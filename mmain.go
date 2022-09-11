package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	v := "value"
	s := sha256.New()
	s.Write([]byte(v))
	fmt.Println(hex.EncodeToString(s.Sum(nil)))
}
