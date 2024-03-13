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

//Model for course - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}


// fake DB

var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}


func main()  {
	fmt.Println("API - LCO")
	r := mux.NewRouter()

	// Seeding

	courses = append(courses, Course{
		CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, 
		Author: &Author{Fullname: "Divya", Website: "lco"}})
	courses = append(courses, Course{
		CourseId: "3", CourseName: "JS", CoursePrice: 199, 
		Author: &Author{Fullname: "Satya", Website: "lco"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// listen
	log.Fatal(http.ListenAndServe(":4000", r))
}
// controllers - file

// server home route

func serveHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>Welcome to API by Go</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) 
}

func getOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// fmt.Printf("Type of params is %T\n", params)

	// loop through courese, find matching id and return response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	
	// what about - {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// TODO: Check if title is duplicate
	// loop, title matches with course.courseName


	// generate unique id, string
	// append course into course

	// rand.Seed(time.Now().UnixNano())

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	course.CourseId = strconv.Itoa(rng.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from req

	params := mux.Vars(r)

	// loop, id, remove, add with params id
	for ind, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:ind], courses[ind+1: ]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Id is not found")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for ind, course := range courses {
		if course.CourseId == params["id"] {
			// decode the message and delete
			courses = append(courses[:ind], courses[ind+1: ]...)
			json.NewEncoder(w).Encode("Successfully deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("No id found")
	return
}


