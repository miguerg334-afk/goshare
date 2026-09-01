package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/mdp/qrterminal/v3"
)

// Obtiene la primera IP local privada (Wi-Fi o Ethernet) descartando loopback (127.0.0.1)
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no se encontró una interfaz de red activa")
}

func main() {
	// 1. Validar que el usuario pase la ruta del archivo como argumento
	if len(os.Args) < 2 {
		fmt.Println("Uso: goshare <ruta-del-archivo>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// 2. Verificar que el archivo existe
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Error: El archivo '%s' no existe.\n", filePath)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		fmt.Println("Error: Por ahora solo se pueden compartir archivos, no carpetas.")
		os.Exit(1)
	}

	// 3. Obtener la IP local de tu máquina en la red
	ip, err := getLocalIP()
	if err != nil {
		fmt.Println("Error obteniendo la IP local:", err)
		os.Exit(1)
	}

	// 4. Buscar un puerto disponible automáticamente asignando :0
	listener, err := net.Listen("tcp", ip+":0")
	if err != nil {
		fmt.Println("Error abriendo el puerto de red:", err)
		os.Exit(1)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	fileName := filepath.Base(filePath)
	downloadURL := fmt.Sprintf("http://%s:%d/%s", ip, port, fileName)

	// 5. Configurar la ruta HTTP para servir el archivo
	http.HandleFunc("/"+fileName, func(w http.ResponseWriter, r *http.Request) {
		// Forzar al navegador a descargar el archivo
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
		http.ServeFile(w, r, filePath)
		fmt.Printf("\n[+] ¡Descarga iniciada por %s!\n", r.RemoteAddr)
	})

	// 6. Imprimir pantalla con el código QR y la URL
	fmt.Print("\033[H\033[2J") // Limpiar consola
	fmt.Println("==================================================")
	fmt.Printf(" Compartiendo: %s (%.2f MB)\n", fileName, float64(fileInfo.Size())/1024/1024)
	fmt.Println(" Escanea el código QR con tu celular:")
	fmt.Println("==================================================")

	// Generar el código QR directamente en la consola
	qrterminal.GenerateHalfBlock(downloadURL, qrterminal.L, os.Stdout)

	fmt.Println("\n O accede directamente a la URL:")
	fmt.Printf(" 👉 %s\n", downloadURL)
	fmt.Println("==================================================")
	fmt.Println("Presiona Ctrl+C para detener el servidor.")

	// 7. Manejo de señales para un cierre limpio con Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := http.Serve(listener, nil); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error en el servidor:", err)
		}
	}()

	<-stop
	fmt.Println("\n[-] Servidor detenido.")
}
