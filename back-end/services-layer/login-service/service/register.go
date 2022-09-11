package service

import (
	"context"
	"errors"
	"login/myInit"
	"login/mylog"
	"login/proto"
	"time"
)

type LoginClient struct {
}

func (s *LoginClient) Register(ctx context.Context, userInfo *proto.UserInfo) (*proto.ResUserInfo, error) {
	flag, err := register(userInfo)
	return &proto.ResUserInfo{
		Flag: flag,
	}, err
}

func register(userInfo *proto.UserInfo) (bool, error) {
	start := time.Now()

	rows, err := myInit.Db.Query("select userID, email from user where userID = ? or email = ?", userInfo.UserID, userInfo.Email)
	if err != nil {
		mylog.LogClient.Infof("| query | %13v | %v |", time.Since(start), err)
		return false, err
	}

	defer rows.Close()

	var userID string
	var email string

	for rows.Next() {
		err := rows.Scan(&userID, &email)
		if err != nil {
			mylog.LogClient.Errorf("| scan | %13v | %v |", time.Since(start), err)
			return false, err
		}
	}

	if userID != "" || email != "" {
		return false, errors.New("user have been create before")
	}

	insertSql, err := myInit.Db.Prepare("insert user (userID, email , password) values (?,?,?)")
	if err != nil {
		mylog.LogClient.Errorf("| prepare | %13v | %v |", time.Since(start), err)
		return false, err
	}

	result, err := insertSql.Exec(userInfo.UserID, userInfo.Email, userInfo.Password)
	if err != nil {
		mylog.LogClient.Errorf("| exec | %13v | %v |", time.Since(start), err)
		return false, err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		mylog.LogClient.Infof("| lastInsertId | %13v | %v |", time.Since(start), err)
		return false, err
	}

	mylog.LogClient.Infof("| %13v | %d |", time.Since(start), insertID)
	return true, nil
}
