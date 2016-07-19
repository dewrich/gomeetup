package main

import "fmt"

type foo struct {
	Message    string `json:"message"`
	Ports      []int  `json:"ports"`
	ServerName string `json:"serverName"`
}

func main() {
	fmt.Println("vim-go")
	myVariable := "testme"
	fmt.Printf("myVariable = %+v\n", myVariable)
}

func one2() {
}

func two() {

}

func three() {

}
