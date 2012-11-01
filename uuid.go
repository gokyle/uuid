package uuid

import (
	"crypto/rand"
	"math/big"
)

var octet_space = []byte("ABCDEF0123456789")
type Uuid []byte

func randomSymbol(keyspace []byte) (b byte, err error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(keyspace))))
	if err == nil {
		index := n.Int64()
		b = keyspace[index]
	}
	return
}

// GenerateV4 returns a version 4 UUID as a byte slice.
func GenerateV4() (u Uuid, err error) {
	u = make([]byte, 36)
	var y byte

	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			u[i] = "-"[0]
		} else if i == 14 || i == 19 {
			continue
		} else {
			u[i], err = randomSymbol(octet_space)
			if err != nil {
				return
			}
		}
	}
	u[14] = "4"[0]
	y, err = randomSymbol([]byte("89AB"))
	if err == nil {
		u[19] = y
	}
	return
}

// GenerateV4String returns a version 4 UUID as a string.
func GenerateV4String() (u string, err error) {
	var uByte []byte
	uByte, err = GenerateV4()
	if err == nil {
		u = string(uByte)
	}
	return
}
