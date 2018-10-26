package entity

import (

	"fmt"

)


func StartAgenda() bool {
	ReadFromFile()
	ReadCurrentUser()
	if CurrentUser.Name == "" {
		return false
	}
	return true
}
func QuitAgenda() {
	writeToFile()
	writeCurrentUser()
}