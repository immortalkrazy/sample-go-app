package main

import (
	"fmt"
	"net/http"
)

var taskOne = "This is task One"
var taskTwo = "This is task Three"
var taskThr = "This is task three"

var taskList = []string{taskOne, taskTwo, taskThr}

func main() {
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/", greetings)
	http.HandleFunc("/show-tasks", showTasks)
}

func greetings(writer http.ResponseWriter, request *http.Request) {
	var greet = "Helle Dear, Greetings...."
	fmt.Fprintln(writer, greet)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskList {
		fmt.Fprintln(writer, task)
	}
}
