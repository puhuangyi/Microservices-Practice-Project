package authentication

import (
	"errors"
	"fmt"
	"io/ioutil"
	"middle-layer/mylog"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	signingKey = "Team22"
	redisCon   redisConfig
	RedisPool  *redis.Pool
	authLog    *logrus.Logger
)

func init() {
	//Create log
	authLog = mylog.AuthLogger()
	authLog.Info("Authentication log is created")

	//Read the redis configuration file
	ymlFile, err := ioutil.ReadFile("./config/redis.yml")
	if err != nil {
		authLog.Error(err)
	}

	//Save data in structure
	err = yaml.Unmarshal(ymlFile, &redisCon)
	if err != nil {
		authLog.Error(err)
	}

	//Set redis pool
	RedisPool = &redis.Pool{
		MaxIdle:     redisCon.MaxIdle,
		MaxActive:   redisCon.MaxActive,
		IdleTimeout: redisCon.IdleTimeout * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", redisCon.Addr) },
	}
}

// redisConfig store redis configuration
type redisConfig struct {
	Addr        string        `yaml:"Addr"`
	MaxIdle     int           `yaml:"MaxIdle"`
	MaxActive   int           `yaml:"MaxActive"`
	IdleTimeout time.Duration `yaml:"IdleTimeout"`
}

// myClaim is my custom JWT payload structure
type myClaim struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func CreateToken(userID string) string {
	claims := &myClaim{
		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signingKey))
	//println(signingKey)
	if err != nil {
		authLog.Infof("| SignedString | %v |", err)
		return ""
	}

	conn := RedisPool.Get()
	defer conn.Close()

	_, err = conn.Do("setex", ss, 3600*24*2, userID)
	if err != nil {
		authLog.Infof("| SignedString | %v |", err)
		return ""
	}

	authLog.Infof("| token create | userID:%s |", ss)
	return ss
}

// ParseToken returns integer value representing whether tokenStr
// was successfully parsed. Return "0" mean success, return "1" mean
// token is unvalid, return "2" mean redis connect is unavailable.
// Successfully parsed mean that token is a valid token.
// This function first uses JWT to parse the token to get the second
// part of JWT, that is, the Payload part, and then connects to redis
// to check whether the value corresponding to the token is the same
// as the UserID of the Payload part. If any part of the above steps
// is inconsistent, the token is invalid.
func ParseToken(tokenStr string) (flag int, err error) {
	//Pasrse with JWT
	token, err := jwt.ParseWithClaims(tokenStr, &myClaim{}, func(*jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	//Not JWT structure
	if err != nil {
		return 1, err
	}

	//Parse Payload part
	if claims, ok := token.Claims.(*myClaim); ok {

		//Connect to redis with redis pool
		conn := RedisPool.Get()
		defer conn.Close()

		//Check connect
		err := conn.Err()
		if err != nil {
			info := fmt.Sprintf("%+v", RedisPool.Stats())
			return 2, errors.New("redis pool unavailable, pool info:" + info)
		}

		//Query the corresponding value of the token in the redis
		s, err := redis.String(conn.Do("GET", tokenStr))
		//Value not exist
		if err != nil {
			return 1, err
		}

		//Contrast
		if s == claims.UserID {
			return 0, nil
		}

		return 1, errors.New("the parsed token is different from the database")

	}

	return 1, errors.New("struct conversion error")
}

// Auth returns a function(middleware) that authenticates a request.
// This function will parse token, if token is empty or unvalid, the
// server will refuse handle request.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

		//Get token from request header
		token := c.GetHeader("Authorization")

		//Check if has token
		if len(token) == 0 {
			authLog.Infof("| %15s | false | Authorization not value |", c.ClientIP())
			c.String(401, "Please provide token")
			c.Abort()
			return
		}

		//Parse token
		flag, err := ParseToken(token)

		//Token is right
		if flag == 0 {
			authLog.Infof("| %15s | true |", c.ClientIP()) //log output

			//Handle request
			c.Next()
			return
		}

		//Not handle, directly return response
		c.Abort()

		if flag == 1 {
			c.String(401, "Please provide valid token")
			authLog.Infof("| %15s | false | %s |", c.ClientIP(), err) //log output
			return
		}
		if flag == 2 {
			c.String(503, "redis pool busy")
			authLog.Warnf("| %15s | false | %s |", c.ClientIP(), err) //log output
			return
		}

	}
}
