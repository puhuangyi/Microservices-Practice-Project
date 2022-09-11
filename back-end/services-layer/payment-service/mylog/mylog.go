package mylog

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	logConfig LogConfigs
	LogClient *logrus.Logger
)

func init() {
	//Read the log configuration file
	ymlFile, err := ioutil.ReadFile("./config/log.yml")
	if err != nil {
		panic(err)
	}

	//Save data in structure
	err = yaml.Unmarshal(ymlFile, &logConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(logConfig)

	initLog()
}

type LogConfigs struct {
	LogName       string `yaml:"logName"`          // The prefix name of the log file to output
	Format        string `yaml:"logFormat"`        // The suffix name of the log file to output
	RotationTime  int    `yaml:"logRotationTime"`  // Maximum time for rotating logs
	MaxSavingTime int    `yaml:"logMaxSavingTime"` // Maximum time for storing logs
}

// MyFormatter define the format of each log output
type MyFormatter struct{}

//initLog initialize log output
func initLog() {
	//New a logrus
	LogClient = logrus.New()

	//Disable logrus output
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	LogClient.Out = src // Set output

	//Set logurus output formatter
	LogClient.SetFormatter(new(MyFormatter))
	LogClient.SetLevel(logrus.DebugLevel)

	//Set periodically rotates log files
	logWriter, _ := rotatelogs.New(
		logConfig.LogName+logConfig.Format,
		rotatelogs.WithLinkName(logConfig.LogName+".log"),                            //Generate a soft link pointing to the latest log file
		rotatelogs.WithMaxAge(time.Duration(logConfig.MaxSavingTime)*time.Hour),      //Set the maximum file retention time
		rotatelogs.WithRotationTime(time.Duration(logConfig.RotationTime)*time.Hour), //Set the log cutting interval
	)

	//Create a Hook
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &MyFormatter{})

	//Add hook
	LogClient.AddHook(lfHook)
}

// Format returns custom log output format header
func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	time := strings.Fields(entry.Time.String())
	time = time[0:2] //strings.Join(time, " ")
	msg := fmt.Sprintf("[%s] [%s %-16s]  %s\n", entry.Level, time[0], time[1], entry.Message)
	return []byte(msg), nil
}
