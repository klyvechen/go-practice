package model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetTopo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	result := []byte{}
	jsonresult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintln(w, string(jsonresult))
}
