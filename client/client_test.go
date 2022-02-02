package client

import (
	"testing"
)

func TestClient(t *testing.T) {
	// Configuração do cliente.
	client := Client{
		Host:      "https://runrun.it/api/v1.0/",
		AppKey:    "f9c650c98eeb28e345e0a38a184d20cb",
		UserToken: "roBknmkPI0ALmwkRuC1q",
	}

	// Consulta de feriados cadastros.
	offDays, err := client.GetOffDays()
	if err != nil {
		t.Error(err)
		return
	}

	if len(offDays) < 1 {
		t.Error("no offday returned")
	}

}
