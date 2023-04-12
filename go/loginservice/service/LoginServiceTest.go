package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	dto "github.com/indranureska/service/dto"
)

func RunLoginServiceTest() {
	fmt.Println("Running user login test...")

	// Login
	var loginUser dto.User
	loginUser.UserEmail = "projectzerofour@gmail.com"
	loginUser.Password = "abcdefg"

	loginUserJson, err := json.Marshal(loginUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User to login : ")
	fmt.Println(string(loginUserJson))

	client := &http.Client{Timeout: time.Duration(5) * time.Second}
	userLoginReq, err := http.NewRequest(http.MethodPost, SERVICE_URL+"/login", bytes.NewBuffer(loginUserJson))
	if err != nil {
		panic(err)
	}

	// Set request header content-type for json
	userLoginReq.Header.Set("Content-Type", "application/json; charset=utf-8")
	loginUserResp, err := client.Do(userLoginReq)
	if err != nil {
		panic(err)
	}

	defer loginUserResp.Body.Close()

	loginRespBody, err := ioutil.ReadAll(loginUserResp.Body)

	if loginUserResp.StatusCode != http.StatusBadRequest {
		if err != nil {
			panic(err)
		}

		var userloggedIn dto.User
		json.Unmarshal(loginRespBody, &userloggedIn)

		fmt.Println("user logged in: " + userloggedIn.UserEmail)

	} else {
		// Display error message from service
		var errorResponse dto.ErrorResponse
		json.Unmarshal(loginRespBody, &errorResponse)

		fmt.Println("login failed with error: ", errorResponse.Error)
		panic(loginUserResp.Status)
	}

	// Logout
	fmt.Println("Logout")
	userLogoutReq, err := http.NewRequest(http.MethodPost, SERVICE_URL+"/logout", bytes.NewBuffer(loginUserJson))
	if err != nil {
		panic(err)
	}

	// Set request header content-type for json
	userLogoutReq.Header.Set("Content-Type", "application/json; charset=utf-8")
	logoutUserResp, err := client.Do(userLogoutReq)
	if err != nil {
		panic(err)
	}

	defer logoutUserResp.Body.Close()

	logoutRespBody, err := ioutil.ReadAll(logoutUserResp.Body)

	if logoutUserResp.StatusCode != http.StatusBadRequest {
		if err != nil {
			panic(err)
		}

		var userloggedOut dto.User
		json.Unmarshal(logoutRespBody, &userloggedOut)

		fmt.Println("user logged out: " + userloggedOut.UserEmail)

	} else {
		// Display error message from service
		var errorResponse dto.ErrorResponse
		json.Unmarshal(logoutRespBody, &errorResponse)

		fmt.Println("logout failed with error: ", errorResponse.Error)
		panic(logoutUserResp.Status)
	}

}
