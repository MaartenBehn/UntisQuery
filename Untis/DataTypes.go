package Untis

type Teacher struct {
	id        int
	name      string
	foreName  string
	longName  string
	foreColor string
	backColor string
}

func parseTeacher(data map[string]interface{}) Teacher {
	return Teacher{
		id:        int(data["id"].(float64)),
		name:      data["name"].(string),
		foreName:  data["foreName"].(string),
		longName:  data["longName"].(string),
		foreColor: data["foreColor"].(string),
		backColor: data["backColor"].(string),
	}
}

type Student struct {
	id       int
	key      string
	name     string
	foreName string
	longName string
	gender   string
}

func parseStudent(data map[string]interface{}) Student {
	return Student{
		id:       int(data["id"].(float64)),
		name:     data["name"].(string),
		foreName: data["foreName"].(string),
		longName: data["longName"].(string),
		gender:   data["gender"].(string),
	}
}

type Class struct {
	id         int
	name       string
	longName   string
	teacherIds []int
}

func parseClass(data map[string]interface{}) Class {
	class := Class{
		id:       int(data["id"].(float64)),
		name:     data["name"].(string),
		longName: data["longName"].(string),
	}

	return class
}

type Room struct {
	id       int
	name     string
	longName string
}

func parseRoom(data map[string]interface{}) Room {
	room := Room{
		id:       int(data["id"].(float64)),
		name:     data["name"].(string),
		longName: data["longName"].(string),
	}

	return room
}

type Subject struct {
	id       int
	name     string
	longName string
}

func parseSubject(data map[string]interface{}) Subject {
	subject := Subject{
		id:       int(data["id"].(float64)),
		name:     data["name"].(string),
		longName: data["longName"].(string),
	}

	return subject
}

type Departments struct {
	id       int
	name     string
	longName string
}

func parseDepartments(data map[string]interface{}) Departments {
	departments := Departments{
		id:       int(data["id"].(float64)),
		name:     data["name"].(string),
		longName: data["longName"].(string),
	}

	return departments
}

type Period struct {
	id        int
	date      int
	startTime int
	endTime   int
	lstype    string
	code      string
	info      string
	subText   string
	lstext    string
	lsnumber  int

	classIds   []int
	teacherIds []int
	subjectIds []int
	roomIds    []int
}

func parsePeriod(data map[string]interface{}) Period {
	period := Period{
		id:         int(data["id"].(float64)),
		date:       int(data["date"].(float64)),
		startTime:  int(data["startTime"].(float64)),
		endTime:    int(data["endTime"].(float64)),
		lstype:     data["lstype"].(string),
		code:       data["code"].(string),
		info:       data["info"].(string),
		subText:    data["subText"].(string),
		lstext:     data["lstext"].(string),
		lsnumber:   int(data["lsnumber"].(float64)),
		classIds:   nil,
		teacherIds: nil,
		subjectIds: nil,
		roomIds:    nil,
	}
	return period
}
