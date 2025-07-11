package main

import "github.com/ignaciomilesi/go-drive-uploader/services/driverUp"

func main() {
	service := driverUp.Autentificacion()
	driverUp.ListFiles(service)
	driverUp.UploadFile(service, "/Users/matiasmartini/Desktop/prueba/go-drive-uploader/prueba.txt")
	driverUp.ListFiles(service)
}
