package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(responseData))
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Id)
	fmt.Println(responseObject.Title)
}
