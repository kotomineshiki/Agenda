package entity

import (
	"Agenda/deepcopy"
	"Agenda/loghelper"
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

type UserFilter func(*User) bool
type MeetingFilter func(*Meeting) bool

var userinfoPath = "/src/Agenda/data/userinfo"
var metinfoPath = "/src/Agenda/data/meetinginfo"
var curUserPath = "/src/Agenda/data/curUser.txt"
var curUserName *string
var dirty bool      //标记脏位
var userData []User //用来存用户信息
var meetingData []Meeting
var errLog *log.Logger

func init() {
	errLog = loghelper.Error
	dirty = false

	userinfoPath = filepath.Join(loghelper.GoPath, userinfoPath)

	metinfoPath = filepath.Join(loghelper.GoPath, metinfoPath)

	curUserPath = filepath.Join(loghelper.GoPath, curUserPath)

}

func Logout() error {

	curUserName = nil

	return Sync()

}
func Sync() error {
	if err := writeToFile(); err != nil {
		errLog.Println("writeToFile fail:", err)
		return err
	}
	return nil
}

func CreateUser(tocreate *User) {
	userData = append(userData, deepcopy.Copy(*tocreate).(User))
	dirty = true

}
func writeToFile() error {
	var e []error
	if err := writeString(curUserPath, curUserName); err != nil {
		e = append(e, err)
	}
	if dirty {
		if err := writeJSON(userinfoPath, userData); err != nil {
			e = append(e, err)
		}
		if err := writeJSON(metinfoPath, meetingData); err != nil {
			e = append(e, err)
		}
	}
	if len(e) == 0 {
		return nil
	}
	er := e[0]
	for i := 1; i < len(e); i++ {
		er = errors.New(er.Error() + e[i].Error())
	}
	return er

}
func readFromFile() error {
	var e []error
	str, err1 := readString(curUserPath)
	if err1 != nil {
		e = append(e, err1)
	}
	curUserName = str
	if err := readUser(); err != nil {
		e = append(e, err)
	}
	if err := readMeeting(); err != nil {
		e = append(e, err)
	}
	if len(e) == 0 {
		return nil
	}
	er := e[0]
	for i := 1; i < len(e); i++ {
		er = errors.New(er.Error() + e[i].Error())
	}
	return er
}
func writeString(path string, data *string) error {
	file, err := os.Create(path)
	if err != nil {
		loghelper.Error.Println("Create file error:", path)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if data != nil {
		if _, err := writer.WriteString(*data); err != nil {
			loghelper.Error.Println("Write file fail:", path)
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		loghelper.Error.Println("Flush file fail:", path)
		return err
	}
	return nil
}
func readString(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		loghelper.Error.Println("Open file error:", path)
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		loghelper.Error.Println("Read file fail:", path)
		return nil, err
	}
	return &str, nil
}
func writeJSON(fpath string, data interface{}) error {
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	if err := enc.Encode(&data); err != nil {
		errLog.Println("writeJSON:", err)
		return err
	}
	return nil
}
func readUser() error {
	file, err := os.Open(userinfoPath)
	if err != nil {
		errLog.Println("Open File Fail:", userinfoPath, err)
		return err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	switch err := dec.Decode(&userData); err {
	case nil, io.EOF:
		return nil
	default:
		errLog.Println("Decode User Fail:", err)
		return err
	}
}
func readMeeting() error {
	file, err := os.Open(metinfoPath)
	if err != nil {
		errLog.Println("Open File Fail:", metinfoPath, err)
		return err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	switch err := dec.Decode(&meetingData); err {
	case nil, io.EOF:
		return nil
	default:
		errLog.Println("Decode Met Fail:", err)
		return err
	}
}

func QueryUser(filter UserFilter) []User {
	var user []User
	for _, v := range userData {
		if filter(&v) {
			user = append(user, v)
		}
	}
	return user
}
