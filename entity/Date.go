package entity

import (
    "fmt"
    "strconv"
    "errors"
)

type Date struct{
    m_year,m_month,m_day,m_hour,m_minute int
}

func (m_date Date) init(t_year, t_month, t_day, t_hour, t_minute int)  {
    m_date.m_year = t_year
    m_date.m_month = t_month
    m_date.m_day = t_day
    m_date.m_hour = t_hour
    m_date.m_minute = t_minute
}

func (m_date Date) CopyDate (t_date Date) Date {
    m_date.SetYear(t_date.GetYear())
    m_date.SetMonth(t_date.GetMonth())
    m_date.SetDay(t_date.GetDay())
    m_date.SetHour(t_date.GetHour())
    m_date.SetMinute(t_date.GetMinute())
    return m_date
}

func StringToDate(dateString string) (Date, error) {
    var result Date
    if (len(dateString) != 16) {
        return result, errors.New("wrong")
    }
    var count int = 0 
    for count < len(dateString) {
        if (count==4 || count==7) && dateString[count]!='-' {
            return result, errors.New("wrong")
        } else if count==10 && dateString[count]!='/'{
            return result, errors.New("wrong")
        } else if count==13 && dateString[count]!=':'{
            return result, errors.New("wrong")
        } else if (count!=4&&count!=7&&count!=10&&count!=13)&&(dateString[count]>'9'||dateString[count]<'0'){
            return result, errors.New("wrong")
        }
        count++
    }