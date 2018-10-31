package entity

import (
	"strings"
)

type Meeting struct {
	M_sponsor              string
	M_participators        []string
	M_startDate, M_endDate Date
	M_title                string
}

func (m_meeting Meeting) init(t_sponsor string, t_participator []string, t_startTime Date, t_endTime Date, t_title string) {
	m_meeting.M_sponsor = t_sponsor
	m_meeting.SetParticipator(t_participator)
	m_meeting.M_startDate.CopyDate(t_startTime)
	m_meeting.M_endDate.CopyDate(t_endTime)
	m_meeting.M_title = t_title
}

func (m_meeting Meeting) CopyMeeting(t_meeting Meeting) {
	m_meeting.M_sponsor = t_meeting.M_sponsor
	m_meeting.SetParticipator(t_meeting.M_participators)
	m_meeting.M_startDate.CopyDate(t_meeting.M_startDate)
	m_meeting.M_endDate.CopyDate(t_meeting.M_endDate)
	m_meeting.M_title = t_meeting.M_title
}

func (m_meeting Meeting) GetSponsor() string {
	return m_meeting.M_sponsor
}

func (m_meeting Meeting) SetSponsor(t_sponsor string) {
	m_meeting.M_sponsor = t_sponsor
}

func (m_meeting Meeting) GetParticipator() []string {
	return m_meeting.M_participators
}

func (m_meeting Meeting) SetParticipator(t_participators []string) {
	var length = len(t_participators)
	for i := 0; i < length; i++ {
		m_meeting.M_participators[i] = t_participators[i]
	}
}

func (m_meeting Meeting) GetStartDate() Date {
	return m_meeting.M_startDate
}

func (m_meeting Meeting) SetStartDate(t_startTime Date) {
	m_meeting.M_startDate.CopyDate(t_startTime)
}

func (m_meeting Meeting) GetEndDate() Date {
	return m_meeting.M_endDate
}

func (m_meeting Meeting) SetEndDate(t_endTime Date) {
	m_meeting.M_endDate.CopyDate(t_endTime)
}

func (m_meeting Meeting) GetTitle() string {
	return m_meeting.M_title
}

func (m_meeting Meeting) SetTitle(t_title string) {
	m_meeting.M_title = t_title
}

func (m_meeting Meeting) IsParticipator(t_username string) bool {
	var i int
	for i = 0; i < len(m_meeting.M_participators); i++ {
		if strings.EqualFold(m_meeting.M_participators[i], t_username) == true {
			return true
		}
	}
	return false
}
func (m_meeting *Meeting) DeleteParticipator(t_username string) {
	var i int
	tl := len(m_meeting.M_participators)
	for i = 0; i < tl; i++ {
		if strings.EqualFold(m_meeting.M_participators[i], t_username) == true {
			m_meeting.M_participators = append(m_meeting.M_participators[:i], m_meeting.M_participators[i+1:]...)
			break
		}
	}
}
func (m_meeting *Meeting) AddParticipator(t_username string) bool {
	if strings.EqualFold(m_meeting.M_sponsor, t_username) || m_meeting.IsParticipator(t_username) {
		return false
	}
	m_meeting.M_participators = append(m_meeting.M_participators, t_username)
	return true
}
