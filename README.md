# GoShare 🚀

> **Transfiere archivos entre dispositivos de tu red local de forma rápida, simple y privada.**

**GoShare** es una herramienta **CLI minimalista escrita en Go** que permite compartir archivos entre dispositivos conectados a la misma red **Wi-Fi / LAN**, sin utilizar servicios en la nube, cables ni aplicaciones de terceros.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge\&logo=go\&logoColor=white)
![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge\&logo=ubuntu\&logoColor=white)
![Git](https://img.shields.io/badge/Git-F05032?style=for-the-badge\&logo=git\&logoColor=white)

---

## ✨ Características

* 📱 **Código QR en la terminal**
  Genera un código QR directamente en la consola para acceder al archivo desde un smartphone.

* ⚡ **Transferencia rápida**
  Los archivos se transmiten directamente a través de tu red local, aprovechando la velocidad de tu conexión Wi-Fi o LAN.

* 🔒 **Privacidad**
  Los archivos no necesitan pasar por servidores externos ni servicios en la nube.

* 🌐 **Detección automática de red**
  Detecta automáticamente la dirección IP de la interfaz de red activa.

* 🔌 **Puerto automático**
  Utiliza un puerto disponible del sistema para iniciar el servidor de transferencia.

* 💻 **Interfaz CLI**
  Diseñado para ser rápido, sencillo y fácil de utilizar desde la terminal.

---

## 🛠️ Instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/TU_USUARIO/goshare.git
cd goshare
```

### 2. Instalar dependencias y compilar

```bash
go mod tidy
go build -o goshare
```

### 3. Instalar globalmente

Para poder ejecutar `goshare` desde cualquier directorio:

```bash
sudo mv goshare /usr/local/bin/
```

Comprueba que la instalación funciona:

```bash
goshare
```

---

## 🚀 Uso

Para compartir un archivo, ejecuta:

```bash
goshare ~/Descargas/mi_documento.pdf
```

También puedes utilizar rutas relativas:

```bash
goshare "mi foto.png"
```

Si el nombre del archivo contiene espacios, utiliza comillas:

```bash
goshare "Mis documentos/documento importante.pdf"
```

GoShare iniciará un servidor local y mostrará un **código QR en la terminal**.

📱 Escanea el código QR desde un dispositivo conectado a la **misma red Wi-Fi** y podrás acceder al archivo para descargarlo.

---

## 🔄 ¿Cómo funciona?

```text
┌──────────────┐
│   Tu PC      │
│              │
│    GoShare   │
└──────┬───────┘
       │
       │  HTTP
       │
   ┌───▼────┐
   │ Wi-Fi  │
   │  LAN   │
   └───┬────┘
       │
       │
┌──────▼───────┐
│ Smartphone   │
│              │
│ 📷 QR Code   │
│      ↓       │
│ 📥 Descargar │
└──────────────┘
```

El archivo se comparte **directamente dentro de la red local**, evitando la necesidad de subirlo a Internet.

---

## 📋 Requisitos

* **Go** instalado en el sistema.
* Dispositivos conectados a la **misma red Wi-Fi / LAN**.
* Un sistema compatible con Go.

El proyecto está pensado principalmente para **Linux / Ubuntu**, aunque Go permite adaptarlo a otros sistemas operativos.

---

## 🧰 Tecnologías

* **Go**
* **HTTP**
* **LAN / Wi-Fi**
* **QR Code**
* **CLI**

---

## 🔐 Privacidad

GoShare está diseñado para realizar las transferencias dentro de tu red local.

> **Tu archivo no necesita ser almacenado en la nube ni enviado a un servidor externo.**

Aun así, recuerda que cualquier dispositivo que pueda acceder a la dirección local mostrada por GoShare podría potencialmente acceder al archivo mientras el servidor esté activo.

---

## 📌 Estado del proyecto

🚧 **En desarrollo**

GoShare es un proyecto personal enfocado en aprender y aplicar conceptos de:

* Programación en Go
* Servidores HTTP
* Redes locales
* Manejo de archivos
* Interfaces CLI
* Comunicación entre dispositivos

---

## 🤝 Contribuciones

Las contribuciones, ideas y mejoras son bienvenidas.

Si encuentras un problema o tienes una sugerencia, puedes abrir un **Issue** o crear un **Pull Request**.

---

## 📄 Licencia

Este proyecto está disponible bajo la licencia **MIT**.
