package entity

import (
	"errors"
	"fmt"
	"strconv"
)

type Date struct {
	m_year, m_month, m_day, m_hour, m_minute int
}

func (m_date Date) init(t_year, t_month, t_day, t_hour, t_minute int) {
	m_date.m_year = t_year
	m_date.m_month = t_month
	m_date.m_day = t_day
	m_date.m_hour = t_hour
	m_date.m_minute = t_minute
}

func (m_date Date) CopyDate(t_date Date) Date {
	m_date.SetYear(t_date.GetYear())
	m_date.SetMonth(t_date.GetMonth())
	m_date.SetDay(t_date.GetDay())
	m_date.SetHour(t_date.GetHour())
	m_date.SetMinute(t_date.GetMinute())
	return m_date
}

func StringToDate(t_dateString string) (Date, error) {
	var result Date
	if len(t_dateString) != 16 {
		return result, errors.New("wrong")
	}
	var count int = 0
	for count < len(t_dateString) {
		if (count == 4 || count == 7) && t_dateString[count] != '-' {
			return result, errors.New("wrong")
		} else if count == 10 && t_dateString[count] != '/' {
			return result, errors.New("wrong")
		} else if count == 13 && t_dateString[count] != ':' {
			return result, errors.New("wrong")
		} else if (count != 4 && count != 7 && count != 10 && count != 13) && (t_dateString[count] > '9' || t_dateString[count] < '0') {
			return result, errors.New("wrong")
		}
		count++
	}
	var err error = nil
	result.m_year, err = strconv.Atoi(t_dateString[0:4])
	result.m_month, err = strconv.Atoi(t_dateString[5:7])
	result.m_day, err = strconv.Atoi(t_dateString[8:10])
	result.m_hour, err = strconv.Atoi(t_dateString[11:13])
	result.m_minute, err = strconv.Atoi(t_dateString[14:])
	return result, err
}

func (m_date Date) GetYear() int {
	return m_date.m_year
}
func (m_date Date) SetYear(t_year int) {
	m_date.m_year = t_year
}
func (m_date Date) GetMonth() int {
	return m_date.m_month
}
func (m_date Date) SetMonth(t_month int) {
	m_date.m_month = t_month
}
func (m_date Date) GetDay() int {
	return m_date.m_day
}
func (m_date Date) SetDay(t_day int) {
	m_date.m_day = t_day
}
func (m_date Date) GetHour() int {
	return m_date.m_hour
}
func (m_date Date) SetHour(t_hour int) {
	m_date.m_hour = t_hour
}
func (m_date Date) GetMinute() int {
	return m_date.m_minute
}
func (m_date Date) SetMinute(t_minute int) {
	m_date.m_minute = t_minute
}

func IsValid(t_date Date) bool {
	if t_date.GetYear() > 9999 || t_date.GetYear() < 1000 {
		return false
	}
	month_day := [12]int{21, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var current_year int = t_date.GetYear()
	var current_month int = t_date.GetMonth()
	var current_day int = t_date.GetDay()
	var current_hour int = t_date.GetHour()
	var current_minute int = t_date.GetMinute()

	if current_year%400 == 0 || (current_year%4 == 0 && current_year%100 != 0) {
		month_day[1] = 29
	}

	if current_month > 12 || current_month <= 0 ||
		current_day > month_day[current_month-1] || current_day <= 0 ||
		current_hour > 23 || current_hour < 0 ||
		current_minute > 59 || current_minute < 0 {
		return false
	}
	return true
}

func DateToString(t_date Date) (string, error) {
	var dateString string = ""
	var initTime string = "0000-00-00/00:00"
	if !IsValid(t_date) {
		return initTime, nil
	}
	dateString = fmt.Sprintf("%04d-%02d-%02d/%02d:%02d", t_date.GetYear(), t_date.GetMonth(), t_date.GetDay(), t_date.GetHour(), t_date.GetMinute())
	if dateString != "" {
		return dateString, nil
	}
	return dateString, errors.New("wrong")
}

func (m_date Date) IsSameDate(t_date Date) bool {
	return (t_date.GetYear() == m_date.GetYear() &&
		t_date.GetMonth() == m_date.GetMonth() &&
		t_date.GetDay() == m_date.GetDay() &&
		t_date.GetHour() == m_date.GetHour() &&
		t_date.GetMinute() == m_date.GetMinute())
}

func (m_date Date) MoreThan(t_date Date) bool {
	if m_date.m_year > t_date.GetYear() {
		return true
	} else if m_date.m_year == t_date.GetYear() {
		if m_date.m_month > t_date.GetMonth() {
			return true
		} else if m_date.m_month == t_date.GetMonth() {
			if m_date.m_day > t_date.GetDay() {
				return true
			} else if m_date.m_day == t_date.GetDay() {
				if m_date.m_hour > t_date.GetHour() {
					return true
				} else if m_date.m_hour == t_date.GetHour() {
					if m_date.m_minute > t_date.GetMinute() {
						return true
					}
				}
			}
		}
	}
	return false
}

func (m_date Date) LessThan(t_date Date) bool {
	if m_date.IsSameDate(t_date) == false && m_date.MoreThan(t_date) == false {
		return true
	}
	return false
}

func (m_date Date) GreateOrEqual(t_date Date) bool {
	return m_date.IsSameDate(t_date) || m_date.MoreThan(t_date)
}

func (m_date Date) LessOrEqual(t_date Date) bool {
	return !m_date.MoreThan(t_date)
}
