package driverUp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const credentialsFile = "credentials.json"
const tokenFile = "token.json"

// Scopes (permisos) de la API
var scopes = []string{
	drive.DriveFileScope, //acceso a los archivos creados o modificados por la aplicación
}

func Autentificacion() *drive.Service {
	ctx := context.Background() //manejar la cancelación, los tiempos de espera (timeouts), y los valores asociados a las solicitudes

	// Carga las credenciales del archivo
	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("No se pudo leer credentials.json: %v", err)
	}

	// Crea el cliente de autenticación
	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		log.Fatalf("No se pudo parsear credentials: %v", err)
	}

	// Verifica si tenemos un token guardado
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		// Si no tenemos token, hacer la autenticación con OAuth
		token = getTokenFromWeb(config)
		saveToken(tokenFile, token)
	}

	// Crea el cliente con el token de acceso
	client := config.Client(ctx, token)

	// Crea el servicio de Google Drive
	service, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
	}

	fmt.Println("Servicio generado")
	// devolvemos servicio para interactuar con Google Drive
	return service
}

func getTokenFromFile(file string) (*oauth2.Token, error) {
	// Leer el token desde un archivo
	tok, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal el token
	token := new(oauth2.Token)
	err = json.Unmarshal(tok, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func saveToken(file string, token *oauth2.Token) {
	// Guardar el token en un archivo para uso futuro
	tokData, err := json.Marshal(token)
	if err != nil {
		log.Fatalf("Unable to marshal token: %v", err)
	}
	err = os.WriteFile(file, tokData, 0644)
	if err != nil {
		log.Fatalf("Unable to save token: %v", err)
	}
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	// Mostrar el enlace de autenticación y obtener el código de verificación
	authURL := config.AuthCodeURL("", oauth2.AccessTypeOffline)
	fmt.Printf("Abrir el siguiente enlace en el navegador:\n%v\n", authURL)

	// Leer el código de verificación
	var authCode string
	fmt.Print("Ingresa el código de autenticación: ")
	fmt.Scan(&authCode)

	// Obtener el token
	token, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token: %v", err)
	}
	return token
}
