//Package service implement query bus location function.
//If you want to understand the specific details of this service,
//please read this:https://data.bus-data.dft.gov.uk/guidance/requirements/?section=casestudies

package service

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"location/myInit"
	"location/mylog"
	"location/proto"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	URL        = "https://data.bus-data.dft.gov.uk/api/v1/datafeed/"  //Query url
	KEY        = "/?api_key=dc56a4f0cfc3feb0e93bd93360a8fc26a7018a1f" //Query API key
	timeFormat = "2006-01-02T15:04:05"                                //Time format
)

//these struct store API response information
//detail:https://data.bus-data.dft.gov.uk/avl/
type Siri struct {
	XMLName         xml.Name        `xml:"Siri"`
	Version         string          `xml:"version,attr"`
	ServiceDelivery ServiceDelivery `xml:"ServiceDelivery"`
}

//store API response information
type ServiceDelivery struct {
	XMLName                   xml.Name                  `xml:"ServiceDelivery"`
	ResponseTimestamp         string                    `xml:"ResponseTimestamp"`
	ProducerRef               string                    `xml:"ProducerRef"`
	VehicleMonitoringDelivery VehicleMonitoringDelivery `xml:"VehicleMonitoringDelivery"`
}

//store API response information
type VehicleMonitoringDelivery struct {
	XMLName               xml.Name          `xml:"VehicleMonitoringDelivery"`
	ResponseTimestamp     string            `xml:"ResponseTimestamp"`
	RequestMessageRef     string            `xml:"RequestMessageRef"`
	ValidUntil            string            `xml:"ValidUntil"`
	ShortestPossibleCycle string            `xml:"ShortestPossibleCycle"`
	VehicleActivitys      []VehicleActivity `xml:"VehicleActivity"`
}

//store API response information
type VehicleActivity struct {
	XMLName                 xml.Name                `xml:"VehicleActivity"`
	RecordedAtTime          string                  `xml:"RecordedAtTime"`
	ItemIdentifier          string                  `xml:"ItemIdentifier"`
	ValidUntilTime          string                  `xml:"ValidUntilTime"`
	MonitoredVehicleJourney MonitoredVehicleJourney `xml:"MonitoredVehicleJourney"`
	Extensions              Extensions              `xml:"Extensions"`
}

//store API response information
type MonitoredVehicleJourney struct {
	XMLName                  xml.Name                  `xml:"MonitoredVehicleJourney"`
	LineRef                  string                    `xml:"LineRef"`
	DirectionRef             string                    `xml:"DirectionRef"`
	FramedVehicleJourneyRefs []FramedVehicleJourneyRef `xml:"FramedVehicleJourneyRefs"`
	PublishedLineName        string                    `xml:"PublishedLineName"`
	OperatorRef              string                    `xml:"OperatorRef"`
	DestinationRef           string                    `xml:"DestinationRef"`
	VehicleLocation          VehicleLocation           `xml:"VehicleLocation"`
	Bearing                  float32                   `xml:"Bearing"`
	Occupancy                string                    `xml:"Occupancy"`
	BlockRef                 string                    `xml:"BlockRef"`
	VehicleRef               string                    `xml:"VehicleRef"`
}

//store API response information
type FramedVehicleJourneyRef struct {
	XMLName                xml.Name `xml:"FramedVehicleJourneyRef"`
	DataFrameRef           string   `xml:"DataFrameRef"`
	DatedVehicleJourneyRef int      `xml:"DatedVehicleJourneyRef"`
}

//store API response information
type VehicleLocation struct {
	XMLName   xml.Name `xml:"VehicleLocation"`
	Longitude float64  `xml:"Longitude"`
	Latitude  float64  `xml:"Latitude"`
}

//store API response information
type Extensions struct {
	XMLName        string         `xml:"Extensions"`
	VehicleJourney VehicleJourney `xml:"VehicleJourney"`
}

//store API response information
type VehicleJourney struct {
	XMLName             xml.Name    `xml:"VehicleJourney"`
	Operationals        Operational `xml:"Operational"`
	VehicleUniqueId     string      `xml:"VehicleUniqueId"`
	SeatedOccupancy     int         `xml:"SeatedOccupancy"`
	SeatedCapacity      int         `xml:"SeatedCapacity"`
	WheelchairOccupancy int         `xml:"WheelchairOccupancy"`
	WheelchairCapacity  int         `xml:"WheelchairCapacity"`
	OccupancyThresholds string      `xml:"OccupancyThresholds"`
}

//store API response information
type Operational struct {
	XMLName        xml.Name        `xml:"Operational"`
	TicketMachines []TicketMachine `xml:"TicketMachine"`
}

//store API response information
type TicketMachine struct {
	XMLName                  xml.Name `xml:"TicketMachine"`
	TicketMachineServiceCode string   `xml:"TicketMachineServiceCode"`
	JourneyCode              int      `xml:"JourneyCode"`
}

