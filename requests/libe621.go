/*
/	libyiff
/
/	Go library for identifying and sorting furry porn
/   written by heyitspuggo with help from spotlight
/   licensed under GNU Public License version 3
/
/   e621 request library
*/


package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetByHash asks e621 for an image's tags based on md5 hash
func GetByHash(hash string, username string, key string) {
	// Request (GET https://e621.net/posts.json?tags=md5:)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://e621.net/posts.json?tags=md5:" + hash, nil)

	// Headers
	req.Header.Add("User-Agent", "libyiff v0.1 (library by heyitspuggo, current user's username: " + username + ")")
	parseFormErr := req.ParseForm(); if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// actually make the request
	resp, err := client.Do(req); if err != nil {
		fmt.Println("Failure: ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}



