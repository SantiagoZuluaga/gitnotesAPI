package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func PreventPrint(arg interface{}) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv != "production" {
		fmt.Println(arg)
	}
}
