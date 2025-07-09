package main

import "go-drive-uploader/driverUp"

func main() {
	service := driverUp.Autentificacion()
	driverUp.ListFiles(service)
	driverUp.UploadFile(service, "/Users/matiasmartini/Desktop/prueba/go-drive-uploader/prueba.txt")
	driverUp.ListFiles(service)
}
