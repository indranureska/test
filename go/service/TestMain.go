package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	dto "github.com/indranureska/service/dto"
)

const SERVICE_URL = "http://localhost:8000"

func main() {
	fmt.Println("Running test...")

	fmt.Println("1. Create user")
	fmt.Println("start")
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

	resp, err := http.Post(SERVICE_URL+"/create-user", "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var userCreateResp dto.Response
		json.Unmarshal(body, &userCreateResp)

		fmt.Println("Inserted doc id: ", userCreateResp.InsertedID)

	} else {
		fmt.Println("User creation failed with error: ", resp.Status)
		panic(resp.Status)
	}
	fmt.Println("end")

	fmt.Println("2. Get list of user")
	fmt.Println("start")
	getAllResp, err := http.Get(SERVICE_URL + "/user-list")
	if err != nil {
		panic(err)
	}

	defer getAllResp.Body.Close()
	getAllRespBody, err := ioutil.ReadAll(getAllResp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response: ", string(getAllRespBody))

	var users []dto.User
	json.Unmarshal(getAllRespBody, &users)

	for i, s := range users {
		fmt.Println(i, "id: "+s.UserEmail)
	}
	fmt.Println("Unmarshal: ", users)

	fmt.Println("end")

	fmt.Println("3. Find user")
	fmt.Println("start")

	fmt.Println("end")
}
