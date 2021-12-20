package calculator

import (
	"testing"
)

func TestAddition(t *testing.T) {
	t.Run("case bilangan positif", func(t *testing.T) {
		if Addition(1, 2) != 3 {
			t.Error("Error, expected 3")
		}
	})
	t.Run("case bilangan negatif", func(t *testing.T) {
		if Addition(-1, -2) != -3 {
			t.Error("Error, expected 3")
		}
	})
}

func TestAdditionNegative(t *testing.T) {
	expected := -3
	actual := Addition(-1, -2)
	if actual != expected {
		t.Error("Error, expected 3")
	}

}

func TestAdditionNegativePositive(t *testing.T) {
	actual := Addition(-1, 2)
	expected := 1
	if actual != expected {
		t.Error("Error, expected 1 actual", actual)
	}
}

func TestAdditionPositiveNegative(t *testing.T) {
	actual := Addition(1, -2)
	expected := -1
	if actual != expected {
		t.Error("Error, expected 1 actual", actual)
	}
}

func TestPengurangan(t *testing.T) {
	actual := Pengurangan(5, 3)
	expected := 2
	if actual != expected {
		t.Error("Error,expedted 2, actual ", actual)
	}

}

func TestPenguranganLebih10(t *testing.T) {
	actual := Pengurangan(11, 3)
	expected := 11
	if actual != expected {
		t.Error("Error,expedted 2, actual ", actual)
	}

}
