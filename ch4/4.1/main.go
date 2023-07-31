package main

import (
	"crypto/sha256"
	"fmt"
)

func bitCountDifferent(hash1, hash2 [32]byte) int {
	count := 0
	for i := 0; i < len(hash1); i++ {
		diffBits := hash1[i] ^ hash2[i]
		count += popCount(diffBits)
	}
	return count
}

// 计算一个字节中不同bit的数目
func popCount(b byte) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		if b&(1<<i) != 0 {
			count++
		}
	}
	return count
}

func main() {
	str1 := "hello"
	str2 := "world"

	hash1 := sha256.Sum256([]byte(str1))
	hash2 := sha256.Sum256([]byte(str2))

	count := bitCountDifferent(hash1, hash2)
	fmt.Printf("The number of different bits is: %d\n", count)
}
