package service

import (
	"Agenda/entity"
	"Agenda/loghelper"
	"log"
)

var curuserinfoPath = "/src/Agenda/data/curuser.txt"

var errLog *log.Logger

type User entity.User
type Meeting entity.Meeting

func init() {

	errLog = loghelper.Error

}
func UserRegister(username string, password string, email string, phone string) (bool, error) {

	user := entity.QueryUser(func(u *entity.User) bool {

		return u.GetName() == username

	})

	if len(user) == 1 {

		errLog.Println("User Register: Already exist username")

		return false, nil

	}

	entity.CreateUser(&entity.User{username, password, email, phone})

	if err := entity.Sync(); err != nil {

		return true, err

	}

	return true, nil

}
