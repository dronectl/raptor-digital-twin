package tripzones

import (
	"testing"
)

func TestTripzoneMinMax(t *testing.T) {
	tz := New(100.0, 0.0)
	m := 50.0
	tz.Min(m)
	if tz.min != m {
		t.Errorf("Expected %v got %v", m, tz.min)
	}
	tz.Max(m)
	if tz.max != m {
		t.Errorf("Expected %v got %v", m, tz.max)
	}
}

func TestTripzoneCheckTripCondition(t *testing.T) {
	tz := New(100.0, 0.0)
	if tz.CheckTripCondition(50.0) {
		t.Errorf("Expected tripped to be false, got %t", tz.tripped)
	}
	if tz.CheckTripCondition(150.0) == false {
		t.Errorf("Expected tripped to be true, got %t", tz.tripped)
	}
}

func TestTripzoneNew(t *testing.T) {
	tz := New(0.0, 0.0)
	if tz.min != 0.0 {
		t.Errorf("Expected min to be 0.0, got %f", tz.min)
	}
	if tz.max != 0.0 {
		t.Errorf("Expected max to be 0.0, got %f", tz.max)
	}
	if tz.tripped != false {
		t.Errorf("Expected tripped to be false, got %t", tz.tripped)
	}
}
