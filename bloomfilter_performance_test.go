// goal here is to create tools that can asses the false positive rates
// and other performance characteristics of bloom filters

// from theory, we know that for a table with:
//      m bits 
//      k hashfunctions
//      n keys randomly inserted
//
// the false positive rate should be (1 - (1-(1/m))^kn)^k ~= (1- exp(kn/m))^k

func randomString() (random string) {
    // TODO
    // idea is to randomly choose a character and then flip a coin to decide whether
    // to add an additional character
    //
    // probably want to have a high probability of going on to get high entropy here!
    // 
    // alternative would be to sample from strings of a fixed, but very large length
}

func randomStrings(length int) (randoms []string) {
    for i:=0; i<length; i++ {
        randoms[i] = randomString()
    }
}

func computeFalsePositiveRate(b *BloomFilter, existingKeys []string, iterations int) (falsePositiveRate float64) {
    existingKeysSet := make(map[string]bool)
    for i:=0; i<len(existingKeys); i++ {
        existingKeysSet[existingKeys[i]] = true
    }

    newKeysChecked := float64(0)
    falsePositives := float64(0)

    randomStringKeys := randomStrings(iterations)
    for i:=0; i<iterations; i++ {
        key := randomStringKeys[i]
        if !existingKeysSet[key] {
            if b.queryKey(key) {
                falsePositives += 1
            }
            newKeysChecked += 1
        }
    }

    falsePositiveRate = falsePositives/newKeysChecked
    return
}