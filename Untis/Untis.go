package Untis

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	username string
	password string
	school   string
	server   string

	sessionId  string
	personType float64
	personId   float64
	klasseId   float64
}

var timetables [][]Period

func NewUser(username string, password string, school string, server string) *User {

	schoolParts := splitAny(school, " +_")
	school = ""
	for i, part := range schoolParts {
		if i > 0 {
			school += "+"
		}
		school += part
	}

	return &User{
		username: username,
		password: password,
		school:   school,
		server:   server,
	}
}

func (u *User) request(mehtode string, jsonParam map[string]interface{}) interface{} {

	url := u.server + "/WebUntis/jsonrpc.do" + "?school=" + u.school

	postBody, _ := json.Marshal(map[string]interface{}{
		"id":      "0",
		"method":  mehtode,
		"params":  jsonParam,
		"jsonrpc": "2.0",
	})
	responseBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, responseBody)
	checkError(err)

	if u.sessionId != "" {
		req.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: u.sessionId})
	}

	res, err := client.Do(req)
	checkError(err)

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)

	var response interface{}
	err = json.Unmarshal(body, &response)
	checkError(err)

	result := response.(map[string]interface{})
	for key, value := range result {
		if key == "error" {
			log.Print(value)
			return map[string]interface{}{}
		}
	}
	log.Println()

	return result["result"]
}

func (u *User) Login() {
	postBody := map[string]interface{}{
		"user":     u.username,
		"password": u.password,
		"client":   "UntisQuerry",
	}

	result := u.request("authenticate", postBody).(map[string]interface{})
	u.personId = result["personId"].(float64)
	u.klasseId = result["klasseId"].(float64)
	u.personType = result["personType"].(float64)
	u.sessionId = result["sessionId"].(string)
}
func (u *User) Logout() {
	postBody := map[string]interface{}{}

	u.request("logout", postBody)
}

func (u *User) GetTeachers() []Teacher {
	postBody := map[string]interface{}{}

	result := u.request("getTeachers", postBody)
	var teachers []Teacher
	for _, t := range result.([]interface{}) {
		teacher := parseTeacher(t.(map[string]interface{}))
		teachers = append(teachers, teacher)
	}
	return teachers
}
func (u *User) GetStudents() []Student {
	postBody := map[string]interface{}{}

	result := u.request("getStudents", postBody)
	var students []Student
	for _, t := range result.([]interface{}) {
		student := parseStudent(t.(map[string]interface{}))
		students = append(students, student)
	}
	return students
}
func (u *User) GetClasses() []Class {
	postBody := map[string]interface{}{}

	result := u.request("getKlassen", postBody)
	var classes []Class
	for _, t := range result.([]interface{}) {
		student := parseClass(t.(map[string]interface{}))
		classes = append(classes, student)
	}
	return classes
}
func (u *User) GetRooms() []Room {
	postBody := map[string]interface{}{}

	result := u.request("getRooms", postBody)
	var rooms []Room
	for _, t := range result.([]interface{}) {
		room := parseRoom(t.(map[string]interface{}))
		rooms = append(rooms, room)
	}
	return rooms
}
func (u *User) GetSubjects() []Subject {
	postBody := map[string]interface{}{}

	result := u.request("getSubjects", postBody)
	var subjects []Subject
	for _, t := range result.([]interface{}) {
		subject := parseSubject(t.(map[string]interface{}))
		subjects = append(subjects, subject)
	}
	return subjects
}

func (u *User) GetTimetable(id int, idtype int /*, startDate int , endDate int */) (periods []Period) {
	postBody := map[string]interface{}{
		"options": map[string]interface{}{
			"element": map[string]interface{}{"id": id, "type": idtype, "keyType": "id"},

			"showStudentgroup": true,
			"showLsText":       true,
			"showLsNumber":     true,
			"showInfo":         true,
			"showBooking":      true,
			"showSubstText":    true,

			// "startDate": startDate,
			// "endDate": endDate,

			"klasseFields":  []interface{}{"id", "name", "longname", "externalkey"},
			"roomFields":    []interface{}{"id", "name", "longname", "externalkey"},
			"subjectFields": []interface{}{"id", "name", "longname", "externalkey"},
			"teacherFields": []interface{}{"id", "name", "longname", "externalkey"},
		},
	}

	timetable := u.request("getTimetable", postBody).([]interface{})
	for _, data := range timetable {
		period := parsePeriod(data.(map[string]interface{}))
		periods = append(periods, period)
	}
	return periods
}

func (u *User) GetData() {
	rooms := u.GetRooms()
	for _, room := range rooms {
		timetable := u.GetTimetable(room.id, 4)
		timetables = append(timetables, timetable)
	}
}

/*
func (u *User) QuerryTeacher(firstname string, lastname string) {
	postBody := map[string]interface{}{
		"type": "2",
		"fn":   firstname,
		"sn":   lastname,
		"dob":  "0",
	}

	var teacherId = u.request("getPersonId", postBody).(float64)
	log.Printf("TeacherId: %.0f\n", teacherId)

	for _, timetable := range timetables {
		for _, s := range timetable {
			subject := s.(map[string]interface{})
			for key, value := range subject {
				if key == "te" {
					datas := value.([]interface{})
					for _, d := range datas {
						data := d.(map[string]interface{})
						if data["id"] == teacherId {

							rooms := subject["ro"].([]interface{})
							roomName := ""
							for _, value := range rooms {
								attr := value.(map[string]interface{})
								for key, value := range attr {
									if key == "name" {
										roomName = value.(string)
									}
								}
							}

							klass := subject["kl"].([]interface{})
							klassName := ""
							for _, value := range klass {
								attr := value.(map[string]interface{})
								for key, value := range attr {
									if key == "name" {
										klassName = value.(string)
									}
								}
							}

							log.Printf("Found: %s %s %s %.0f %.0f", data["name"], roomName, klassName, subject["date"], subject["startTime"])
						}
					}
				}
			}
		}
	}
}
*/
