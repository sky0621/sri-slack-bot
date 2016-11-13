package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/stretchr/graceful"
)

func main() {
	SetupLog(".")
	log.Println("App Start")

	mux := http.NewServeMux()
	mux.HandleFunc("/slackbot/", handleMsg)
	log.Println("Server Start")
	graceful.Run(":8140", 1*time.Second, mux)
}

func handleMsg(w http.ResponseWriter, r *http.Request) {
	const fname = "handleMsg"
	log.Println(fname, "START")

	switch r.Method {
	case "POST":
		r.ParseForm()
		log.Println(r.Form)

		sa := r.Form["text"]
		res := map[string]string{"text": strings.Join(sa, "|") + " ？"}
		unames := r.Form["user_name"]
		for _, uname := range unames {
			if strings.Contains(uname, "slackbot") {
				res = map[string]string{"text": ""}
				respond(w, r, http.StatusOK, res)
				return
			}
		}
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

// SetupLog ...
func SetupLog(outputDir string) (*os.File, error) {
	logfile, err := os.OpenFile(filepath.Join(outputDir, "slackbot.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("[%s]のログファイル「slackbot.log」オープンに失敗しました。 [ERROR]%s\n", outputDir, err)
		return nil, err
	}

	// [MEMO]内容に応じて出力するファイルを切り替える場合はどうするんだ・・・？
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	return logfile, nil
}
