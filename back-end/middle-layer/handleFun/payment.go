package handleFun

import (
	"context"
	"encoding/json"
	"errors"
	"middle-layer/mylog"
	"middle-layer/proto"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	paymentLog *logrus.Logger //Define payment log output

	connPayment *grpc.ClientConn
)

func init() {
	//Init payment log output
	paymentLog = mylog.PaymentLogger()

	var err error

	//Connect to payment service
	connPayment, err = grpc.Dial("payment-service:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		paymentLog.Errorf("| Pay.init | %v |", err)
	}
}

//paymentResponseJson store payment response data
type paymentResponseJson struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

//Pay function handle payment request.
//Firstly, check the parameter if correct.
//Secondly, call payment function in payment service with gRPC.
//Finally, check payment if successful then return response.
func Pay(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	customerID, ok1 := c.GetPostForm("userID")
	transactionAmount, ok2 := c.GetPostForm("amount")
	startStop, ok3 := c.GetPostForm("start")
	destination, ok4 := c.GetPostForm("destination")
	number, ok5 := c.GetPostForm("number")
	date, ok6 := c.GetPostForm("date")
	ticketType, ok7 := c.GetPostForm("type")
	routerID, ok8 := c.GetPostForm("routerDetailID")

	//Check parameter
	if !ok1 || !ok2 || len(customerID) == 0 || len(transactionAmount) == 0 || !ok3 || !ok4 || len(startStop) == 0 || len(destination) == 0 || !ok5 || !ok6 || len(number) == 0 || len(date) == 0 || !ok7 || !ok8 || len(ticketType) == 0 || len(routerID) == 0 {
		paymentLog.Infof("| Pay.getPostForm | %13v | %15s | %v |", time.Since(start), c.ClientIP(), errors.New("parameter not enough"))
		c.String(406, "Please provide full parameter")
		return
	}

	//Parse number type
	number64, err := strconv.ParseInt(number, 10, 32)
	if err != nil {
		paymentLog.Infof("| Pay.parseInt | %13v | %15s | %v |", time.Since(start), c.ClientIP(), errors.New("parameter not vaild"))
		c.String(406, "Please provide full parameter")
	}
	routerID64, err := strconv.ParseInt(routerID, 10, 32)
	if err != nil {
		paymentLog.Infof("| Pay.parseInt | %13v | %15s | %v |", time.Since(start), c.ClientIP(), errors.New("parameter not vaild"))
		c.String(406, "Please provide full parameter")
	}

	number32 := int32(number64)
	royterID32 := int32(routerID64)

	//gprc client
	server := proto.NewPaymentServiceClient(connPayment)

	//Packing parameters for gRPC call
	req := proto.PaymentInfo{
		CustomerID:        customerID,
		TransactionAmount: getHeadBytesToNum(transactionAmount),
		Start:             startStop,
		Destination:       destination,
		Number:            number32,
		RouteDetailID:     royterID32,
		Date:              date,
		Type:              ticketType,
	}

	//Payment
	res, err := server.Payment(context.Background(), &req)
	if err != nil {
		paymentLog.Infof("| Pay.payment | %13v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(200, err.Error())
		return
	}

	//Check if successful, then packing resopnse
	var response paymentResponseJson
	if res.Status {
		response = paymentResponseJson{
			Status:  res.Status,
			Message: res.Reason,
		}
	} else {
		response = paymentResponseJson{
			Status:  false,
			Message: res.Reason,
		}
	}

	//Change to Json type
	resByte, err := json.Marshal(&response)
	if err != nil {
		paymentLog.Infof("| Pay.marshal | %13v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Write(resByte)

	paymentLog.Infof("| Pay | success | %13v | %15s |", time.Since(start), c.ClientIP())
}

//QueryOrder
func QueryOrder(c *gin.Context) {
	//Add header
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	userID, ok := c.GetQuery("userID")
	if !ok || len(userID) == 0 {
		locationLog.Infof("| QueryOrder.getQuery | %13v | %15s | %v |", time.Since(start), c.ClientIP(), errors.New("request did not provide user name"))
		c.String(406, "Please provide user name")
		return
	}

	//gprc client
	server := proto.NewPaymentServiceClient(connPayment)

	//Packing parameters for gRPC call
	req := proto.QueryOrderUserID{
		UserID: userID,
	}

	//Query
	res, err := server.QueryOrder(context.Background(), &req)
	if err != nil {
		paymentLog.Infof("| QueryOrder.QueryOrder | %13v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}
	if res == nil {
		paymentLog.Infof("| QueryOrder.QueryOrder | %13v | %15s | %v |", time.Since(start), c.ClientIP(), errors.New("unkonow error, res = nil"))
		c.String(500, err.Error())
		return
	}

	//Change to Json type
	resByte, err := json.Marshal(&res)
	if err != nil {
		paymentLog.Infof("| QueryOrder.marshal | %13v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Write(resByte)

	paymentLog.Infof("| QueryOrder | success | %13v | %15s |", time.Since(start), c.ClientIP())
}

//getHeadBytesToNum return float64 which at the beginning of the string
func getHeadBytesToNum(s string) float64 {
	reg := regexp.MustCompile(`^[+-]?\d*(\.)?\d*`)
	numStr := reg.FindString(s)
	num, _ := strconv.ParseFloat(numStr, 64)
	return num
}