//feedList store feedID list
type feedList struct {
	feed []int
}

//LocationService used for gPRC server
type LocationService struct {
}

//QueryLocation returns ResBusInfos which include bus location
//and seat information.
//Firstly, check bus company if in common bus list, if yes, use
//the corresponding ID to query the bus location. Otherwise,
//query all the feedIDs of this company and save them in the list.
//Then, Find the bus location one by one according to the feedID
//in the list.
func (s *LocationService) QueryLocation(ctx context.Context, busInfo *proto.BusInfo) (*proto.ResBusInfos, error) {
	start := time.Now()

	//Query company through noc
	var company string
	rows, err := myInit.Db.Query("select operator from operator_noc where noc = ?", busInfo.Reginol)
	if err != nil {
		mylog.LogClient.Infof("| QueryLocation.queryCompany | %10v | %15s | %6s | %v |", time.Since(start), busInfo.Reginol, busInfo.BusName, err)
		return nil, err
	}

	defer rows.Close()

	//Read company name
	for rows.Next() {
		err = rows.Scan(&company)
		if err != nil {
			mylog.LogClient.Infof("| QueryLocation.checkCompany | %10v | %15s | %6s | %v |", time.Since(start), busInfo.Reginol, busInfo.BusName, err)
			return nil, err
		}
	}

	//Check
	if len(company) == 0 {
		err = errors.New("noc invaild")
		mylog.LogClient.Infof("| QueryLocation.checkCompany | %10v | %15s | %6s | %v |", time.Since(start), busInfo.Reginol, busInfo.BusName, err)
		return nil, err
	}

	//Check company if in the common list
	for _, x := range myInit.CommonList.List {

		//If company in the list
		if company == x.Name {

			//find table name
			cNAME := strings.ReplaceAll(x.Name, "-", " ")
			tableName := "location_commonTable_" + strings.ReplaceAll(cNAME, " ", "_")

			var feed int
			//Query feedID
			err := myInit.Db.QueryRow("SELECT Datafeed_ID FROM "+tableName+" WHERE LineName=?", busInfo.BusName).Scan(&feed) //Save result
			if err != nil {
				mylog.LogClient.Infof("| QueryLocation.queryFeedID | %10v | %s | %6s | %v |", time.Since(start), busInfo.Reginol, busInfo.BusName, err)
				return nil, err
			}

			//Query location
			result, err := queryLocationWithFeedID(busInfo.BusName, feed)
			if err != nil {
				mylog.LogClient.Infof("| QueryLocation.queryLocationWithFeedID | %10v | %s | %6s | %v |", time.Since(start), busInfo.Reginol, busInfo.BusName, err)
				return nil, err
			}

			mylog.LogClient.Infof("| QueryLocation | success | %10v | %s | %6s | %1s | feed:%d |", time.Since(start), busInfo.Reginol, busInfo.BusName, tableName, feed)
			return result, err
		}
	}

	//Query all the feedIDs of this company
	list, err := queryFeedIDList(company, busInfo.BusName)
	if err != nil {
		mylog.LogClient.Infof("| QueryLocation.queryFeedIDList | %10v | %s | %6s | %v |", time.Since(start), busInfo.Reginol, busInfo.BusName, errors.New("could not find bus name"))
		return nil, err
	}

	mylog.LogClient.Infof("| QueryLocation | %10v | %s | %6s | %v |", time.Since(start), busInfo.Reginol, busInfo.BusName, list)

	//Find the bus location according to the feedID in the list
	return queryLocationWithFeedIDList(list, busInfo.BusName)
}

//queryFeedIDList returns a list which include all the feedIDs of this company
func queryFeedIDList(busCompany string, busName string) (*feedList, error) {

	//Query FeedID
	rows, err := myInit.Db.Query("select Datafeed_ID from location_data_overall where Organisation_Name = ?", busCompany)
	if err != nil {
		mylog.LogClient.Infof("| queryFeedIDList.Query | %s | %6s | %v |", busCompany, busName, err)
		return nil, err
	}

	//Close rows
	defer rows.Close()

	//Init feedID list
	list := &feedList{
		feed: make([]int, 0),
	}

	//Read data in loop
	for rows.Next() {
		var result int
		err := rows.Scan(&result) //Save feedID
		if err != nil {
			mylog.LogClient.Infof("| queryFeedIDList.Scan | %s | %6s | %v |", busCompany, busName, err)
			return nil, err
		}

		if result != 0 {
			list.feed = append(list.feed, result) //Append feedID in list
		}
	}

	//Check if have data
	if len(list.feed) == 0 {
		err = errors.New("could not found feedID")
		mylog.LogClient.Infof("| queryFeedIDList.checkFeedID | %s | %6s | %v |", busCompany, busName, err)
		return nil, err
	}

	return list, nil
}

