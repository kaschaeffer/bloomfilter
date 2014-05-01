package main

import (
	"testing"
	"bytes"
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
			t.Errorf("setBitInByte(%#02x, %d) = %#02x, want %#02x", in, testValues.inWhichBit, testValues.in, testValues.out)
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
		t.Errorf("setBitInByte(%#02x, %d) should return error, but did not", in, bitPosition)
	}

	bitPosition = 7
	in = 0x00
	err = setBitInByte(&in, bitPosition)

	if err != nil {
		t.Errorf("setBitInByte(%#02x, %d) should not return error, but did", in, bitPosition)
	}
}

func TestNewBloomFilterStringKeyed(t *testing.T) {
	byteCapacity := 100
	numHashes := 5
	bf := NewBloomFilterStringKeyed(byteCapacity, numHashes)
	if bf.byteCapacity != byteCapacity {
		t.Errorf("NewBloomFilterStringKeyed(%d, %d) returned object, b, with incorrect b.byteCapacity = %d", byteCapacity, numHashes, bf.byteCapacity)
	}
	if bf.numHashes != numHashes {
		t.Errorf("NewBloomFilterStringKeyed(%d, %d) returned object, b, with incorrect b.numHashes = %d", byteCapacity, numHashes, bf.numHashes)
	}

	defaultBitHashTable := make([]byte, byteCapacity)
	for i:=0; i<len(defaultBitHashTable); i++ {
		defaultBitHashTable[i] = 0x00
	}
	if !bytes.Equal(bf.bitHashTable, defaultBitHashTable) {
		t.Errorf("NewBloomFilterStringKeyed(%d, %d) returned object, b, with incorrect b.bitHashTable = %x", byteCapacity, numHashes, bf.bitHashTable)
	}
}

var setBitFromIndexTestValues = []struct {
	bitToSet         		uint64
	expectedArrayIndex		int 
	expectedModifiedByte    byte // TODO: is the variable name 'out' confusing here since SetbitInByte is in-place?
}{
	{0, 0, 0x01},
	{1, 0, 0x02},
	{2, 0, 0x04},
	{3, 0, 0x08},
	{8, 1, 0x01},
	{9, 2, 0x02},
	{15, 2, 0x80},
}

func TestSetBitFromIndex(t *testing.T) {
	// initialize empty bloom filter (maybe this should be its own function?)
	byteCapacity := 100
	numHashes := 5
	bf := NewBloomFilterStringKeyed(byteCapacity, numHashes)

	// set a bit
	arrayIndex := 0
	var modifiedByte byte
	var bitToSet uint64

	bitToSet = 0
	modifiedByte = 0x01

	// maybe should factor out this boilerplait...
	expectedBitHashTable := make([]byte, byteCapacity)
	for i:=0; i<len(expectedBitHashTable); i++ {
		expectedBitHashTable[i] = 0x00
	}	
	expectedBitHashTable[arrayIndex] = modifiedByte

	bf.setBitFromIndex(bitToSet)

	if !bytes.Equal(bf.bitHashTable, expectedBitHashTable) {
		t.Errorf("Attempted to set bit %d.  Expected array %x, returned array %x", bitToSet, bf.bitHashTable, expectedBitHashTable)
	}

	// TODO
}

var getBitInByteTestValues = []struct {
	in         byte
	inWhichBit uint
	out        bool
}{
	{0x01, 0, true},
	{0x01, 1, false},
	{0x02, 1, true},
	{0x04, 2, true},
	{0x08, 3, true},
	{0x10, 4, true},
	{0x10, 0, false},
	{0x20, 5, true},
	{0x40, 6, true},
	{0x80, 7, true},
	{0x01, 0, true},
	{0x03, 1, true},
	{0x03, 2, false},
}

func TestGetBitInByte(t *testing.T) {
	for i := 0; i < len(getBitInByteTestValues); i++ {
		testValues := getBitInByteTestValues[i]

		if result, _ :=getBitInByte(&testValues.in, testValues.inWhichBit); result != testValues.out {
			t.Errorf("getBitInByte(%#02x, %d) = %t, want %t", testValues.in, testValues.inWhichBit, result, testValues.out)
		}
	}
}

func TestGetBitInByteIncorrectBitPosition(t *testing.T) {
	var bitPosition uint
	var in byte

	bitPosition = 9
	in = 0x00
	_, err := getBitInByte(&in, bitPosition)

	if err == nil {
		t.Errorf("setBitInByte(%#02x, %d) should return error, but did not", in, bitPosition)
	}

	bitPosition = 7
	in = 0x00
	_, err = getBitInByte(&in, bitPosition)

	if err != nil {
		t.Errorf("setBitInByte(%#02x, %d) should not return error, but did", in, bitPosition)
	}
}

var addKeyValues = []struct {
	byteCapacity         int
	numHashes int
	key        string
}{
	{10000, 3, "foobar"},
	{10000, 4, "i am a test"},
	{10000, 5, "whatever"},
	{10000, 6, "hope this test passes!!!!!!!!"},
}

func TestAddKey(t *testing.T) {
	for i:=0; i<len(addKeyValues); i++ {
		testValues := addKeyValues[i]
		b := NewBloomFilterStringKeyed(testValues.byteCapacity, testValues.numHashes)
		
		b.AddKey(testValues.key)
		numBitsSet := countBitsInBytes(&b.bitHashTable)
		
		if numBitsSet != testValues.numHashes {
			t.Errorf("addKey(%s) should set %d bits, but only set %d bits", testValues.key, testValues.numHashes, numBitsSet)
		}
	}
}

var queryKeyValues = []struct {
	byteCapacity         int
	numHashes int
	addKey        string
	queryKey 	string
	result		bool
}{
	{10000, 3, "foobar", "foobar", true},
	{10000, 3, "foobar", "fooBar", false},
	{10000, 3, "foobar", "foobar ", false},
	{10000, 3, "foobar", "foobaz", false},
	{10000, 3, "foobar", "lasfdahfak", false},
	{10000, 13, "zzzz", "ZZZZ", false},
	{10000, 13, "zzzz", "zzzz", true},
}

func TestQueryKey(t *testing.T) {
	for i:=0; i<len(queryKeyValues); i++ {
		testValues := queryKeyValues[i]
		b := NewBloomFilterStringKeyed(testValues.byteCapacity, testValues.numHashes)
		
		b.AddKey(testValues.addKey)
		result := b.QueryKey(testValues.queryKey)
		
		if result != testValues.result {
			t.Errorf("Added key %s, then queried key %s.  Expected result to be %t, but returned %t",
				testValues.addKey,
				testValues.queryKey,
				testValues.result,
				result,
				)
		} 
	}
}

// Useful utility method for performing comparisons
func countBitsInBytes(byteArray *[]byte)(bitCount int) {
	for i:=0; i<len(*byteArray); i++ {
		bitCount += countBitsInByte(&(*byteArray)[i])
	}
	return
}

func countBitsInByte(b *byte)(bitCount int) {
	for i:=uint(0); i<8; i++ {
		bitCount += int((*b >> i) & 1)
	}
	return
}
