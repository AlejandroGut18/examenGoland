package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readerStrings(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Hubo un error, intente nuevamente", err)
		return ""
	}
	input = strings.TrimSpace(input)
	return input
}

func readerInt(message string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s", message)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Hubo un error, intente nuevamente\n")
			continue
		}
		input = strings.TrimSpace(input)
		inputConv, errConv := strconv.Atoi(input)
		if errConv != nil {
			fmt.Print("Por favor ingrese un valor valido\n")
			continue
		}
		return inputConv
	}
}

type Song struct {
	Titulo  string
	Artista string
	Next    *Song
}
type Linkendlist struct {
	Head *Song
	Size uint32
}

func main() {
	lista := Linkendlist{}
	for {
		Clear()
		fmt.Print("==== LISTA DE REORDUCCIÓN ====\n1 - AGEGAR CANCIÓN\n2 - ELIMINAR CANCIÓN\n3 - BUSCAR CANCIÓN \n4 - IMPRIMIR LISTA\n5 - Salir\n")
		option := readerInt("Ingrese opción: ")
		Clear()
		if option == 5 {
			break
		}
		switch option {
		case 1:
			titulo := readerStrings("Ingrese el titulo de la canción: ")
			artista := readerStrings("Ingrese el artista de la canción: ")
			lista.insertLast(titulo, artista)

			readerStrings("Presione enter para continuar...")
		case 2:
			titulo := readerStrings("Ingrese el titulo de la canción: ")
			lista.delete(titulo)
			readerStrings("Presione enter para continuar...")
		case 3:
			titulo := readerStrings("Ingrese el titulo de la canción: ")
			lista.buscar(titulo)
			readerStrings("Presione enter para continuar...")
		case 4:
			lista.imprimir()
			readerStrings("Presione enter para continuar...")
		}
	}
}

func (lista *Linkendlist) imprimir() {
	current := lista.Head
	for current != nil {
		//fmt.Printf("%s --> ", current.Data.Titulo)
		fmt.Printf("Titulo: %s, Artista: %s --> ", current.Titulo, current.Artista)
		current = current.Next
	}
	fmt.Print("nil\n")
}
func (lista *Linkendlist) delete(titulo string) {
	if lista.Head == nil {
		fmt.Print("Lista vacía\n")
		return
	}
	if lista.Head.Titulo == titulo {
		lista.Head = lista.Head.Next
		lista.Size--
		fmt.Printf("La Canción (%s) fue eliminada con exito\n", titulo)
		return
	}
	prev := lista.Head
	current := lista.Head.Next
	for current != nil && current.Titulo != titulo {
		prev = current
		current = current.Next
	}
	if current == nil {
		fmt.Print("No se encontró\n")
		return
	}
	prev.Next = current.Next
	lista.Size--
	fmt.Printf("La Canción (%s) fue eliminada con exito\n", titulo)
}

func (lista *Linkendlist) buscar(titulo string) {
	current := lista.Head
	for current != nil {
		if current.Titulo == titulo {
			fmt.Printf("Atista: %s \nTitulo: %s\n", current.Artista, current.Titulo)
			return
		}
		current = current.Next
	}
	fmt.Printf("%s no encontrada\n", titulo)
}

func (lista *Linkendlist) insertLast(titulo, artista string) {
	newSong := Song{Titulo: titulo, Artista: artista}
	if lista.Head == nil {
		lista.Head = &newSong
	} else {
		current := lista.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &newSong
	}
	lista.Size++
}
