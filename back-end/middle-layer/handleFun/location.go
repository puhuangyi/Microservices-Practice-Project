package handleFun

import (
	"context"
	"encoding/json"
	"errors"
	"middle-layer/mylog"
	"middle-layer/proto"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	locationLog  *logrus.Logger //Define location log output
	connLocation *grpc.ClientConn
)

func init() {
	//Init log output
	locationLog = mylog.LocationLogger()

	var err error

	//Connect to location service
	connLocation, err = grpc.Dial("location-service:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		locationLog.Errorf("| Location.init | %v |", err)
	}
}

//QueryLocation function handle query bus location request.
//Firstly, check the parameter if correct.
//Secondly, call QueryLocation function in location service with gRPC.
//Finally, check QueryLocation if successful then return response.
func QueryLocation(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	noc, ok := c.GetQuery("noc")

	//Check parameter
	if !ok || len(noc) == 0 {
		locationLog.Infof("| getQuery | %10v | %15s | %v |", time.Since(start), c.ClientIP(), errors.New("request did not provide noc"))
		c.String(406, "Please provide noc(national operator code)")
		return
	}

	server := proto.NewLocationServiceClient(connLocation)

	//Packing parameters for gRPC call
	req := &proto.BusInfo{
		BusName: c.Param("busName"),
		Reginol: noc,
	}

	//Query location
	res, err := server.QueryLocation(context.Background(), req)
	if err != nil {
		locationLog.Infof("| queryLocation | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(400, "please enter correct parameter")
		return
	}

	//Change to Json type
	resByte, err := json.Marshal(&res)
	if err != nil {
		locationLog.Infof("| marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(resByte)

	locationLog.Infof("| success | %10v | %15s | %s %s |", time.Since(start), c.ClientIP(), noc, c.Param("busName"))
}
