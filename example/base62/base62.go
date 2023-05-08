package main

import (
	"fmt"

	"github.com/tuannguyenandpadcojp/utils/base62"
)

func main() {
	id := int64(755129780)
	encodedID := base62.Encode(id)
	// encode value:
	fmt.Printf("encoded: %s\n", encodedID)

	// decode value:
	decodedID, _ := base62.Decode(encodedID)
	fmt.Printf("decoded: %d\n", decodedID)
}
