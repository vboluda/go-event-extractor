// abi_parser/abi_parser.go
package abi_parser

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

// Inicializar loggers con prefijos para diferentes niveles de logging
var (
	infoLogger  = log.New(os.Stdout, "INF: ", log.Ldate|log.Ltime)
	logLogger   = log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

// Event representa un evento de contrato con su información
type Event struct {
	Name       string
	Signature  string
	ID         string
	Parameters []EventParameter
}

// EventParameter representa un parámetro de un evento
type EventParameter struct {
	Type string
	Name string
}

// ParseABIFile analiza el archivo ABI y extrae todos los eventos
func ParseABIFile(filePath string) ([]Event, error) {
	infoLogger.Printf("Iniciando el análisis del archivo ABI: %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		errorLogger.Printf("Error al abrir el archivo: %v", err)
		return nil, fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()
	logLogger.Println("Archivo ABI abierto correctamente")

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		errorLogger.Printf("Error al leer el archivo: %v", err)
		return nil, fmt.Errorf("error al leer el archivo: %v", err)
	}
	logLogger.Println("Archivo ABI leído correctamente")

	var parsedABI abi.ABI
	if err := json.Unmarshal(byteValue, &parsedABI); err != nil {
		errorLogger.Printf("Error al parsear el ABI: %v", err)
		return nil, fmt.Errorf("error al parsear el ABI: %v", err)
	}
	logLogger.Println("Archivo ABI parseado correctamente")

	var events []Event
	for _, event := range parsedABI.Events {
		logLogger.Printf("Procesando evento: %s", event.Name)
		eventID := calculateEventID(event.Sig)
		ev := Event{
			Name:      event.Name,
			Signature: event.Sig,
			ID:        eventID,
		}

		for _, input := range event.Inputs {
			logLogger.Printf("Añadiendo parámetro: Name=%s, Type=%s", input.Name, input.Type.String())
			ev.Parameters = append(ev.Parameters, EventParameter{
				Type: input.Type.String(),
				Name: input.Name,
			})
		}

		events = append(events, ev)
		logLogger.Printf("Evento procesado: %s, ID: %s", ev.Name, ev.ID)
	}

	logLogger.Printf("Total de eventos extraídos: %d", len(events))
	return events, nil
}

// calculateEventID calcula el hash Keccak-256 de la signatura del evento
func calculateEventID(signature string) string {
	infoLogger.Printf("Calculando el ID del evento para la signatura: %s", signature)
	hash := crypto.Keccak256Hash([]byte(signature))
	eventID := hex.EncodeToString(hash.Bytes())
	logLogger.Printf("ID del evento calculado: %s", eventID)
	return eventID
}
