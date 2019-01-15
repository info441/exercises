package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type context struct {
	port    string
	message string
}

func main() {
	port := getENVOrExit("PORT")
	fileName := getENVOrExit("FILEPATH")

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("error opening %s: %s", fileName, err.Error())
	}
	defer f.Close()
	messageBytes, _ := ioutil.ReadAll(f)
	http.HandleFunc("/", context{port: port, message: string(messageBytes)}.successHandler)
	log.Println("listening on 0.0.0.0:" + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (ctx context) successHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		Success!
		
		You are able to connect to your docker container on port ` + ctx.port +
		`	The secret message is: ` + ctx.message))
}

// gets the value of the environment variable of envName and returns it
// terminates the process if there is not value for envName
func getENVOrExit(envName string) string {
	if env := os.Getenv(envName); len(env) > 0 {
		return env
	}
	log.Fatalf("no value set for %s, please set a value for %s", envName, envName)
	return ""
}
