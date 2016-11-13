package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/stretchr/graceful"
)

func main() {
	// token := flag.String("t", "dummyToken", "API Token")
	// flag.Parse()
	// log.Println(*token)

	SetupLog(".")
	log.Println("App Start")

	mux := http.NewServeMux()
	mux.HandleFunc("/slackbot/", handleMsg)
	log.Println("Server Start")
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
