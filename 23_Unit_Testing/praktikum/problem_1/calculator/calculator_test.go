package calculator

import "testing"

func TestAddition(t *testing.T) {
	if Adidition(1, 2) != 3 {
		t.Error("Harusnya 1 (+) 2 sama dengan 3")
	}
	if Adidition(-1, -2) != -3 {
		t.Error("Harusnya -1 (+) -2 sama dengan -3")
	}
}
func TestSubtration(t *testing.T) {
	if Subtration(1, 2) != -1 {
		t.Error("Harusnya 1 (-) 2 sama dengan -1")
	}
	if Subtration(-1, -2) != 1 {
		t.Error("Harusnya -1 (-) -2 sama dengan 1")
	}
}

func TestDivision(t *testing.T) {
	if Division(1, 2) != 1/2 {
		t.Error("Harusnya 1 (/) 2 sama dengan 0.5")
	}
	if Division(-1, -2) != -1/-2 {
		t.Error("Harusnya -1 (/) -2 sama dengan 0.5")
	}
}
func TestMultiplication(t *testing.T) {
	if Multiplication(1, 2) != 2 {
		t.Error("Harusnya 1 (*) 2 sama dengan 2")
	}
	if Multiplication(-1, -2) != 2 {
		t.Error("Harusnya -1 (*) -2 sama dengan 2")
	}
}
