package main

import (
	"testing"
)

var setBitInByteTestValues = []struct {
	in         byte
	inWhichBit uint
	out        byte // TODO: is the variable name 'out' confusing here since SetbitInByte is in-place?
}{
	{0x00, 0, 0x01},
	{0x00, 1, 0x02},
	{0x00, 2, 0x04},
	{0x00, 3, 0x08},
	{0x00, 4, 0x10},
	{0x00, 5, 0x20},
	{0x00, 6, 0x40},
	{0x00, 7, 0x80},
	{0x01, 0, 0x01}, // setting a bit that's already 1 should have no effect
	{0x01, 1, 0x03},
	{0x03, 1, 0x03},
}

func TestSetBitInByte(t *testing.T) {
	for i := 0; i < len(setBitInByteTestValues); i++ {
		testValues := setBitInByteTestValues[i]

		in := testValues.in // make a copy since setByteInBit mutates testValues.in
		if setBitInByte(&testValues.in, testValues.inWhichBit); testValues.in != testValues.out {
			t.Errorf("setByteInBit(%#02x, %d) = %#02x, want %#02x", in, testValues.inWhichBit, testValues.in, testValues.out)
		}
	}
}

func TestSetBitInByteIncorrectBitPosition(t *testing.T) {
	var bitPosition uint
	var in byte

	bitPosition = 9
	in = 0x00
	err := setBitInByte(&in, bitPosition)

	if err == nil {
		t.Errorf("setByteInBit(%#02x, %d) should return error, but did not", in, bitPosition)
	}
}
