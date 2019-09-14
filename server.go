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
