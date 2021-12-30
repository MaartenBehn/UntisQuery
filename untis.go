package main

import "github.com/Stroby241/UntisAPI"

var user *UntisAPI.User

func Login(username string, password string, school string, server string) {
	user = UntisAPI.NewUser(username, password, school, server)
	err := user.Login()
	if err != nil {
		user = nil
		return
	}

	initCalls()
}

func Logout() {
	if user != nil {
		return
	}
	user.Logout()
	user = nil
}

var rooms map[int]UntisAPI.Room

func initCalls() {
	if user != nil {
		return
	}
	var err error

	rooms, err = user.GetRooms()
	if err != nil {
		return
	}
}
