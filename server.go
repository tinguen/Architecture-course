package main;

import (
	"fmt"
	"net/http"
)

type Time struct
{
	Time string
}




func handleMain(responsewriter http.ResponseWriter, request *http.Request) {
	responsewriter.Header().Set("Content-Type", "text/html")
		responsewriter.WriteHeader(200)
		responsewriter.Write([]byte("The main page\n"))
		fmt.Println("/ page accessed");
}


