package multimedias

import "fmt"

type ContenidoWeb struct {
	ContenidosWeb []Multimedia
}

func (cw *ContenidoWeb) Mostrar() {
	for _, v := range cw.ContenidosWeb {
		v.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales int64
}

func (i *Imagen) Mostrar() {
	fmt.Println("Imagen:")
	fmt.Println("Titulo: ", i.Titulo, ", Formato: ", i.Formato, ", Canales: ", i.Canales)
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int64
}

func (a *Audio) Mostrar() {
	fmt.Println("Audio:")
	fmt.Println("Titulo: ", a.Titulo, ", Formato: ", a.Formato, ", Duracion: ", a.Duracion)
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int64
}

func (v *Video) Mostrar() {
	fmt.Println("Video:")
	fmt.Println("Titulo: ", v.Titulo, ", Formato: ", v.Formato, ", Frames: ", v.Frames)
}
