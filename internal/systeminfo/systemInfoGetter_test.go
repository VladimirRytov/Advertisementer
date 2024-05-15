package systeminfo

import "testing"

func TestCPU(t *testing.T) {
	_, err := CPU()
	if err != nil {
		t.Fatalf("gor error %v", err)
	}
}

func TestComputerName(t *testing.T) {
	_, err := ComputerName()
	if err != nil {
		t.Fatalf("gor error %v", err)
	}
}
