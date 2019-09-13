package api

import (
	"fmt"
	"net/http"
)

func GetUser() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.v2.vapor.cloud/v2/auth/users/me", nil)
	req.Header.Set("Authorization", "Foo")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Authorization failed")
	} else {
		if res.StatusCode >= 200 && res.StatusCode <= 299 {
			fmt.Println("Success!")
		} else {
			fmt.Println("Authentication failed")
		}
		// data, _ := ioutil.ReadAll(res.Body)
		// fmt.Println(string(data))
	}
}
