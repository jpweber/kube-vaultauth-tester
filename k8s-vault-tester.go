package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func sendRequest(jwt, roleName, vaultUrl string) {
	reqData := map[string]string{
		"jwt":  jwt,
		"role": roleName,
	}

	jsonReq, _ := json.Marshal(reqData)
	body := bytes.NewBuffer(jsonReq)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", vaultUrl, body)

	// Headers
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status: ", resp.Status)
	fmt.Println("response Headers: ", resp.Header)
	fmt.Println("response Body: ", string(respBody))
}

func readToken() string {
	tokenFile := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	file, err := os.Open(tokenFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	log.Println("Service token: ", string(b))
	return string(b)
}

func main() {
	vaultURL := os.Getenv("VAULT_URL")
	role := os.Getenv("ROLE")
	token := readToken()
	sendRequest(token, role, vaultURL)
}
