package entity

type User struct {
	M_name, M_password, M_email, M_phone string
}

func (m_User User) init(t_userName, t_userPassword, t_userEmail, t_userPhone string) {
	m_User.M_name = t_userName
	m_User.M_password = t_userPassword
	m_User.M_email = t_userEmail
	m_User.M_phone = t_userPhone
}

func (m_User User) CopyUser(t_user User) {
	m_User.M_name = t_user.M_name
	m_User.M_password = t_user.M_password
	m_User.M_email = t_user.M_email
	m_User.M_phone = t_user.M_phone
}

func (m_User User) GetName() string {
	return m_User.M_name
}
func (m_User User) SetName(t_name string) {
	m_User.M_name = t_name
}
func (m_User User) GetPassword() string {
	return m_User.M_password
}
func (m_User User) SetPassword(t_password string) {
	m_User.M_password = t_password
}
func (m_User User) GetEmail() string {
	return m_User.M_email
}
func (m_User User) SetEmail(t_email string) {
	m_User.M_email = t_email
}
func (m_User User) GetPhone() string {
	return m_User.M_phone
}

func (m_User User) SetPhone(t_phone string) {
	m_User.M_phone = t_phone
}
