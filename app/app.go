package app

import (
	"net/http"
	"os"

	"github.com/SantiagoZuluaga/gitnotesAPI/app/common"
)

func RunServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	//router := representation.Router()
	common.PreventPrint("Server running in 	port: http://localhost" + port)
	http.ListenAndServe(port, nil)
}
