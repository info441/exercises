package main

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

// YOU DO NOT NEED TO MAKE CHANGES TO THIS FILE

//usage is the usage string shown when not enough
//arguments are passed to this command
const usage = `
usage:
	hmac sign secretKey < file-to-sign
	hmac verify secretKey signature < file-that-was-signed
When using 'sign' the output will be a base64-encoded
HMAC digital signature for the file using the secretKey.
When using 'verify' the output will be "signature is valid"
if the provided signature is valid for the file and secretKey,
or "signature is INVALID" if the provided signature is
invalid for the file and secretKey
`

//showErrorAndExit prints the error message
//and exits with a non-success code
func showErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

//showUsageAndExit shows the usage string
//and exits with a non-success code
func showUsageAndExit() {
	showErrorAndExit(errors.New(usage))
}

func main() {
	
	if len(os.Args) < 2 {
		showUsageAndExit()
	}
	
	command := os.Args[1]
	secretKey := []byte(os.Args[2])
	reader := bufio.NewReader(os.Stdin)
	
	switch command {
	case "sign":
		signature, err := Sign(reader, secretKey)
		if err != nil {
			showErrorAndExit(err)
		}
		
		id := base64.URLEncoding.EncodeToString(signature)
		fmt.Println(id)
	case "verify":
		signature, _ := base64.URLEncoding.DecodeString(os.Args[3])
		valid, err := Verify(reader, secretKey, signature)
		if err != nil {
			showErrorAndExit(err)
		}
		
		if valid {
			fmt.Println("Signature is Valid")
		} else {
			fmt.Println("Signature is Invalid")
		}
	default:
		showUsageAndExit()
		os.Exit(1)
	}
}
