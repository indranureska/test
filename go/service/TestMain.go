package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	user.UserEmail = "projectzerofour@gmail.com"
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

	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusCreated {
		if err != nil {
			panic(err)
		}

		var userCreateResp dto.Response
		json.Unmarshal(body, &userCreateResp)

		fmt.Println("Inserted doc id: ", userCreateResp.InsertedID)

	} else {
		// TODO: To display error message from service
		fmt.Println("User creation failed with error: ", string(body))
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
		fmt.Println(i, "User email: "+s.UserEmail)
	}
	fmt.Println("end")

	fmt.Println("3. Find user by email address")
	fmt.Println("start")
	findUserResp, err := http.Get(SERVICE_URL + "/find-user/" + user.UserEmail)
	if err != nil {
		panic(err)
	}

	defer findUserResp.Body.Close()
	findUserRespBody, err := ioutil.ReadAll(findUserResp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("User found: ", string(findUserRespBody))

	fmt.Println("end")

	fmt.Println("4. Update user")
	fmt.Println("start")

	var userToUpdate dto.User
	json.Unmarshal(findUserRespBody, &userToUpdate)

	fmt.Println("User ID to update : ", userToUpdate.ID)

	userToUpdate.FirstName = "Jhonny"

	// Marshall to JSON
	userToUpdateJson, err := json.Marshal(userToUpdate)
	if err != nil {
		panic(err)
	}

	fmt.Println("User Data to update : ", string(userToUpdateJson))

	client := &http.Client{Timeout: time.Duration(1) * time.Second}
	updateUserReq, err := http.NewRequest(http.MethodPut, SERVICE_URL+"/update-user", bytes.NewBuffer(userToUpdateJson))
	if err != nil {
		panic(err)
	}

	// Set request header content-type for json
	updateUserReq.Header.Set("Content-Type", "application/json; charset=utf-8")
	updateUserResp, err := client.Do(updateUserReq)
	if err != nil {
		panic(err)
	}

	defer updateUserResp.Body.Close()

	fmt.Println("Update response status code: " + strconv.Itoa(updateUserResp.StatusCode))

	fmt.Println("end")

	fmt.Println("5. Delete user")
	fmt.Println("start")

	fmt.Println("User ID to delete : ", userToUpdate.ID.Hex())

	deleteUserReq, err := http.NewRequest(http.MethodDelete, SERVICE_URL+"/delete-user/"+userToUpdate.ID.Hex(), nil)
	if err != nil {
		panic(err)
	}

	// Set request header content-type for json
	deleteUserReq.Header.Set("Content-Type", "application/json; charset=utf-8")
	deleteUserResp, err := client.Do(deleteUserReq)
	if err != nil {
		panic(err)
	}

	fmt.Println("Delete response status code: " + strconv.Itoa(deleteUserResp.StatusCode))
	fmt.Println("end")
}
