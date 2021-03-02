package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
}

func main() {
	response, _ := http.Get("http://192.168.1.43:8081/jenkins/")
	user := User{xml.Name{"", "user"}}
	temp, _ := ioutil.ReadAll(response.Body)
	body := []byte(temp)
	xml.Unmarshal(body, &user)
	fmt.Printf("status: %s", user)
}
