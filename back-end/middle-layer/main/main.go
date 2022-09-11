package main

import (
	"io"
	"middle-layer/authentication"
	"middle-layer/handleFun"
	"middle-layer/mylog"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {

	//New gin
	engine := gin.New()

	//Set output
	a := io.MultiWriter(os.Stdout, mylog.GinLogger())
	gin.DefaultWriter = a

	//Close
	defer authentication.RedisPool.Close()
	defer handleFun.CloseResource()

	//Use log middleware
	engine.Use(mylog.ServerLogger())

	//Test
	engine.GET("/test", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("hello"))
	})

	//Login service
	{
		engine.POST("/login", handleFun.Login)
		engine.POST("/registered", handleFun.Register)
	}

	//Query service
	queryRouter := engine.Group("/query")
	{
		queryRouter.Use(authentication.Auth())
		queryRouter.GET("/location/:busName", handleFun.QueryLocation)
		queryRouter.GET("/order", handleFun.QueryOrder)

		queryRouter.GET("/fare/castle", handleFun.QueryCastleFare)
		queryRouter.GET("/fare/bus/:busName", handleFun.QueryBusFare)
		queryRouter.GET("/fare/train/:trainName", handleFun.QueryTrainFare)

		queryRouter.GET("/info/castle", handleFun.QueryCastleInfo)
		queryRouter.GET("/info/startTime", handleFun.QueryStartTime)

		queryRouter.GET("/route", handleFun.QueryRouteToCastle)
		queryRouter.GET("/routePlus", handleFun.QueryRouteAllInfo)
	}

	//Payment service
	paymentRouter := engine.Group("/payment")
	{
		paymentRouter.Use(authentication.Auth())
		paymentRouter.POST("/horsepay", handleFun.Pay)
	}

	//Other path show 404
	{
		engine.POST("/:other", func(c *gin.Context) {
			c.String(404, "404 page not found")
		})

		engine.GET("/:other", func(c *gin.Context) {
			c.String(404, "404 page not found")
		})
	}

	//Start
	engine.Run(":8080")
}
