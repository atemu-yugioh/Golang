package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Id int
	Name string
	Age int
	Class []int
}

type Students []Student

func main() {
	fmt.Println("Server running on port",5000)

	http.HandleFunc("/", MyHomePage)
	http.HandleFunc("/about", MyAboutPage)
	http.HandleFunc("/api/music", MusicApi)
	http.HandleFunc("/api/student", StudentApi)
	http.HandleFunc("/api/students", ListStudentApi)

	http.ListenAndServe(":5000", nil)
}

func ListStudentApi (w http.ResponseWriter, r *http.Request) {
	var listStudent = Students{
		Student{1, "Diep", 18, []int{1,2,3}},
		Student{2, "Nam", 18, []int{1,2,3}},
		Student{3, "Trong", 18, []int{1,2,3}},
		Student{4, "Tuan", 18, []int{1,2,3}},
	}

	json.NewEncoder(w).Encode(listStudent)
}

func StudentApi (w http.ResponseWriter, r *http.Request) {
	var student = Student{Name: "nguyen van a", Id: 1, Age: 15, Class: []int{1,2,3}}

	json.NewEncoder(w).Encode(student)
}


func MusicApi(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"name": "con mua ngang qua",
		"singer": "Son Tung",
	}
	json.NewEncoder(w).Encode(data)
}

func MyHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This is home page</h1>")
}

func MyAboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>this is about page</h1>")
}
