// abi_parser/abi_parser_test.go
package abi_parser

import (
	"path/filepath"
	"testing"
)

func TestParseABIFile(t *testing.T) {
	// Ruta del archivo ABI de prueba
	abiFilePath := filepath.Join("..", "ABIs", "ERC20.json")

	// Llamada a la función ParseABIFile
	events, err := ParseABIFile(abiFilePath)
	if err != nil {
		t.Fatalf("Error al parsear el archivo ABI: %v", err)
	}

	// Verificación del número de eventos encontrados
	expectedEventCount := 2
	if len(events) != expectedEventCount {
		t.Errorf("Se esperaban %d eventos, pero se encontraron %d", expectedEventCount, len(events))
	}

	// Verificación del primer evento (Transfer)
	transferEvent := events[1]
	if transferEvent.Name != "Transfer" {
		t.Errorf("Se esperaba el evento 'Transfer', pero se encontró '%s'", transferEvent.Name)
	}

	expectedTransferID := calculateEventID("Transfer(address,address,uint256)")
	if transferEvent.ID != expectedTransferID {
		t.Errorf("Se esperaba el ID del evento '%s', pero se encontró '%s'", expectedTransferID, transferEvent.ID)
	}

	// Verificación del segundo evento (Approval)
	approvalEvent := events[0]
	if approvalEvent.Name != "Approval" {
		t.Errorf("Se esperaba el evento 'Approval', pero se encontró '%s'", approvalEvent.Name)
	}

	expectedApprovalID := calculateEventID("Approval(address,address,uint256)")
	if approvalEvent.ID != expectedApprovalID {
		t.Errorf("Se esperaba el ID del evento '%s', pero se encontró '%s'", expectedApprovalID, approvalEvent.ID)
	}
}
