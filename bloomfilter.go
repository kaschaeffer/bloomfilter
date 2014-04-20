package main

import (
    "fmt"
    "crypto/sha256"
)

type Bloomfilter struct {
    capacity, numHashes int   
}

def (b *Bloomfilter) SetCapacity(capacity int) {
    &b.capacity = capacity
}

def (b *Bloomfilter) SetNumHashes(numHashes int) {
    &b.numHashes = numHashes
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
    fmt.Printf("%x", result)
}