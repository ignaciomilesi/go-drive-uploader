package driverUp

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

func uploadFile(service *drive.Service, filePath string) (*drive.File, error) {
	// Abre el archivo local
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Crear los metadatos del archivo a subir
	fileMetadata := &drive.File{
		Name: "MiArchivo.txt", // Cambia el nombre del archivo según tu necesidad
	}

	// Crear el archivo en Google Drive
	createdFile, err := service.Files.Create(fileMetadata).
		Media(file).
		Do()

	if err != nil {
		if apiErr, ok := err.(*googleapi.Error); ok {
			if apiErr.Code == 403 {
				return nil, fmt.Errorf("permission denied for upload")
			}
		}
		return nil, err
	}

	return createdFile, nil
}

func listFiles(service *drive.Service) {
	// Solicitar los primeros 10 archivos
	files, err := service.Files.List().PageSize(10).Fields("files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	if len(files.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		fmt.Println("Files:")
		for _, file := range files.Files {
			fmt.Printf("%s (%s)\n", file.Name, file.Id)
		}
	}
}
