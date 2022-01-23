package main

import (
	"encoding/base64"
	"os"
)

type UserData struct {
	Script string
}

func loadEncodedUserData(userDataPath string) (string, error) {
	data, err := os.ReadFile(userDataPath)
	if err != nil {
		panic(err)
	}

	encodedData := base64.StdEncoding.EncodeToString(data)

	return encodedData, err
}
