package main;

import(
	"fmt"
	"log"
	"time"
	"encoding/json"
	"net/http")

type Time struct
{
	Time string
}


func handleTime(responsewriter http.ResponseWriter, request *http.Request) {

	datatosend := Time{time.Now().Format(time.RFC3339)}

  responsewriter.Header().Set("Content-Type", "application/json")
	responsewriter.WriteHeader(200)
  json.NewEncoder(responsewriter).Encode(datatosend)

	}

	func handleMain(responsewriter http.ResponseWriter, request *http.Request) {

    responsewriter.Header().Set("Content-Type", "text/html")
		responsewriter.WriteHeader(200)
		responsewriter.Write([]byte("The main page\n"))
		fmt.Println("/ page accessed");

	}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/time", handleTime)

	fmt.Println("Listening on 8795...");

	log.Fatal(http.ListenAndServe(":8795", nil))
}
