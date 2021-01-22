package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Student struct {
	Number int    `json:"student_number"`
	Name   string `json:"name"`
}
type Class struct {
	Number   int       `json:"class_number"`
	Students []Student `json:"students"`
}

func GetStudentNumberHandler(c echo.Context) error {
	classNum, err := strconv.Atoi(c.Param("class"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string {"error": "Student Not Found"})
	}
	studentNum, err := strconv.Atoi(c.Param("studentNumber"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string {"error": "Student Not Found"})
	}

	var classes []Class

	_ = json.Unmarshal([]byte(`[
	{"class_number": 1, "students": [
		{"student_number": 1, "name": "Humming"},
		{"student_number": 2, "name": "masutech16"},
		{"student_number": 3, "name": "ninja"}
	]},
	{"class_number": 2, "students": [
		{"student_number": 1, "name": "hukuda222"},
		{"student_number": 2, "name": "takashi_trap"},
		{"student_number": 3, "name": "nagatech"},
		{"student_number": 4, "name": "whiteonion"}
	]},
	{"class_number": 3, "students": [
		{"student_number": 1, "name": "yamada"},
		{"student_number": 2, "name": "tubotu"},
		{"student_number": 3, "name": "tsukatomo"}
	]},
	{"class_number": 4, "students": [
		{"student_number": 1, "name": "g2"},
		{"student_number": 2, "name": "hatasa-y"}
	]}
	]`), &classes)

	var result Student
	for _, v := range classes {
		if v.Number == classNum {
			for _, u := range v.Students {
				if u.Number == studentNum {
					result = u
					return c.JSON(http.StatusOK, result)
				}
			}
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]string {"error": "Student Not Found"})
}
