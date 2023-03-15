package list_tool

import (
	"crypto/rand"
	"math/big"
)

func Get_random_element_from_list[T any](list []T) any {
	// it may return nil if it can't find a random source
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		return nil
	}
	n := nBig.Int64()
	return list[n]
}
