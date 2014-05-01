package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// possibly pass in the hash function too??
type BloomFilter struct {
	byteCapacity, numHashes int
	bitHashTable            []byte
	convertKeyToByteArray	func(interface{}) []byte
}

// constructor
func NewBloomFilterStringKeyed(byteCapacity int, numHashes int) *BloomFilter {
	return &BloomFilter{
		byteCapacity: byteCapacity,
		numHashes: numHashes,
		bitHashTable: make([]byte, byteCapacity),
		convertKeyToByteArray: func(key interface{}) []byte {
			return []byte(key.(string))
		},
	}
}


// IDEA: maybe pass in which hashing algorithm you want in the constructor
// (are functions first class in golang?)

// Somewhere there should be appropriate error handling if filter is too big
// for the hash functions...

func (b *BloomFilter) QueryKey(key interface{}) (keyFound bool) {
	hashes := b.generateHashes(key)
	indices := b.convertHashesToCorrectRange(hashes)
	bits := b.getBitsFromIndices(indices)
	
	for i:=0; i<b.numHashes; i++ {
		if !bits[i] {
			keyFound = false
			return
		}
	}
	keyFound = true
	return
}

func (b *BloomFilter) AddKey(key interface{}) {
	hashes := b.generateHashes(key)
	indices := b.convertHashesToCorrectRange(hashes)
	b.setBitsFromIndices(indices)
}

func (b *BloomFilter) setBitsFromIndices(indices []uint64) {
    // check length of the array

    // this should be an atomic operation (how to do??)

    // could also be parallelized via goroutines potentially
    for i:=0; i<b.numHashes; i++ {
    	b.setBitFromIndex(indices[i])
    }
}

func (b *BloomFilter) setBitFromIndex(index uint64) {
	arrayIndex := index / 8
	bitPosition := uint(index % 8)
	setBitInByte(&b.bitHashTable[arrayIndex], bitPosition)
}

func (b *BloomFilter) getBitsFromIndices(indices []uint64) (bits []bool) {
    // check length of the array

    // this should be an atomic operation (how to do??)

    // could also be parallelized via goroutines potentially
    bits = make([]bool, b.numHashes)

    for i:=0; i<b.numHashes; i++ {
    	bits[i] = b.getBitFromIndex(indices[i])
    }
    return
}

func (b *BloomFilter) getBitFromIndex(index uint64)(bit bool) {
	arrayIndex := index / 8
	bitPosition := uint(index % 8)
	bit, _ = getBitInByte(&b.bitHashTable[arrayIndex], bitPosition)
	return

}

func (b *BloomFilter) convertHashesToCorrectRange(hashes [][]byte) []uint64 {

	// TODO first should check the length of the array

	indices := make([]uint64, b.numHashes)

	for i := 0; i < b.numHashes; i++ {
		hashAsInt := convertByteArraytoUInt64(hashes[i])
		indices[i] = hashAsInt % uint64(b.byteCapacity)
	}

	return indices
}

func convertByteArraytoUInt64(byteArray []byte) uint64 {
	convertedByteArray := binary.LittleEndian.Uint64(byteArray)
	return convertedByteArray
}

func (b *BloomFilter) HashKey(key interface{}) [32]byte {
	byteKey := b.convertKeyToByteArray(key)
	hashedKey := HashByteArray(byteKey)
	return hashedKey
}

func (b *BloomFilter) generateHashes(key interface{}) [][]byte {
	// get a 32-byte hash of the key
	hashedKey := b.HashKey(key)
	fmt.Println(hashedKey)

	// call this "seedHash" instead?

	// split into two 8-byte hashes
	// TODO: don't hardcode 8 below..
	hashedKey0 := hashedKey[:8]
	hashedKey1 := hashedKey[8:16]

	// max filter size should be 2^(16*8)

	fmt.Printf("hashedKey0 = %d\n", hashedKey0)
	fmt.Printf("hashedKey1 = %d\n", hashedKey1)

	// list of hashes
	fmt.Printf("numHashes = %d\n", b.numHashes)
	hashes := make([][]byte, b.numHashes)

	for k := 0; k < b.numHashes; k++ {
		kByte := byte(k)

		hashes[k] = make([]byte, 8)
		for i := 0; i < len(hashedKey0); i++ {
			// hashes[k][i] = hashedKey0[i] + (hashedKey1[i]*byte(k))
			hashes[k][i] = hashedKey0[i] ^ (hashedKey1[i] * kByte)

		}
	}
	return hashes
}

// func byteMultiplication(b byte, k int) {
//     b
// }

// func byteArrayToInt64(inputBytes *[8]byte) {
//     var outputInt64 int64
//     buf := bytes.NewBuffer(inputBytes)
//     binary.Read(buf, binary.LittleEndian, &outputInt64)
//     return outputInt64
// }

// Want function that takes (x is int ) and returns a flips one bit in a byte (8 bits)
// **should be able to do this via bit-shifting tricks**

func setBitInByte(b *byte, bitPosition uint) (err error) {
	if bitPosition > 7 {
		err = fmt.Errorf(
			"Position of bit to set must be in the range [0,7), attempted to set position %", bitPosition)
	}
	*b = *b | (1 << bitPosition)
	return
}

func getBitInByte(b *byte, bitPosition uint) (bit bool, err error) {
	if bitPosition > 7 {
		err = fmt.Errorf(
			"Position of bit to set must be in the range [0,7), attempted to set position %", bitPosition)
	}
	byteCheck := *b & (1 << bitPosition)
	bit = (byteCheck != 0x00)
	return
}

/////////////////////////////////////////////////
//       general purpose hashing functions     //
/////////////////////////////////////////////////


// general function that will be used for all hashing
func HashByteArray(byteArray []byte) [32]byte {
	hash := sha256.Sum256(byteArray)
	return hash
}

// Testing out how to hash some object
func main() {
	fmt.Println("Hello, playground")

	// sha := sha256.New()

	result := sha256.Sum256([]byte("foobar"))
	length := len(result)
	fmt.Printf("%d\n", length)
	fmt.Printf("%x\n", result)

	part0 := result[:16]
	fmt.Printf("%d\n", len(part0))
	fmt.Printf("%x\n", part0)

	// converting byte array to signed int
	buf := bytes.NewBuffer(part0)
	intPart0, _ := binary.ReadVarint(buf)
	// note: probably want unsigned int instead...

	fmt.Printf("%d\n\n\n", intPart0)

	part1 := result[16:]
	fmt.Printf("%d\n", len(part1))
	fmt.Printf("%x\n", part1)

	b := NewBloomFilterStringKeyed(32, 10)

	fmt.Println(b.byteCapacity)
	fmt.Println(b.numHashes)
	fmt.Println(b.bitHashTable)

	b.AddKey("foo")

}
