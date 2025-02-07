package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestParseFtpwho(t *testing.T) {
	file, err := os.Open("examples/sftp.json")
	if err != nil {
		fmt.Printf("[ERROR] Cant open examples/sftp.json: %s", err)
		return
	}
	defer file.Close()

	var ftpwhoOutput ftpwho
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ftpwhoOutput)
	if err != nil {
		fmt.Printf("[ERROR] Cant parse examples/sftp.json: %s", err)
		return
	}
	prettyOutput, _ := json.MarshalIndent(ftpwhoOutput, "", "  ")
	fmt.Printf("%s", prettyOutput)
}
