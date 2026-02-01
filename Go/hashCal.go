package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
)

type Options struct {
	InputFile string
}

func callOptions() *Options {
	fmt.Print("Este programa permite calcular los siguientes hash para un archivo: MD5, SHA1, SHA256\n\n")
	inputFile := flag.String("F", "", "Ingresar archivo ")
	flag.Parse()

	return &Options{
		InputFile: *inputFile,
	}
}

func main() {
	menu := callOptions()

	if menu.InputFile == "" {
		log.Fatal("[!] Debe ingresar un archivo.\nEjemplo: hashcal -F <File>")
	}

	entradaData, err := os.ReadFile(menu.InputFile)
	if err != nil {
		log.Fatal("[!] Error al leer el archivo.")
	}

	md5sum := md5.Sum(entradaData)
	sha1sum := sha1.Sum(entradaData)
	sha256sum := sha256.Sum256(entradaData)

	fmt.Printf("MD5: %x\n", md5sum)
	fmt.Printf("SHA1: %x\n", sha1sum)
	fmt.Printf("SHA256: %x\n", sha256sum)
}
