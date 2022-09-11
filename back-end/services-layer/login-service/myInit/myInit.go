package myInit

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"login/mylog"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var (
	Db *sql.DB
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
