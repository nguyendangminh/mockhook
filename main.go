package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

var port int
var verifyToken string
var hook string

func main() {
	flag.IntVar(&port, "p", 1203, "Port number for binding")
	flag.StringVar(&verifyToken, "t", "", "App verify token")
	flag.StringVar(&hook, "h", "webhook", "Hook path")
	flag.Parse()
	if verifyToken == "" {
		logrus.Error("App verify token is required")
		Usage()
		os.Exit(1)
	}

	http.HandleFunc("/", handle)
	logrus.Infoln(fmt.Sprintf("The webhook is ready at port %d/%s", port, hook))
	logrus.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		verify(w, r)
	case "POST":
		printMessage(r)
	default:
		logrus.Errorf("Method %s is not supported", r.Method)
	}
}

func verify(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("hub.mode") == "subscribe" && r.FormValue("hub.verify_token") == verifyToken {
		fmt.Fprintf(w, r.FormValue("hub.challenge"))
		logrus.Info("Verified")
		return
	}
	logrus.WithFields(logrus.Fields{"verifyToken": verifyToken, "hub.verify_token": r.FormValue("hub.verify_token")}).Error("Tokens does NOT match.")
	http.Error(w, "Validation failed, make sure the tokens match.", http.StatusForbidden)
	return
}

func printMessage(r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error("ReadAll the request body failed. Error ", err)
		return
	}
	defer r.Body.Close()

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		logrus.WithFields(logrus.Fields{"body": string(body)}).Error("Unmarshal request body failed: ", err)
		return
	}

	pj, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		logrus.Error("MarshalInden failed: ", err)
		return
	}
	logrus.Infof("New message: \n%s\n", string(pj))
}
