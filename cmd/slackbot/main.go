package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/stretchr/graceful"
)

func main() {
	// token := flag.String("t", "dummyToken", "API Token")
	// flag.Parse()
	// log.Println(*token)

	mux := http.NewServeMux()
	mux.HandleFunc("/slackbot/", handleMsg)
	graceful.Run(":8140", 1*time.Second, mux)

	// resp, err := http.Get("https://slack.com/api/auth.test?token=" + *token + "&pretty=1")
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer resp.Body.Close()
	// byteArray, _ := ioutil.ReadAll(resp.Body)
	// log.Println(string(byteArray))
}

func handleMsg(w http.ResponseWriter, r *http.Request) {
	const fname = "handleMovies"
	log.Println(fname, "START")

	switch r.Method {
	case "POST":
		res := map[string]string{"text": "African or European?"}
		log.Printf("%+v\n", r)
		ba, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%+v\n", string(ba))
		respond(w, r, http.StatusOK, res)
	default:
	}
	log.Println(fname, "END")
}

func encodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
	if data != nil {
		encodeBody(w, r, data)
	}
}
