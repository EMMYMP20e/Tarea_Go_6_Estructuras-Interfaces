package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"./multimedias"
)

func main() {
	cw := multimedias.ContenidoWeb{
		[]multimedias.Multimedia{},
	}
	opc := 1
	for opc != 0 {
		opc = menu()
		switch opc {
		case 1:
			agregarImagen(&cw)
		case 2:
			agregarAudio(&cw)
		case 3:
			agregarVideo(&cw)
		case 4:
			cw.Mostrar()
			pausa()
		}
	}
}

func pausa() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("[Presiona Enter para Continuar]")
	scanner.Scan()
}

func menu() int {
	var opc int
	limpiarPantalla()
	fmt.Println("1.- Capturar Imagen")
	fmt.Println("2.- Capturar Audio")
	fmt.Println("3.- Capturar Video")
	fmt.Println("4.- Mostrar")
	fmt.Println("0.- Salir")
	fmt.Print("Opcion: ")
	fmt.Scan(&opc)
	limpiarPantalla()
	return opc
}

func limpiarPantalla() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func agregarImagen(cw *multimedias.ContenidoWeb) {
	var canales int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("Capturar Imagen")
	fmt.Print("Titulo: ")
	scanner.Scan()
	titulo := scanner.Text()
	fmt.Print("Formato: ")
	scanner.Scan()
	formato := scanner.Text()
	fmt.Print("Canales: ")
	fmt.Scan(&canales)
	cw.ContenidosWeb = append(cw.ContenidosWeb, &multimedias.Imagen{titulo, formato, canales})
}

func agregarAudio(cw *multimedias.ContenidoWeb) {
	var duracion int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("Capturar Audio")
	fmt.Print("Titulo: ")
	scanner.Scan()
	titulo := scanner.Text()
	fmt.Print("Formato: ")
	scanner.Scan()
	formato := scanner.Text()
	fmt.Print("Duracion: ")
	fmt.Scan(&duracion)
	cw.ContenidosWeb = append(cw.ContenidosWeb, &multimedias.Audio{titulo, formato, duracion})
}

func agregarVideo(cw *multimedias.ContenidoWeb) {
	var frames int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("Capturar Video")
	fmt.Print("Titulo: ")
	scanner.Scan()
	titulo := scanner.Text()
	fmt.Print("Formato: ")
	scanner.Scan()
	formato := scanner.Text()
	fmt.Print("Frames: ")
	fmt.Scan(&frames)
	cw.ContenidosWeb = append(cw.ContenidosWeb, &multimedias.Video{titulo, formato, frames})
}
