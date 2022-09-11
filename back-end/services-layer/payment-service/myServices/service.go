package myServices

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"payment/mylog"
	"payment/proto"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var (
	Db *sql.DB //Mysql connection pool
)

const METHOD = "POST"
const CONTENT_TYPE = "application/json"
const layout = "02/01/2006 15:04"
const URL = "http://homepages.cs.ncl.ac.uk/daniel.nesbitt/CSC8019/HorsePay/HorsePay.php"

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
	Db.SetMaxOpenConns(100)
	Db.SetMaxIdleConns(4)
}

//mysqlConfig store Mysql config
type mysqlConfig struct {
	DbName   string `yaml:"dbName"`
	Host     string `yaml:"host"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}

type MyService struct {
}

type payInfo struct {
	StoreID           string  `json:"storeID"`
	CustomerID        string  `json:"customerID"`
	Date              string  `json:"date"`
	Time              string  `json:"time"`
	TimeZone          string  `json:"timeZone"`
	TransactionAmount float64 `json:"transactionAmount"`
	CurrencyCode      string  `json:"currencyCode"`
}

type ResPayStatus struct {
	Status bool   `json:"Status"`
	Reason string `json:"reason"`
}

type ResponseInfo struct {
	payInfo
	ResPayStatus `json:"paymetSuccess"`
}

//Payment function return payment result, if pay success, will record
//order in database
func (s *MyService) Payment(ctx context.Context, info *proto.PaymentInfo) (*proto.ResPayInfo, error) {
	//Now time
	start := time.Now()
	timeNow := strings.Split(start.Format(layout), " ")

	//Get parameter
	config := &payInfo{
		StoreID:           "TEAM22",
		CustomerID:        info.CustomerID,
		Date:              timeNow[0],
		Time:              timeNow[1],
		TimeZone:          "GMT",
		TransactionAmount: info.TransactionAmount,
		CurrencyCode:      "GBP",
	}

	//Change to json
	configData, _ := json.Marshal(config)
	param := bytes.NewBuffer([]byte(configData))

	//Init http request
	client := &http.Client{}
	req, err := http.NewRequest(METHOD, URL, param)
	if err != nil {
		mylog.LogClient.Infof("| newRequest | %13v | %s | %v |", time.Since(start), info.CustomerID, err)
		return nil, err
	}

	//Add header
	req.Header.Add("Content-Type", CONTENT_TYPE)

	//Sent request
	res, err := client.Do(req)
	if err != nil {
		mylog.LogClient.Infof("| sentRequest | %13v | %s | %v |", time.Since(start), info.CustomerID, err)
		return nil, err
	}
	defer res.Body.Close()

	//Check respose
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		mylog.LogClient.Infof("| readResponse | %13v | %s | %v |", time.Since(start), info.CustomerID, err)
		return nil, err
	}

	//Unmarshal
	resData := &ResponseInfo{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		mylog.LogClient.Infof("| unmarshal | %13v | %s | %v |", time.Since(start), info.CustomerID, err)
		return nil, err
	}

	//Record order in database
	if resData.Status {
		sqlStr := "insert into orderInfo( userID, amount, start, destination, number, routerDetailID, time, date, type) values(?,?,?,?,?,?,?,?,?)"
		_, err := Db.Exec(sqlStr, info.CustomerID, info.TransactionAmount, info.Start, info.Destination, info.Number, info.RouteDetailID, start.String()[:22], info.Date, info.Type)
		if err != nil {
			mylog.LogClient.Errorf("| Exec | %13v | %s | %v|", time.Since(start), info.CustomerID, err)
		}
	}

	mylog.LogClient.Infof("| %13v | %s | successful |", time.Since(start), info.CustomerID)

	return &proto.ResPayInfo{
		Status: resData.Status,
		Reason: resData.Reason,
	}, nil
}

func (s *MyService) QueryOrder(ctx context.Context, queryInfo *proto.QueryOrderUserID) (*proto.OrderInfoList, error) {
	//Now time
	start := time.Now()

	//Query
	rows, err := Db.Query("select orderID, amount, start, destination, number, routerDetailID, time, date, type from orderInfo where userID = ?", queryInfo.UserID)
	if err != nil {
		mylog.LogClient.Infof("| query | %13v | %s | %v |", time.Since(start), queryInfo.UserID, err)
		return nil, err
	}

	defer rows.Close()
	result := &proto.OrderInfoList{}

	//Read data
	for rows.Next() {
		order := &proto.OrderInfo{}
		err = rows.Scan(&order.OrderID, &order.Amount, &order.Start, &order.Destination, &order.Number, &order.RouteDetailID, &order.Time, &order.Date, &order.Type)
		if err != nil {
			mylog.LogClient.Infof("| rows.next | %13v | %s | %v |", time.Since(start), queryInfo.UserID, err)
			return nil, err
		}
		result.Orders = append(result.Orders, order)
	}

	mylog.LogClient.Infof("| %13v | %s | successful |", time.Since(start), queryInfo.UserID)
	return result, nil
}
