package main
// goal here is to create tools that can asses the false positive rates
// and other performance characteristics of bloom filters

// from theory, we know that for a table with:
//      m bits 
//      k hashfunctions
//      n keys randomly inserted
//
// the false positive rate should be (1 - (1-(1/m))^kn)^k ~= (1- exp(kn/m))^k

import (
    "testing"
    "bytes"
    "math/rand"
    "fmt"
    "time"
)

func randomCharacter() (randomCharacter string) {
    randomCharacter = string(rand.Intn(256))
    return
}

func randomString(length int) (random string) {
    var buffer bytes.Buffer

    for i:=0; i<length; i++ {
        buffer.WriteString(randomCharacter())
    }

    random = buffer.String()
    return
}

func randomStrings(numStrings, stringLength int) (randoms []string) {
    randoms = make([]string, numStrings)
    for i:=0; i<numStrings; i++ {
        randoms[i] = randomString(stringLength)
    }
    return
}

func computeFalsePositiveRate(b *BloomFilter, existingKeys []string, stringLength, iterations int) (falsePositiveRate float64) {
    existingKeysSet := make(map[string]bool)
    for i:=0; i<len(existingKeys); i++ {
        existingKeysSet[existingKeys[i]] = true
    }

    newKeysChecked := float64(0)
    falsePositives := float64(0)

    randomStringKeys := randomStrings(iterations, stringLength)
    for i:=0; i<iterations; i++ {
        key := randomStringKeys[i]
        if !existingKeysSet[key] {
            if b.QueryKey(key) {
                falsePositives += 1
            }
            newKeysChecked += 1
        }
    }

    falsePositiveRate = falsePositives/newKeysChecked
    return
}

func TestFalsePositiveRate(t *testing.T) {
    // todo
    rand.Seed(time.Now().UTC().UnixNano())

    stringLength := 512
    numKeysAdded := 20
    capacity := 512
    numHashes := 5
    iterations := 10000

    keysAdded := randomStrings(numKeysAdded, stringLength)

    b := NewBloomFilterStringKeyed(capacity, numHashes)
    for i:=0; i<numKeysAdded; i++ {
        b.AddKey(keysAdded[i])
    }

    fpr := computeFalsePositiveRate(b, keysAdded, stringLength, iterations)
    fmt.Printf("The false positive rate is %f\n", fpr)
}