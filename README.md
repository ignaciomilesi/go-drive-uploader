# go-drive-uploader
API en Go para subir documentos a Google Drive y registre los datos en MySQL


Generar las credenciales necesarias. 

### Paso 1: Crear un proyecto en Google Cloud Console

1. Ir a [Google Cloud Console](https://console.cloud.google.com/).

2. En la esquina superior izquierda, clic en  **selector de proyectos** y luego en **Nuevo Proyecto**. Dale un nombre, elegir la ubicación (puede ser la predeterminada) y **Crear**.

### Paso 2: Habilitar la API de Google Drive

1. En el menú de navegación (en la barra lateral izquierda), seleccionar **APIs y Servicios** > **Biblioteca**. Buscar "Google Drive API" y habilitarla.

### Paso 3: Configurar OAuth 2.0 para la autenticación

1. En el menú de navegación, ve a **APIs y Servicios** > **Credenciales**.

2. Clic en el botón **Crear credenciales** y selecciona **ID de cliente de OAuth**.
   
   * Si es la primera vez, pedirá configurar la pantalla de consentimiento:

3. En **Tipo de aplicación**, selecciona **Aplicación de escritorio** y darle un nombre al cliente.

4. Después de crear el cliente, se generará un archivo `credentials.json` para descargar.


Cuando usas OAuth 2.0, el primer acceso requerirá que un usuario inicie sesión y autorice el acceso. La API generará un token de acceso que se guarda en el archivo `token.json`. Los siguientes inicios utilizara el token generado (no solicitara el inicio de sesion) 


Cuando se crea la aplicacion, se coloca en modo de prueba y solo los ususarios marcados como testers pueden utilizarlar, para agregarlos:

   * En el menú de la izquierda, **APIs y Servicios** > **Pantalla de consentimiento OAuth**.

   * En la sección **Publico**, al final, se encontrara **Usuarios de prueba**, agrega las direcciones de correo electrónico de las personas podran usar la aplicación.

Si la aplicación está lista, en la sección **Publico** se puede publicar la aplicacion para que pueda ser utilizada por usuarios fuera los testers.