package menu

import (
	"strings"
	"testing"
)

func TestFormatCheatSuccessMessage(t *testing.T) {
	message := formatCheatSuccessMessage(250)

	if !strings.Contains(message, "Vous avez bien triché") {
		t.Fatalf("message attendu avec texte de triche, obtenu : %q", message)
	}

	if !strings.Contains(message, "+250") {
		t.Fatalf("message attendu avec montant, obtenu : %q", message)
	}

	if !strings.Contains(message, "argent") {
		t.Fatalf("message attendu avec mention de l'argent, obtenu : %q", message)
	}
}
