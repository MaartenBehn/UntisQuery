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

	sessionId  string
	personType float64
	personId   float64
	klasseId   float64
}

var timetables [][]interface{}

func NewUser(username string, password string, school string) *User {

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
	}
}

func request(mehtode string, jsonParam map[string]interface{}, seesionId string, urlParam ...string) interface{} {
	url := "https://tipo.webuntis.com/WebUntis/jsonrpc.do"

	if len(urlParam) > 0 {
		url += "?"
	}
	for i, para := range urlParam {
		if i > 0 {
			url += "&"
		}
		url += para
	}

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

	if seesionId != "" {
		req.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: seesionId})
	}

	res, err := client.Do(req)
	checkError(err)

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)

	var response interface{}
	err = json.Unmarshal(body, &response)
	checkError(err)

	return response
}

func (u *User) Login() {
	postBody := map[string]interface{}{
		"user":     u.username,
		"password": u.password,
		"client":   "UntisQuerry",
	}

	response := request("authenticate", postBody, "", "school="+u.school).(map[string]interface{})

	result := response["result"].(map[string]interface{})
	u.personId = result["personId"].(float64)
	u.klasseId = result["klasseId"].(float64)
	u.personType = result["personType"].(float64)
	u.sessionId = result["sessionId"].(string)
}

func (u *User) GetData() {
	postBody := map[string]interface{}{}
	rooms := request("getRooms", postBody, u.sessionId, "school="+u.school).(map[string]interface{})["result"].([]interface{})

	for _, r := range rooms {
		room := r.(map[string]interface{})

		postBody := map[string]interface{}{
			"options": map[string]interface{}{
				"element":          map[string]interface{}{"id": room["id"], "type": 4, "keyType": "id"},
				"showStudentgroup": true,
				"showLsText":       true,
				"showLsNumber":     true,
				"showInfo":         true,
				"roomFields":       []interface{}{"id", "name"},
				"klasseFields":     []interface{}{"id", "name", "longname", "externalkey"},
				"teacherFields":    []interface{}{"id", "name"},
			},
		}

		response := request("getTimetable", postBody, u.sessionId, "school="+u.school).(map[string]interface{})
		timetable := response["result"].([]interface{})
		timetables = append(timetables, timetable)
	}
}

func (u *User) QuerryTeacher(firstname string, lastname string) {
	postBody := map[string]interface{}{
		"type": "2",
		"fn":   firstname,
		"sn":   lastname,
		"dob":  "0",
	}
	response := request("getPersonId", postBody, u.sessionId, "school="+u.school).(map[string]interface{})
	var teacherId = response["result"].(float64)
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

func (u *User) Logout() {
	postBody := map[string]interface{}{}

	request("logout", postBody, u.sessionId, "school="+u.school)
}
