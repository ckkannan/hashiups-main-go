package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp-demoapp/hashicups-client-go"
)

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id`
	Username string `json:"username`
	Token    string `json:"token"`
}

const HostURL string = "http://192.168.1.137:19090"

var h string = HostURL
var user string = "educatio1n"
var pass string = "test123"
var host *string = &h
var puser *string = &user
var ppass *string = &pass

func main() {

	h, err := hashicups.NewClient(host, puser, ppass)
	if err != nil {
		fmt.Println("New client Failed")
		os.Exit(1)
	}
	h.SignIn()
	var coff []hashicups.Coffee
	coff, err = h.GetCoffees()
	s, _ := json.Marshal(coff)
	fmt.Println(string(s))
	for i := 0; i < len(coff); i++ {
		fmt.Println(coff[i].ID, "  ", coff[i].Name)
	}
}
