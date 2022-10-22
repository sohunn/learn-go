package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - should be ideally in a file
type Course struct {
	CourseID    string  `json:"courseID"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware/helpers - should be ideally in a file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - Online")
	r := mux.NewRouter()

	courses = append(courses, Course{
		CourseID:    "2",
		CourseName:  "React JS",
		CoursePrice: 299,
		Author:      &Author{Fullname: "Sohan Shashi", Website: "lco.dev"},
	})

	courses = append(courses, Course{
		CourseID:    "4",
		CourseName:  "MERN Stack",
		CoursePrice: 199,
		Author:      &Author{Fullname: "Sohan Shashi", Website: "go.dev"},
	})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/courses/update/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/delete/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

// controllers - should be ideally in a file

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Sohan's API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received on getAllCourses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received on getCourse")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses and find matching ID and return response
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received on createCourse")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid JSON. No data sent")
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Invalid JSON. No data sent inside JSON")
	}

	rand.Seed(time.Now().Unix())
	course.CourseID = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received on updateCourse")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)

			var newCourse Course
			json.NewDecoder(r.Body).Decode(&newCourse)

			newCourse.CourseID = params["id"]
			courses = append(courses, newCourse)

			json.NewEncoder(w).Encode(newCourse)
			break
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received on deleteCourse")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for i, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)

			json.NewEncoder(w).Encode("Course successfully deleted")
			break
		}
	}
}
