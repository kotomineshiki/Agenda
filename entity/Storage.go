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

	readFromFile()
}

func Logout() error {
	if *curUserName != "" {
		curUserName = nil
		return Sync()
	} else {
		return errors.New("No curUserName!")
	}

}
func Sync() error {
	if err := writeToFile(); err != nil {
		errLog.Println("writeToFile fail:", err)
		return err
	}
	return nil
}
func writeToFile() error {
	var e []error
	if err := writeString(curUserPath, curUserName); err != nil {
		e = append(e, err)
	}
	if dirty {
		if err := writeUser(); err != nil {
			e = append(e, err)
		}
		if err := writeMeeting(); err != nil {
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
func writeUser() error {
	file, err := os.Create(userinfoPath)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	if err := enc.Encode(&userData); err != nil {
		errLog.Println("writeJSON:", err)
		return err
	}
	return nil
}
func writeMeeting() error {
	file, err := os.Create(metinfoPath)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	if err := enc.Encode(&meetingData); err != nil {
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

func CreateUser(tocreate *User) {
	userData = append(userData, deepcopy.Copy(*tocreate).(User))
	dirty = true

}

func QueryUser(filter UserFilter) []User {
	var user []User
	for _, it := range userData {
		if filter(&it) {
			user = append(user, it)
		}
	}
	return user
}

func UpdateUser(filter UserFilter, switcher func(*User)) int {
	count := 0
	for i := 0; i < len(userData); i++ {
		if temp := &userData[i]; filter(temp) {
			switcher(temp)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

func DeleteUser(filter UserFilter) int {
	count := 0
	u_l := len(userData)
	for i := 0; i < len(userData); {
		if filter(&userData[i]) {
			u_l--
			userData[i] = userData[u_l]
			userData = userData[:u_l]
			count++
		} else {
			i++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}
func CreateMeeting(v *Meeting) {
	meetingData = append(meetingData, deepcopy.Copy(*v).(Meeting))
	dirty = true
}
func QueryMeeting(filter MeetingFilter) []Meeting {
	var met []Meeting
	for _, v := range meetingData {
		if filter(&v) {
			met = append(met, v)
		}
	}
	return met
}
func UpdateMeeting(filter MeetingFilter, switcher func(*Meeting)) int {
	count := 0
	for i := 0; i < len(meetingData); i++ {
		if temp := &meetingData[i]; filter(temp) {
			switcher(temp)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}
func DeleteMeeting(filter MeetingFilter) int {
	count := 0
	m_l := len(meetingData)
	for i := 0; i < len(meetingData); {
		if filter(&meetingData[i]) {
			m_l--
			meetingData[i] = meetingData[m_l]
			meetingData = meetingData[:m_l]
			count++
		} else {
			i++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}
func GetCurUser() (User, error) {
	if curUserName == nil {
		return User{}, errors.New("Current user does not exist")
	}
	for _, v := range userData {
		if v.M_name == *curUserName {
			return v, nil
		}
	}
	return User{}, errors.New("Current user does not exist")
}
func SetCurUser(u *User) {
	if u == nil {
		curUserName = nil
		return
	}
	if curUserName == nil {
		p := u.M_name
		curUserName = &p
	} else {
		*curUserName = u.M_name
	}
}
