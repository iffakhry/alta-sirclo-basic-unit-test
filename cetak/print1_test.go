package cetak

import "testing"

func TestCetakNama(t *testing.T) {
	expected := "alta"
	actual := CetakNama("alta")
	if actual != expected {
		t.Error("error")
	}
}
