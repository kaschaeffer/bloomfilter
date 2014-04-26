package main

import (
    "fmt"
    "crypto/sha256"
    "encoding/binary"
    "bytes"
)

// constructor
func NewBloomFilterStringKeyed(capacity int, numHashes int) *BloomFilterStringKeyed {
    bloomFilter := new(BloomFilterStringKeyed)
    bloomFilter.capacity = capacity
    bloomFilter.numHashes = numHashes
    bloomFilter.filter = make([]byte, capacity)
    return bloomFilter   
}

type BloomFilterStringKeyed struct {
    capacity, numHashes int
    filter []byte
}


func (b *BloomFilterStringKeyed) AddKey(key string) {
    // get a 32-byte hash of the key
    hashedKey := HashString(key)
    fmt.Println(hashedKey)

    // now convert this to something that hashes to appropriate-sized filter
    // do something with the key...
}

/////////////////////////////////////////////////
//       general purpose hashing functions     //
/////////////////////////////////////////////////

func HashString(key string)([32]byte) {
    byteKey := StringToByteArray(key)
    hashedKey := HashByteArray(byteKey)
    return hashedKey
}

func StringToByteArray(str string)([]byte) {
    return []byte(str)
}

// general function that will be used for all hashing
func HashByteArray(byteArray []byte)([32]byte) {
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

}