//Package myInit used to perform some preparatory work.
//For example, add data to the database.
//For companies with too many buses, separate forms will
//be created to store their bus data.

package myInit

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"location/mylog"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var (
	Db         *sql.DB                                                //Mysql connection pool
	wg         sync.WaitGroup                                         //WaitGroup
	URL        = "https://data.bus-data.dft.gov.uk/api/v1/datafeed/"  //Query url
	KEY        = "/?api_key=dc56a4f0cfc3feb0e93bd93360a8fc26a7018a1f" //Query key
	CommonList *commonCompanyList                                     //Common company list store the names of companies that own too many bus routes
)

func init() {
	//Read mysql config
	sqlConfig := &mysqlConfig{}
	sqlFile, err := ioutil.ReadFile("./config/mysql.yml") //Read the log configuration file
	if err != nil {
		mylog.LogClient.Errorf("| read mysql config | %v |", err)
		panic(err)
	}

	//Unmarshal config
	err = yaml.Unmarshal(sqlFile, &sqlConfig) //Save data in structure
	if err != nil {
		mylog.LogClient.Errorf("| unmarshal mysql config | %v |", err)
		panic(err)
	}
	fmt.Println(sqlConfig)

	//Conenect to Mysql
	Db, err = sql.Open("mysql", sqlConfig.UserName+":"+sqlConfig.Password+"@tcp("+sqlConfig.Host+")/"+sqlConfig.DbName+"?charset=utf8")
	if err != nil {
		mylog.LogClient.Errorf("| connect to mysql | %v |", err)
		panic(err)
	}

	//Test
	err = Db.Ping()
	if err != nil {
		mylog.LogClient.Errorf("| ping | %v |", err)
		panic(err)
	}

	//Set connection pool config
	Db.SetMaxOpenConns(0)
	Db.SetMaxIdleConns(10)
	Db.SetConnMaxLifetime(4 * time.Second)

	//Init common company list
	CommonList = &commonCompanyList{
		List: make([]commonCompany, 0),
	}

	//Create bus query table
	creatTable()
}

