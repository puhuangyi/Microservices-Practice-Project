//Package handleFun implements functions to handle request

package handleFun

//Login file define login and registered service handle function.
//When request is login or registered, programme will run functions
//below.

import (
	"context"
	"encoding/json"
	"errors"
	"middle-layer/authentication"
	"middle-layer/mylog"
	"middle-layer/proto"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	loginLog  *logrus.Logger //Define login log output
	connLogin *grpc.ClientConn
)

func init() {
	//Init login log output
	loginLog = mylog.LoginLogger()

	var err error

	//Connect to login service
	connLogin, err = grpc.Dial("login-service:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		loginLog.Errorf("| grpc.Dial | %v |", err)
		return
	}
}

//registeredResponseJson store registered response data
type registeredResponseJson struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

//loginResponseJson store login response data
type loginResponseJson struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

//Login function handle login request.
//Firstly, check the parameter if correct.
//Secondly, call login function in login service with gRPC.
//Finally, check login if successful then return response.
func Login(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get parameter
	userID, ok1 := c.GetPostForm("userID")
	password, ok2 := c.GetPostForm("password")

	//Check parameter
	if !ok1 || !ok2 || len(userID) == 0 || len(password) == 0 {
		loginLog.Infof("| Login.getPostForm | %10v | %15s | %v |", time.Since(start), c.ClientIP(), errors.New("parameter not enough"))
		c.String(406, "Please provide full parameter")
		return
	}

	//Packing parameters for gRPC call
	req := proto.LoginInfo{
		UserID:   userID,
		Password: password,
	}

	//grpc client
	server := proto.NewLoginServiceClient(connLogin)
	var response loginResponseJson

	//Login (call login function in login service)
	res, err := server.Login(context.Background(), &req)
	if err != nil { //Unsuccessful
		loginLog.Infof("| Login.login | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.Writer.WriteHeader(200)
		//Packing resopnse
		response = loginResponseJson{
			Status:  false,
			Message: err.Error(),
			Token:   "",
		}
		resByte, _ := json.Marshal(&response) //Change to Json type
		c.Writer.Write(resByte)               //Send response
		return
	}

	//Packing resopnse
	response = loginResponseJson{
		Status:  res.Flag,
		Message: "login successful",
		Token:   authentication.CreateToken(userID),
	}

	//Change to Json type
	resByte, err := json.Marshal(&response)
	if err != nil {
		loginLog.Infof("| Login.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Write(resByte)

	loginLog.Infof("| Login | success | %10v | %15s |", time.Since(start), c.ClientIP())
}

//Register function responsible for the user registration.
//Firstly, check the parameter if correct.
//Secondly, call register function in login service with gRPC.
//Finally, check register if successful then return response.
func Register(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	start := time.Now() //Start time

	//Get the parameters
	usrID, ok1 := c.GetPostForm("userID")
	email, ok2 := c.GetPostForm("email")
	password, ok3 := c.GetPostForm("password")

	//Check parameters if exist
	if !ok1 || !ok2 || !ok3 || len(usrID) == 0 || len(email) == 0 || len(password) == 0 {
		loginLog.Infof("| Register.checkParameter | %10v | %15s | %v |", time.Since(start), c.ClientIP(), errors.New("parameter not enough"))
		c.String(406, "Please provide full parameter")
		return
	}

	//Check parameters if vaild
	if len(usrID) > 20 || len(email) > 20 || len(password) > 30 {
		loginLog.Infof("| Register.checkParameter | %10v | %15s | %v |", time.Since(start), c.ClientIP(), errors.New("parameter not enough"))
		c.String(406, "Parameter is invaild (too long)")
		return
	}

	//grpc client
	server := proto.NewLoginServiceClient(connLogin)
	var response registeredResponseJson //result

	//Packing parameters for gRPC call
	req := proto.UserInfo{
		UserID:   usrID,
		Email:    email,
		Password: password,
	}

	//Register
	res, err := server.Register(context.Background(), &req)
	if err != nil {
		loginLog.Infof("| Register.register | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.Writer.WriteHeader(200)
		response = registeredResponseJson{
			Status:  false,
			Message: err.Error(),
		}
		resByte, _ := json.Marshal(&response)
		c.Writer.Write(resByte)
		return
	}

	//Packing resopnse
	response = registeredResponseJson{
		Status:  res.Flag,
		Message: "user created successful",
	}

	//Change to Json type
	resByte, err := json.Marshal(&response)
	if err != nil {
		loginLog.Infof("| Register.marshal | %10v | %15s | %v |", time.Since(start), c.ClientIP(), err)
		c.String(500, err.Error())
		return
	}

	//Send response
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(resByte)

	loginLog.Infof("| Register | success | %10v | %15s |", time.Since(start), c.ClientIP())
}
