package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ignaciomilesi/go-drive-uploader/handlers"
)

func main() {

	http.HandleFunc("/uploadDrive", handlers.UploadToDrive)

	fmt.Println("Server montado")

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
