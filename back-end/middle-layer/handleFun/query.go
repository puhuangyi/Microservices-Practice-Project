package handleFun

import (
	"context"
	"encoding/json"
	"middle-layer/mylog"
	"middle-layer/proto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	queryLog  *logrus.Logger //Define query log output
	connQuery *grpc.ClientConn
)

func init() {
	//Init log output
	queryLog = mylog.QueryLogger()

	var err error

	//Connect to query service
	connQuery, err = grpc.Dial("query-service:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		queryLog.Errorf("| Query.init | %v |", err)
	}
}

type TransportFare struct {
	TransportName string  `json:"name,omitempty"`
	TransportType string  `json:"type,omitempty"`
	Fare          float32 `json:"fare"`
}

type RouteAllInfo struct {
	RouteID       int32                      `json:"routeID"`
	Original      string                     `json:"original,omitempty"`
	Destination   string                     `json:"destination,omitempty"`
	Info          []*proto.RouteToCastleInfo `json:"info,omitempty"`
	Time          []string                   `json:"time,omitempty"`
	TransportFare []*TransportFare           `json:"transportFare,omitempty"`
}

//QueryTimetable
// func QueryTimetable(c *gin.Context) {
// 	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

// 	start := time.Now() //Start time

// 	//Get parameter
// 	noc, ok1 := c.GetQuery("noc")
// 	day, ok2 := c.GetQuery("day")

// 	//Check parameter if exist
// 	if !ok1 || !ok2 || len(noc) == 0 || len(day) == 0 {
// 		queryLog.Infof("| QueryTimetable.checkParameter | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter not full")
// 		c.String(406, "please provide full parameter")
// 		return
// 	}

// 	//Check parameter if vaild
// 	dayInt, err := strconv.Atoi(day)
// 	if err != nil || dayInt < 1 || dayInt > 7 || len(noc) > 5 || len(c.Param("busName")) > 6 {
// 		queryLog.Infof("| QueryTimetable.checkParameter | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter invaild")
// 		c.String(406, "please provide vaild parameter")
// 		return
// 	}

// 	//gprc client
// 	server := proto.NewQueryServiceClient(connQuery)

// 	//Packing parameters for gRPC call
// 	req := &proto.TimetableRequest{
// 		Noc:     noc,
// 		BusName: c.Param("busName"),
// 		Day:     dayChange(dayInt),
// 	}

// 	//Query timetable
// 	res, err := server.QueryTimetable(context.Background(), req)
// 	if err != nil || res == nil {
// 		c.String(406, "no result, please check you parameter is correct!")
// 		if err != nil {
// 			queryLog.Infof("| QueryTimetable.querytimetable | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
// 			return
// 		}
// 		queryLog.Infof("| QueryTimetable.querytimetable | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "res = nil")
// 		return
// 	}

// 	//Change to Json type
// 	resByte, err := json.Marshal(&res)
// 	if err != nil {
// 		queryLog.Infof("| QueryTimetable.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
// 		c.String(500, err.Error())
// 		return
// 	}

// 	//Send response
// 	c.Writer.Header().Add("Content-Type", "application/json")
// 	c.Writer.WriteHeader(200)
// 	c.Writer.Write(resByte)

// 	queryLog.Infof("| QueryTimetable | success | %10v | %15s |", time.Since(start), c.ClientIP())
// }

