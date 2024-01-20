package main

import (
	"fmt"

	"github.com/Infatzmut/cryptip/encrypt"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
}
