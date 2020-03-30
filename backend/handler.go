package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(w, "Hola desde el handler")
}

func HandleHome(w http.ResponseWriter, res *http.Request) {

	fmt.Fprintf(w, "<h1>hola desde el handler home </h1>")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}
	fmt.Fprintf(w, "payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CheckAut(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user Login
	err := decoder.Decode(&user)
	enableCors(&w)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}
	//response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "true ")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
