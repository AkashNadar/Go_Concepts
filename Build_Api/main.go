package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Models
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Database
var courses []Course

// Middlewares
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// Controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Api by Golang</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside the getAllCourse controller")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
