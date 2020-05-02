package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Set the api that returns the public ip below.
	api := "https://api.ipify.org/?format=text"

	// Create the logging file.
	file, err := os.OpenFile("wmip.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file) // Set the logging path

	// Make an HTTP call to the api.
	resp, err := http.Get(api)
	check(err) // check for errors..

	// Read the data.
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	ipaddress := string(body)
	// Close the response.Body.
	resp.Body.Close()

	f, err := os.Create("ip.txt")
	check(err)

	l, err := f.WriteString(ipaddress)
	check(err)

	log.Println("IP ["+string(ipaddress)+"] written successfully; Bytes written:", l)

	err = f.Close()
	check(err)
}
