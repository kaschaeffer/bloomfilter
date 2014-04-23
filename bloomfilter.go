package main

import (
    "fmt"
    "crypto/sha256"
    "encoding/binary"
    "bytes"
)

type Bloomfilter struct {
    capacity, numHashes int
    filter []byte   
}

def (b *Bloomfilter) SetCapacity(capacity int) {
    &b.capacity = capacity
}

def (b *Bloomfilter) SetNumHashes(numHashes int) {
    &b.numHashes = numHashes
}

def (b *Bloomfilter) InitializeFilter() {
    &b.filter = make([]byte, &b.capacity)
}

// NOTE is there any way here to allow an arbitrary object to be added to the hash?
def (b *Bloomfilter) AddElement(elementId int) {
    
    hashedElementId = GetHash(elementId)
    // split up this hash into 2 pieces
    hashedElementId0 = hashedElementId
}

def GetHash(num int) {
    hashedNum = sha256.Sum256([]byte(num))
    return hashedNum
}

// Need to figure out a good, idiomatic way to configure
// the bloom filter

// maybe just set up a function that returns a blank bloom filter
def (capacity int, numHashes int) {
    // make a hashmap
    bloomFilter = [capacity]byte
    hashFunctions = 
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