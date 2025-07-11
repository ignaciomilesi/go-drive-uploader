# go-drive-uploader
API en Go para subir documentos a Google Drive y registre los datos en MySQL

### estructura

```Python
project/
├── main.go           # Entry point, router y servidor
├── handlers/         # Lógica de los endpoints HTTP
│   └── upload.go
├── services/         # Lógica de negocio (Google Drive)
│   └── drive.go
├── storage/          # Base de datos
│   └── db.go                
├── utils/            # Funciones auxiliares
├── credentials.json  # Google service account credentials
├── go.mod 
├──  go.sum
```