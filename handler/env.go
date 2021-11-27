package handler

import "os"

var URL string = os.Getenv("API_URL")
var TOKEN_MANAGER_NAME string = os.Getenv("TOKEN_MANAGER")
