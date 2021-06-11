package main

import (
	"UntisQuerry/Untis"
)

func main() {
	user := Untis.NewUser("maarten8", "behn500", "TBZ Mitte Bremen")
	user.Login()
	user.GetData()
	user.QuerryTeacher("Daniel", "Dibbern")
	user.QuerryTeacher("Jan", "Benje")
	user.QuerryTeacher("Julian", "Scheichel")
	user.Logout()
}
