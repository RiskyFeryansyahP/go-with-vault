package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ResponseVault struct {
	Auth Auth `json:"auth"`
}

type Auth struct {
	ClientToken string `json:"client_token"`
}

const VAULT_ADDR = "http://18.141.194.85:8200"

func main() {
	fmt.Println("Simple example authentication vault")

	res, _ := userpassMethod()

	fmt.Println(res.Auth.ClientToken)
}

func userpassMethod() (*ResponseVault, error) {
	var resp *ResponseVault

	pass := os.Getenv("PASSWORD")

	password := map[string]string{
		"password": pass,
	}

	payload, _ := json.Marshal(password)

	username := os.Getenv("USERNAME")

	url := fmt.Sprintf("%s/v1/auth/userpass/login/%s", VAULT_ADDR, username)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(respBody, &resp)

	return resp, nil
}