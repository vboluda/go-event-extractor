// main.go
package main

import (
	"fmt"
	"path/filepath"

	"github.com/vboluda/go-event-extractor/abi_parser"
)

func main() {
	// Ruta del archivo ABI en el directorio raíz del proyecto
	abiFilePath := filepath.Join("ABIs", "ERC20.json")

	// Llamada a la función ParseABIFile del paquete abi_parser
	events, err := abi_parser.ParseABIFile(abiFilePath)
	if err != nil {
		fmt.Printf("Error al parsear el archivo ABI: %v\n", err)
		return
	}

	// Imprimir los eventos extraídos
	for _, event := range events {
		fmt.Printf("Evento: %s\n", event.Name)
		fmt.Printf("  Signature: %s\n", event.Signature)
		fmt.Printf("  ID: %s\n", event.ID)
		fmt.Printf("  Parameters:\n")
		for _, param := range event.Parameters {
			fmt.Printf("    - %s: %s\n", param.Type, param.Name)
		}
		fmt.Println()
	}
}
