package entity

import (
	"strings"
)

type Meeting struct {
	m_sponsor              string
	m_participators        []string
	m_startDate, m_endDate Date
	m_title                string
}

func (m_meeting Meeting) init(t_sponsor string, t_participator []string, t_startTime Date, t_endTime Date, t_title string) {
	m_meeting.m_sponsor = t_sponsor
	m_meeting.SetParticipator(t_participator)
	m_meeting.m_startDate.CopyDate(t_startTime)
	m_meeting.m_endDate.CopyDate(t_endTime)
	m_meeting.m_title = t_title
}

func (m_meeting Meeting) CopyMeeting(t_meeting Meeting) {
	m_meeting.m_sponsor = t_meeting.m_sponsor
	m_meeting.SetParticipator(t_meeting.m_participators)
	m_meeting.m_startDate.CopyDate(t_meeting.m_startDate)
	m_meeting.m_endDate.CopyDate(t_meeting.m_endDate)
	m_meeting.m_title = t_meeting.m_title
}

func (m_meeting Meeting) GetSponsor() string {
	return m_meeting.m_sponsor
}

func (m_meeting Meeting) SetSponsor(t_sponsor string) {
	m_meeting.m_sponsor = t_sponsor
}

func (m_meeting Meeting) GetParticipator() []string {
	return m_meeting.m_participators
}

func (m_meeting Meeting) SetParticipator(t_participators []string) {
	var length = len(t_participators)
	for i := 0; i < length; i++ {
		m_meeting.m_participators[i] = t_participators[i]
	}
}

func (m_meeting Meeting) GetStartDate() Date {
	return m_meeting.m_startDate
}

func (m_meeting Meeting) SetStartDate(t_startTime Date) {
	m_meeting.m_startDate.CopyDate(t_startTime)
}

func (m_meeting Meeting) GetEndDate() Date {
	return m_meeting.m_endDate
}

func (m_meeting Meeting) SetEndDate(t_endTime Date) {
	m_meeting.m_endDate.CopyDate(t_endTime)
}

func (m_meeting Meeting) GetTitle() string {
	return m_meeting.m_title
}

func (m_meeting Meeting) SetTitle(t_title string) {
	m_meeting.m_title = t_title
}

func (m_meeting Meeting) IsParticipator(t_username string) bool {
	var i int
	for i = 0; i < len(m_meeting.m_participators); i++ {
		if strings.EqualFold(m_meeting.m_participators[i], t_username) == true {
			return true
		}
	}
	return false
}
func (m_meeting *Meeting) DeleteParticipator(t_username string) {
	var i int
	tl := len(m_meeting.m_participators)
	for i = 0; i < tl; i++ {
		if strings.EqualFold(m_meeting.m_participators[i], t_username) == true {
			m_meeting.m_participators = append(m_meeting.m_participators[:i], m_meeting.m_participators[i+1:]...)
			break
		}
	}
}
func (m_meeting *Meeting) AddParticipator(t_username string) bool {
	if strings.EqualFold(m_meeting.m_sponsor, t_username) || m_meeting.IsParticipator(t_username) {
		return false
	}
	m_meeting.m_participators = append(m_meeting.m_participators, t_username)
	return true
}
