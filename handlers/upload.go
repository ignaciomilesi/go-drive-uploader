package handlers

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	//"github.com/ignaciomilesi/go-drive-uploader/services"
)

func UploadToDrive(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Nueva solicitud")

	// Revisamos que sea un método post
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido, solo se admite POST", http.StatusMethodNotAllowed)
		fmt.Println("Método no permitido. Solicitud -> ", r.Method)
		return
	}

	// Parseamos el formulario
	// (el ParseMultipartForm es por que hay un archivo en el formulario, sino podría usar ParseForm)
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
		fmt.Println("Error al parsear el formulario")
		return
	}

	// Cargamos el archivo
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error al cargar el archivo", http.StatusBadRequest)
		fmt.Println("Error al cargar el archivo")
		return
	}
	defer file.Close()

	// hacemos una copia del archivo en el temporal
	tempPath, err := guardarEnTemp(file, header.Filename)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempPath)

	fmt.Println("todo bien ", tempPath)
	//services.DriverUpload()

	// Respondemos al cliente con un mensaje de texto plano
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)                      // 200 OK
	w.Write([]byte("Archivo recibido correctamente")) // Respuesta de texto plano
}

func guardarEnTemp(file multipart.File, fileName string) (string, error) {

	// Creamos un archivo temporal
	tempPath := filepath.Join(os.TempDir(), fileName)
	tempFile, err := os.Create(tempPath)

	if err != nil {
		return "", errors.New("Error al crear archivo temporal")
	}
	defer tempFile.Close()

	// Copiamos el archivo subido en el archivo temporal
	_, err = io.Copy(tempFile, file)
	if err != nil {
		return "", errors.New("Error al guardar archivo")
	}

	return tempPath, nil

}
