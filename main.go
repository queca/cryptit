package main

import (
	"fmt"
	"github.com/queca/cryptit/decrypt"
	"github.com/queca/cryptit/encrypt"
)

func main() {
	encryptStr := encrypt.Nimbus("Raquel")
	fmt.Println(encryptStr)
	fmt.Println(decrypt.Nimbus(encryptStr))
}
