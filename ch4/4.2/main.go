package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	algorithm := flag.String("algorithm", "sha256", "Hash algorithm to use (sha256, sha384 or sha512)")
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading standard input:", err)
		return
	}

	switch *algorithm {
	case "sha256":
		hash := sha256.Sum256(bytes)
		fmt.Printf("%x\n", hash)
	case "sha384":
		hash := sha512.Sum384(bytes)
		fmt.Printf("%x\n", hash)
	case "sha512":
		hash := sha512.Sum512(bytes)
		fmt.Printf("%x\n", hash)
	default:
		fmt.Println("Invalid hash algorithm specified.")
	}
}
