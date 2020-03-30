package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

type Login struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func (u Login) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

type MetaData interface{}