//queryLocationWithFeedIDList function find the bus location
//according to the feedID in the list.
func queryLocationWithFeedIDList(list *feedList, busName string) (*proto.ResBusInfos, error) {

	//Query in loop
	for _, feed := range list.feed {
		//Query with feedID
		result, err := queryLocationWithFeedID(busName, feed)
		if err != nil {
			mylog.LogClient.Infof("| queryLocationWithFeedIDList | %6s | feed:%d | %v |", busName, feed, err)
			return nil, err
		}

		if len(result.Info) > 0 {
			return result, nil
		}
	}

	err := errors.New("could not found info")
	mylog.LogClient.Infof("| queryLocationWithFeedIDList | %6s | %v |", busName, err)
	return nil, err
}

//queryLocationWithFeedID function find the bus location based on the specified feedID
func queryLocationWithFeedID(busName string, feedID int) (*proto.ResBusInfos, error) {
	//Init result
	result := &proto.ResBusInfos{
		Info: make([]*proto.ResBusInfo, 0),
	}

	//Spliced into a complete URL
	requestUrl := URL + strconv.Itoa(feedID) + KEY
	fmt.Printf("url is:%s\n", requestUrl)

	//Send http request
	resp, err := http.Get(requestUrl)
	if err != nil {
		mylog.LogClient.Infof("| queryLocationWithFeedID | %6s | feed:%d | %v |", busName, feedID, err)
		return nil, err
	}

	defer resp.Body.Close()

	//Read the response
	var res []byte
	if resp.Status == "200 OK" {
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			mylog.LogClient.Infof("| queryLocationWithFeedID | %6s | feed:%d | %v |", busName, feedID, err)
			return nil, err
		}
	} else {
		err = errors.New("response status is not 200")
		mylog.LogClient.Infof("| queryLocationWithFeedID | %6s | feed:%d | %v |", busName, feedID, err)
		return nil, err
	}

	//Change to struct
	data := &Siri{}
	err = xml.Unmarshal(res, data) //change
	if err != nil {
		mylog.LogClient.Infof("| queryLocationWithFeedID | %6s | feed:%d | %v |", busName, feedID, err)
		return nil, err
	}

	//Find each bus data in loop
	for _, x := range data.ServiceDelivery.VehicleMonitoringDelivery.VehicleActivitys {

		//Check bus name if equal
		if x.MonitoredVehicleJourney.LineRef == busName {

			//Get record time
			recoedTime := strings.Split(x.RecordedAtTime, "+")
			timeString, _ := time.Parse(timeFormat, recoedTime[0])

			//Filter out records that have been logged for a long time
			if time.Since(timeString).Minutes() < 60 {

				//Get seat information
				var seat int
				if x.Extensions.VehicleJourney.SeatedCapacity == 0 {
					seat = -1
				} else {
					seat = x.Extensions.VehicleJourney.SeatedCapacity - x.Extensions.VehicleJourney.SeatedOccupancy
				}

				//Get wheel seat information
				var wheel int
				if x.Extensions.VehicleJourney.WheelchairCapacity == 0 {
					wheel = -1
				} else {
					wheel = x.Extensions.VehicleJourney.WheelchairCapacity - x.Extensions.VehicleJourney.WheelchairOccupancy
				}

				//Result
				result.Info = append(result.Info, &proto.ResBusInfo{
					BusRef:    x.MonitoredVehicleJourney.VehicleRef,
					Latitude:  x.MonitoredVehicleJourney.VehicleLocation.Latitude,
					Longitude: x.MonitoredVehicleJourney.VehicleLocation.Longitude,
					SeatFree:  int32(seat),
					WheelFree: int32(wheel),
				})
			}
		}
	}

	return result, nil
}

//QueryTest for test. This function is useless so no comment.
/*
func QueryTest(busCompany string, busName string) {
	for _, x := range myInit.CommonList.List {
		if busCompany == x.Name {

			//table name
			cNAME := strings.ReplaceAll(x.Name, "-", " ")
			tableName := "commonTable_" + strings.ReplaceAll(cNAME, " ", "_")
			fmt.Printf("Query table result is:%s\n", tableName)

			var feed int
			//Query feedID
			err := myInit.Db.QueryRow("SELECT Datafeed_ID FROM "+tableName+" WHERE LineName=?", busName).Scan(&feed) //Save result
			if err != nil {
				panic(err)
			}

			fmt.Printf("Query id result is:%d\n", feed)

			//Query location
			result, err := queryLocationWithFeedID(busName, feed)
			if err != nil {
				panic(err)
			}

			fmt.Printf("result:%v\n", &result.Info)

			return
		}
	}

	list, err := queryFeedIDList(busCompany, busName)
	if err != nil {
		panic(err)
	}

	result, err := queryLocationWithFeedIDList(list, busName)
	if result != nil {
		fmt.Printf("no result")
	}

	fmt.Printf("result is:%v", &result.Info)

}
*/
