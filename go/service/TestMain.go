package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	dto "github.com/indranureska/service/dto"
)

func main() {
	log.Println("running test...")

	log.Println("1. Get list of user")

	log.Println("2. Find user")

	log.Println("3. Create user")
	var user dto.User
	user.FirstName = "John"
	user.LastName = "Doe"
	user.UserEmail = "john.doe@mail.com"
	user.Password = "abcdefg"
	user.LastLogin = time.Now().UTC().String()

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User to create : ")
	fmt.Println(string(userJson))

	resp, err := http.Post("http://localhost:8000/create-user", "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		jsonStr := string(body)
		fmt.Println("Response : ", jsonStr)
	} else {
		fmt.Println("Get failed with error: ", resp.Status)
	}
}
