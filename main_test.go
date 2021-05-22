package main

import "testing"

func TestDecodeFixedLengthInteger(t *testing.T) {
	t.Run("len=1", func(t *testing.T) {
		d := []byte{1}
		val, err := decodeFixedLengthInteger(d)
		if val != 1 || err != nil {
			t.Fatalf("Fail: %v -> (%d, %v)", d, val, err)
		}
	})
	t.Run("len=2", func(t *testing.T) {
		d := []byte{1, 2}
		val, err := decodeFixedLengthInteger(d)
		if val != 513 || err != nil {
			t.Fatalf("Fail: %v -> (%d, %v)", d, val, err)
		}
	})
	t.Run("len=3", func(t *testing.T) {
		d := []byte{1, 2, 3}
		val, err := decodeFixedLengthInteger(d)
		if val != 197121 || err != nil {
			t.Fatalf("Fail: %v -> (%d, %v)", d, val, err)
		}
	})
	t.Run("len=4", func(t *testing.T) {
		d := []byte{1, 2, 3, 4}
		val, err := decodeFixedLengthInteger(d)
		if val != 67305985 || err != nil {
			t.Fatalf("Fail: %v -> (%d, %v)", d, val, err)
		}
	})
	t.Run("len=5", func(t *testing.T) {
		d := []byte{1, 2, 3, 4, 5}
		val, err := decodeFixedLengthInteger(d)
		if val != 0 || err == nil {
			t.Fatalf("Fail: %v -> (%d, %v)", d, val, err)
		}
	})
	t.Run("len=6", func(t *testing.T) {
		d := []byte{1, 2, 3, 4, 5, 6}
		val, err := decodeFixedLengthInteger(d)
		if val != 6618611909121 || err != nil {
			t.Fatalf("Fail: %v -> (%d, %v)", d, val, err)
		}
	})
	t.Run("len=7", func(t *testing.T) {
		d := []byte{1, 2, 3, 4, 5, 6, 7}
		val, err := decodeFixedLengthInteger(d)
		if val != 0 || err == nil {
			t.Fatalf("Fail: %v -> (%d, %v)", d, val, err)
		}
	})
	t.Run("len=8", func(t *testing.T) {
		d := []byte{1, 2, 3, 4, 5, 6, 7, 8}
		val, err := decodeFixedLengthInteger(d)
		if val != 578437695752307201 || err != nil {
			t.Fatalf("Fail: %v -> (%d, %v)", d, val, err)
		}
	})
	t.Run("len=9", func(t *testing.T) {
		d := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
		val, err := decodeFixedLengthInteger(d)
		if val != 0 || err == nil {
			t.Fatalf("Fail: %v -> (%d, %v)", d, val, err)
		}
	})

}