//mysqlConfig store Mysql config
type mysqlConfig struct {
	DbName   string `yaml:"dbName"`
	Host     string `yaml:"host"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}

//feedIdList store feedID list
type feedIdList struct {
	feedID []int
}

//commonCompany store Company name and its number of feedID
type commonCompany struct {
	Name   string
	Number int
}

//commonCompanyList store commonCompany list
type commonCompanyList struct {
	List []commonCompany
}

//these struct store API response information
//detail:https://data.bus-data.dft.gov.uk/avl/
type Siri struct {
	XMLName         xml.Name        `xml:"Siri"`
	Version         string          `xml:"version,attr"`
	ServiceDelivery ServiceDelivery `xml:"ServiceDelivery"`
}

//these struct store API response information
//detail:https://data.bus-data.dft.gov.uk/avl/
type ServiceDelivery struct {
	XMLName                   xml.Name                  `xml:"ServiceDelivery"`
	VehicleMonitoringDelivery VehicleMonitoringDelivery `xml:"VehicleMonitoringDelivery"`
}

//these struct store API response information
//detail:https://data.bus-data.dft.gov.uk/avl/
type VehicleMonitoringDelivery struct {
	XMLName               xml.Name          `xml:"VehicleMonitoringDelivery"`
	ResponseTimestamp     string            `xml:"ResponseTimestamp"`
	RequestMessageRef     string            `xml:"RequestMessageRef"`
	ValidUntil            string            `xml:"ValidUntil"`
	ShortestPossibleCycle string            `xml:"ShortestPossibleCycle"`
	VehicleActivitys      []VehicleActivity `xml:"VehicleActivity"`
}

//these struct store API response information
//detail:https://data.bus-data.dft.gov.uk/avl/
type VehicleActivity struct {
	XMLName                 xml.Name                `xml:"VehicleActivity"`
	MonitoredVehicleJourney MonitoredVehicleJourney `xml:"MonitoredVehicleJourney"`
}

//these struct store API response information
//detail:https://data.bus-data.dft.gov.uk/avl/
type MonitoredVehicleJourney struct {
	XMLName xml.Name `xml:"MonitoredVehicleJourney"`
	LineRef string   `xml:"LineRef"`
}

//createTable function check which have have too many bus route,
//then store their name in list, finally, create separate tables
//for each company on the list.
func creatTable() {
	//Query which company have more then 5 feedID
	rows, err := Db.Query("select Organisation_Name, num from (select Organisation_Name, count(*) as num  from location_data_overall group by Organisation_Name) as ttt where num >= 5")
	if err != nil {
		mylog.LogClient.Errorf("| creatTable:Query | %v |", err)
		return
	}
	defer rows.Close()

	//read data in loop
	for rows.Next() {
		var data commonCompany

		//Get data
		err := rows.Scan(&data.Name, &data.Number)
		if err != nil {
			mylog.LogClient.Errorf("| creatTable:Scan | %v |", err)
			return
		}

		//Append company in list
		CommonList.List = append(CommonList.List, data)
	}

	//Create table for each company in list
	for _, x := range CommonList.List {
		wg.Add(1)
		go create(x.Name) //create
	}

	//Wait for all goroutine to finish
	wg.Wait()

	mylog.LogClient.Infof("| creatTable successful!!! |")
}

//create function gets a company name to create a separate table
//for this company.
func create(companyName string) {

	start := time.Now()

	//Query all feedID of this company
	sqlStr := "select Datafeed_ID from location_data_overall where Organisation_Name = ?"
	rows, err := Db.Query(sqlStr, companyName)
	if err != nil {
		mylog.LogClient.Errorf("| create:Query | %v |", err)
		return
	}

	defer rows.Close()

	//Init feedID list
	list := &feedIdList{
		feedID: make([]int, 0),
	}

	//read data in loop
	for rows.Next() {
		var result int

		//Get feedID
		err := rows.Scan(&result)
		if err != nil {
			mylog.LogClient.Errorf("| create:Scan | %v |", err)
			return
		}

		//Append
		list.feedID = append(list.feedID, result)
	}

	busMap := make(map[string]int)

	//Query all bus names contained in each feedID
	for _, feed := range list.feedID {

		//Init http request url
		requestUrl := URL + strconv.Itoa(feed) + KEY
		resp, err := http.Get(requestUrl)
		if err != nil {
			mylog.LogClient.Errorf("| create:http.Get | URL:%s | %v |", requestUrl, err)
			continue
		}

		//Get http response body
		res, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			mylog.LogClient.Errorf("| create:ReadAll | %v |", err)
			continue
		}

		data := &Siri{}

		//Unmarshal response body
		err = xml.Unmarshal(res, data)
		if err != nil {
			mylog.LogClient.Errorf("| create:Unmarshal | %v |", err)
			continue
		}

		//Add the bus name with this feedID to the MAP
		for _, x := range data.ServiceDelivery.VehicleMonitoringDelivery.VehicleActivitys {
			busMap[x.MonitoredVehicleJourney.LineRef] = feed
		}
	}

	//Init table name
	cName := strings.ReplaceAll(companyName, "-", " ")
	cName = strings.ReplaceAll(cName, " ", "_")
	tableName := "location_commonTable_" + cName

	//Drop table
	deleteSql, err := Db.Prepare("DROP TABLE  IF EXISTS `" + tableName + "`;")
	if err != nil {
		mylog.LogClient.Errorf("| create:Drop.Prepare | %v |", err)
		return
	}
	deleteSql.Exec()

	//Create table
	createSql, err := Db.Prepare("CREATE TABLE " + tableName + " (Datafeed_ID int(8),LineName varchar(8));")
	if err != nil {
		mylog.LogClient.Errorf("| create:Create.Prepare | %v |", err)
		return
	}
	createSql.Exec()

	//Insert data
	insertSql, err := Db.Prepare("INSERT " + tableName + " (Datafeed_ID, LineName) values (?,?)")
	for d := range busMap {
		insertSql.Exec(busMap[d], d)
		if err != nil {
			mylog.LogClient.Errorf("| create:Insert.Prepare | %v |", err)
			continue
		}
	}

	mylog.LogClient.Infof("| create successful | %13v | companyName is %s |", time.Since(start), companyName)
	fmt.Printf("Time of Handle %s is %v\n", companyName, time.Since(start))

	wg.Done()
}
