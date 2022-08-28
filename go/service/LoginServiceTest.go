package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RunLoginServiceTest() {
	fmt.Println("Running user login test...")
	findUserResp, err := http.Get(SERVICE_URL + "/find-user/projectzerofour@gmail.com")
	if err != nil {
		panic(err)
	}

	fmt.Println("Status code :", findUserResp.StatusCode)

	defer findUserResp.Body.Close()
	findUserRespBody, err := ioutil.ReadAll(findUserResp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("User found: ", string(findUserRespBody))
}
