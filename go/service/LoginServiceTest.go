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
	// findUserResp, err := http.Get(SERVICE_URL + "/find-user/projectzerofour@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Status code :", findUserResp.StatusCode)

	// defer findUserResp.Body.Close()
	// findUserRespBody, err := ioutil.ReadAll(findUserResp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("User found: ", string(findUserRespBody))

	var loginUser dto.User
	loginUser.UserEmail = "projectzerofour@gmail.com"
	loginUser.Password = "abcedfg"

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

	body, err := ioutil.ReadAll(loginUserResp.Body)

	if loginUserResp.StatusCode != http.StatusBadRequest {
		if err != nil {
			panic(err)
		}

		var userloggedIn dto.User
		json.Unmarshal(body, &userloggedIn)

		fmt.Println("user logged in: " + userloggedIn.UserEmail)

	} else {
		// Display error message from service
		var errorResponse dto.ErrorResponse
		json.Unmarshal(body, &errorResponse)

		fmt.Println("login failed with error: ", errorResponse.Error)
		panic(loginUserResp.Status)
	}
}