//QueryCastleFare
func QueryCastleFare(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	castle, ok := c.GetQuery("castle")

	//Check parameter
	if !ok || len(castle) == 0 {
		queryLog.Infof("| QueryCastleFare.checkParameter | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter not full")
	}

	//gprc client
	server := proto.NewQueryServiceClient(connQuery)

	//Packing parameters for gRPC call
	req := &proto.CastleFareRequest{
		Castle: castle,
	}

	//Query castle fare
	res, err := server.QueryCastleFare(context.Background(), req)
	if err != nil || res == nil {
		c.String(406, "no result, please check you parameter is correct!")
		if err != nil {
			queryLog.Infof("| QueryCastleFare.queryCastleFare | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
			return
		}
		queryLog.Infof("| QueryCastleFare.queryCastleFare | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "res = nil")
		return
	}

	//Change to Json type
	resByte, err := json.Marshal(&res)
	if err != nil {
		queryLog.Infof("| QueryCastleFare.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(resByte)

	queryLog.Infof("| QueryCastleFare | success | %10v | %15s |", time.Since(start), c.ClientIP())

}

//QueryCastleInfo
func QueryCastleInfo(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	castle, ok1 := c.GetQuery("castle")
	day, ok2 := c.GetQuery("day")

	//Check parameter if exist
	if !ok1 || !ok2 || len(castle) == 0 || len(day) == 0 {
		queryLog.Infof("| QueryCastleInfo.checkParameterExist | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter not full")
		c.String(406, "please provide full parameter")
		return
	}

	//Check parameter if vaild
	dayInt, err := strconv.Atoi(day)
	if err != nil || dayInt < 1 || dayInt > 7 {
		queryLog.Infof("| QueryCastleInfo.checkParameterVaild | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter invaild")
		c.String(406, "please provide vaild parameter")
		return
	}

	//gprc client
	server := proto.NewQueryServiceClient(connQuery)

	//Packing parameters for gRPC call
	req := &proto.CastleInfoRequest{
		Castle: castle,
		Day:    dayChange(dayInt),
	}

	//Query castle information
	res, err := server.QueryCastleInfo(context.Background(), req)
	if err != nil || res == nil {
		c.String(406, "no result, please check you parameter is correct!")
		if err != nil {
			queryLog.Infof("| QueryCastleInfo.queryCastleInfo | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
			return
		}
		queryLog.Infof("| QueryCastleInfo.queryCastleInfo | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "res = nil")
		return
	}

	//Change to Json type
	resByte, err := json.Marshal(&res)
	if err != nil {
		queryLog.Infof("| QueryCastleInfo.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(resByte)

	queryLog.Infof("| QueryCastleInfo | success | %10v | %15s |", time.Since(start), c.ClientIP())
}

//QueryBusFare
func QueryBusFare(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	noc, ok := c.GetQuery("noc")

	//Check parameter if exist
	if !ok || len(noc) == 0 {
		queryLog.Infof("| QueryBusFare.checkParameterExist | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter not full")
		c.String(406, "please provide full parameter")
		return
	}

	//Check parameter if vaild
	if len(noc) > 6 {
		queryLog.Infof("| QueryBusFare.checkParameterVaild | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter invaild")
		c.String(406, "please provide vaild parameter")
		return
	}

	//gprc client
	server := proto.NewQueryServiceClient(connQuery)

	//Packing parameters for gRPC call
	req := &proto.BusFareRequest{
		BusName: c.Param("busName"),
		Noc:     noc,
	}

	//Query Bus Fare
	res, err := server.QueryBusFare(context.Background(), req)
	if err != nil || res == nil {
		c.String(406, "no result, please check you parameter is correct!")
		if err != nil {
			queryLog.Infof("| QueryBusFare.queryBusFare | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
			return
		}
		queryLog.Infof("| QueryBusFare.queryBusFare | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "res = nil")
		return
	}

	//Change to Json type
	resByte, err := json.Marshal(&res)
	if err != nil {
		queryLog.Infof("| QueryBusFare.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(resByte)

	queryLog.Infof("| QueryBusFare | success | %10v | %15s |", time.Since(start), c.ClientIP())
}

//QueryTrainFare
func QueryTrainFare(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//gprc client
	server := proto.NewQueryServiceClient(connQuery)

	//Packing parameters for gRPC call
	req := &proto.TrainFareRequest{
		TrainName: c.Param("trainName"),
	}

	//Query Train Fare
	res, err := server.QueryTrainFare(context.Background(), req)
	if err != nil || res == nil {
		c.String(406, "no result, please check you parameter is correct!")
		if err != nil {
			queryLog.Infof("| QueryTrainFare.queryTrainFare | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
			return
		}
		queryLog.Infof("| QueryTrainFare.queryTrainFare | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "res = nil")
		return
	}

	//Change to Json type
	resByte, err := json.Marshal(&res)
	if err != nil {
		queryLog.Infof("| QueryTrainFare.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(resByte)

	queryLog.Infof("| QueryTrainFare | success | %10v | %15s |", time.Since(start), c.ClientIP())
}

//QueryRoutetoCastle
func QueryRouteToCastle(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	original, ok1 := c.GetQuery("original")
	destination, ok2 := c.GetQuery("destination")

	//Check parameter if exist
	if !ok1 || !ok2 || len(original) == 0 || len(destination) == 0 {
		queryLog.Infof("| QueryRoutetoCastle.checkParameter | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter not full")
		c.String(406, "please provide full parameter")
		return
	}

	//gprc client
	server := proto.NewQueryServiceClient(connQuery)

	//Packing parameters for gRPC call
	req := &proto.RouteToCastleRequest{
		Original:    original,
		Destination: destination,
	}

	//Query Route To Castle
	res, err := server.QueryRouteToCastle(context.Background(), req)
	if err != nil || res == nil {
		c.String(406, "no result, please check you parameter is correct!")
		if err != nil {
			queryLog.Infof("| QueryRoutetoCastle.queryRoutetoCastle | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
			return
		}
		queryLog.Infof("| QueryRoutetoCastle.queryRoutetoCastle | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "res = nil")
		return
	}

	//Change to Json type
	resByte, err := json.Marshal(&res)
	if err != nil {
		queryLog.Infof("| QueryRoutetoCastle.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(resByte)

	queryLog.Infof("| QueryRoutetoCastle | success | %10v | %15s |", time.Since(start), c.ClientIP())
}

//QueryStartTime
func QueryStartTime(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	routeID, ok1 := c.GetQuery("routeID")
	day, ok2 := c.GetQuery("day")
	if !ok1 || !ok2 {
		queryLog.Infof("| QueryStartTime.getParameter | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter invaild")
		c.String(406, "please provide full parameter")
		return
	}

	//Check parameter if vaild
	routeIDNum, err1 := strconv.Atoi(routeID)
	dayNum, err2 := strconv.Atoi(day)
	if err1 != nil || err2 != nil || routeIDNum <= 0 || dayNum <= 0 || dayNum > 7 {
		queryLog.Infof("| QueryStartTime.checkParameter | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter invaild")
		c.String(406, "please provide vaild parameter")
		return
	}

	//gprc client
	server := proto.NewQueryServiceClient(connQuery)

	//Packing parameters for gRPC call
	req := &proto.StartTimeRequest{
		RouteID: int32(routeIDNum),
		Day:     int32(dayNum),
	}

	//Query start time
	res, err := server.QueryStartTime(context.Background(), req)
	if err != nil || res == nil {
		c.String(406, "no result, please check you parameter is correct!")
		if err != nil {
			queryLog.Infof("| QueryStartTime.queryStartTime | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
			return
		}
		queryLog.Infof("| QueryStartTime.queryStartTime | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "res = nil")
		return
	}

	//Change to Json type
	resByte, err := json.Marshal(&res)
	if err != nil {
		queryLog.Infof("| QueryStartTime.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(resByte)

	queryLog.Infof("| QueryStartTime | success | %10v | %15s |", time.Since(start), c.ClientIP())
}

//QueryRouteAllInfo
func QueryRouteAllInfo(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	original, ok1 := c.GetQuery("original")
	destination, ok2 := c.GetQuery("destination")
	day, ok3 := c.GetQuery("day")

	//Check parameter if exist
	if !ok1 || !ok2 || !ok3 || len(original) == 0 || len(destination) == 0 || len(day) == 0 {
		queryLog.Infof("| QueryRouteAllInfo.checkParameterExist | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter not full")
		c.String(406, "please provide full parameter")
		return
	}

	//Check parameter if vaild
	dayNum, err := strconv.Atoi(day)
	if err != nil || dayNum <= 0 || dayNum > 7 {
		queryLog.Infof("| QueryRouteAllInfo.checkParameterVaild | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "parameter invaild")
		c.String(406, "please provide vaild parameter")
		return
	}

	//gprc client
	server := proto.NewQueryServiceClient(connQuery)

	//Packing parameters for gRPC call
	routeReq := &proto.RouteToCastleRequest{
		Original:    original,
		Destination: destination,
	}

	//Query Route To Castle firstly
	routeRes, err := server.QueryRouteToCastle(context.Background(), routeReq)
	if err != nil || routeRes == nil || len(routeRes.Info) == 0 {
		c.String(406, "no result, please check you parameter is correct!")
		if err != nil {
			queryLog.Infof("| QueryRouteAllInfo.queryRoutetoCastle | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
			return
		}
		queryLog.Infof("| QueryRouteAllInfo.queryRoutetoCastle | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "res = nil")
		return
	}

	fareList := make([]*TransportFare, 0)

	//Packing
	timeReq := &proto.StartTimeRequest{
		RouteID: routeRes.RouteID,
		Day:     int32(dayNum),
	}

	//Query start time
	timeRes, err := server.QueryStartTime(context.Background(), timeReq)
	if err != nil || timeRes == nil {
		c.String(406, "no result, please check you parameter is correct!")
		if err != nil {
			queryLog.Infof("| QueryRouteAllInfo.queryStartTime | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
			return
		}
		queryLog.Infof("| QueryRouteAllInfo.queryStartTime | %10v | %15s | %s |", time.Since(start), c.ClientIP(), "res = nil")
		return
	}

	//Add transport fare
	for _, data := range routeRes.Info {
		if data.Type == "bus" {
			fareReq := &proto.BusFareRequest{
				BusName: data.TransportName,
				Noc:     data.Noc,
			}

			//Query fare
			fareRes, err := server.QueryBusFare(context.Background(), fareReq)
			if err != nil || fareRes == nil {
				queryLog.Infof("| QueryRouteAllInfo.queryBusFare | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
				continue
			}

			//Add fare
			fareList = append(fareList, &TransportFare{
				TransportName: data.TransportName,
				TransportType: data.Type,
				Fare:          float32(fareRes.Fare),
			})
		} else if data.Type == "train" {
			fareReq := &proto.TrainFareRequest{
				TrainName: data.TransportName,
			}

			//Query fare
			fareRes, err := server.QueryTrainFare(context.Background(), fareReq)
			if err != nil || fareRes == nil {
				queryLog.Infof("| QueryRouteAllInfo.queryTrainFare | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
				continue
			}

			//Add fare
			fareList = append(fareList, &TransportFare{
				TransportName: data.TransportName,
				TransportType: data.Type,
				Fare:          float32(fareRes.Fare),
			})
		}
	}

	//Result
	result := &RouteAllInfo{
		RouteID:       routeRes.RouteID,
		Original:      original,
		Destination:   destination,
		Info:          routeRes.Info,
		Time:          timeRes.Time,
		TransportFare: fareList,
	}

	//Change to Json type
	resByte, err := json.Marshal(&result)
	if err != nil {
		queryLog.Infof("| QueryRouteAllInfo.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(resByte)

	queryLog.Infof("| QueryRouteAllInfo | success | %10v | %15s |", time.Since(start), c.ClientIP())
}

//dayChange return day string which corresponding number
// 1 -> Monday, 5 -> Friday...
func dayChange(dayInt int) string {
	switch dayInt {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	}
	return ""
}

//CloseResource close resource
func CloseResource() {
	connLogin.Close()
	connPayment.Close()
	connLocation.Close()
	connQuery.Close()
}
