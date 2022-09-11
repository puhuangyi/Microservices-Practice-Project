package service

import (
	"context"
	"errors"
	"login/myInit"
	"login/mylog"
	"login/proto"
	"time"
)

func (s *LoginClient) Login(ctx context.Context, logInfo *proto.LoginInfo) (*proto.ResLoginInfo, error) {
	//println(logInfo.UserID, logInfo.Password)
	flag, err := login(logInfo)
	return &proto.ResLoginInfo{
		Flag: flag,
	}, err
}

func login(logInfo *proto.LoginInfo) (bool, error) {
	start := time.Now()

	rows, err := myInit.Db.Query("select password from user where userID = ?", &logInfo.UserID)
	if err != nil {
		mylog.LogClient.Infof("| login:query | %13v | %v |", time.Since(start), err)
		return false, err
	}
	defer rows.Close()

	var password string

	for rows.Next() {
		err := rows.Scan(&password)
		if err != nil {
			mylog.LogClient.Errorf("| login:scan | %13v | %v |", time.Since(start), err)
			return false, err
		}
	}

	if password == "" {
		mylog.LogClient.Errorf("| login:ckeckEmailIfNil | %13v | %v |", time.Since(start), err)
		return false, errors.New("user have not create")
	}

	if logInfo.Password != password {
		mylog.LogClient.Errorf("| login:ckeckEmailIfEqual | %13v | %v |", time.Since(start), err)
		return false, errors.New("password is incorrect")
	}

	mylog.LogClient.Infof("| %13v | %s login success |", time.Since(start), logInfo.UserID)
	return true, nil
}
