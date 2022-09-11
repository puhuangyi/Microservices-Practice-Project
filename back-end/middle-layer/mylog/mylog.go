// Package mylog implements functions to define
// the log output format and create log output.

package mylog

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	LogAllConfig LogAllConfigs
)

// init initialize log configuration.
// Read the log configuration file from the config directory
// and save the data in the file to the corresponding structure.
func init() {
	ymlFile, err := ioutil.ReadFile("./config/log.yml") //Read the log configuration file
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(ymlFile, &LogAllConfig) //Save data in structure
	if err != nil {
		panic(err)
	}
	fmt.Println(LogAllConfig)
}

// LogConfigs store all log configuration information
type LogAllConfigs struct {
	LogConfigs []LogConfig `yaml:"LogConfig"`
}

// ServerLogConfig store server(middle layer) log configuration information
type LogConfig struct {
	Name          string `yaml:"name"`
	LogName       string `yaml:"logName"`          // The prefix name of the log file to output
	Format        string `yaml:"logFormat"`        // The suffix name of the log file to output
	RotationTime  int    `yaml:"logRotationTime"`  // Maximum time for rotating logs
	MaxSavingTime int    `yaml:"logMaxSavingTime"` // Maximum time for storing logs
}

// MyFormatter define the format of each log output
type MyFormatter struct{}

//Logger returns a logger for log output
func Logger(logCon *LogConfig) *logrus.Logger {
	//New a logrus
	logClient := logrus.New()

	//Disable logrus output
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logClient.Out = src // Set output

	//Set logurus output formatter
	//logClient.SetFormatter(new(MyFormatter))
	logClient.SetLevel(logrus.DebugLevel)

	//Set periodically rotates log files
	logWriter, _ := rotatelogs.New(
		logCon.LogName+logCon.Format,
		rotatelogs.WithLinkName(logCon.LogName+".log"),                            //Generate a soft link pointing to the latest log file
		rotatelogs.WithMaxAge(time.Duration(logCon.MaxSavingTime)*time.Hour),      //Set the maximum file retention time
		rotatelogs.WithRotationTime(time.Duration(logCon.RotationTime)*time.Hour), //Set the log cutting interval
	)

	//Create a Hook
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &MyFormatter{})

	//Add hook
	logClient.AddHook(lfHook)

	return logClient
}

// ServerLogger returns a function(middleware) that calculate the request
// processing time and so on, then prints the log.
// This middleware is provided for server log.
func ServerLogger() gin.HandlerFunc {
	var serverLog *logrus.Logger
	for _, config := range LogAllConfig.LogConfigs {
		if config.Name == "server" {
			serverLog = Logger(&config)
			break
		}
	}
	serverLog.Info("Server log is created")

	//Return log middleware
	return func(c *gin.Context) {
		start := time.Now() //Start time

		//Handle the request
		c.Next()

		end := time.Now()         //End time
		latency := end.Sub(start) //Execution time

		path := c.Request.URL.Path      //The path to request
		clientIP := c.ClientIP()        //IP address of the requester
		method := c.Request.Method      //The method to request
		statusCode := c.Writer.Status() //The response status code

		//Print log
		serverLog.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}

// GinLogger returns rotates log files output
func GinLogger() *rotatelogs.RotateLogs {
	var GinLogConfig LogConfig

	//Find gin log config
	for _, x := range LogAllConfig.LogConfigs {
		if x.Name == "gin" {
			GinLogConfig = x
		}
	}

	//Create gin log output
	ginWriter, _ := rotatelogs.New(
		GinLogConfig.LogName+GinLogConfig.Format,
		rotatelogs.WithLinkName(GinLogConfig.LogName+".log"),                            //Generate a soft link pointing to the latest log file
		rotatelogs.WithMaxAge(time.Duration(GinLogConfig.MaxSavingTime)*time.Hour),      //Set the maximum file retention time
		rotatelogs.WithRotationTime(time.Duration(GinLogConfig.RotationTime)*time.Hour), //Set the log cutting interval
	)

	return ginWriter
}

//AuthLogger returns a Logger for authentication moduel log output
func AuthLogger() *logrus.Logger {
	var authConfig LogConfig
	for _, config := range LogAllConfig.LogConfigs {
		if config.Name == "authentication" {
			authConfig = config
		}
	}
	return Logger(&authConfig)
}

//PaymentLogger returns a Logger for payment moduel log output
func PaymentLogger() *logrus.Logger {
	var paymentConfig LogConfig
	for _, config := range LogAllConfig.LogConfigs {
		if config.Name == "payment" {
			paymentConfig = config
		}
	}
	return Logger(&paymentConfig)
}

//LocationLogger returns a Logger for location moduel log output
func LocationLogger() *logrus.Logger {
	var locationConfig LogConfig
	for _, config := range LogAllConfig.LogConfigs {
		if config.Name == "location" {
			locationConfig = config
		}
	}
	return Logger(&locationConfig)
}

//LoginLogger returns a Logger for login moduel log output
func LoginLogger() *logrus.Logger {
	var loginConfig LogConfig
	for _, config := range LogAllConfig.LogConfigs {
		if config.Name == "login" {
			loginConfig = config
		}
	}
	return Logger(&loginConfig)
}

//QueryLogger returns a Logger for query moduel log output
func QueryLogger() *logrus.Logger {
	var loginConfig LogConfig
	for _, config := range LogAllConfig.LogConfigs {
		if config.Name == "query" {
			loginConfig = config
		}
	}
	return Logger(&loginConfig)
}

// Format returns custom log output format header
func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	time := strings.Fields(entry.Time.String())
	time = time[0:2] //strings.Join(time, " ")
	msg := fmt.Sprintf("[%s] [%s %-16s]  %s\n", entry.Level, time[0], time[1], entry.Message)
	return []byte(msg), nil
}
