package bim

import "crypto/sha1"

// Checksum is a placeholder for a SHA-1 checksum.
type Checksum [sha1.Size]byte

// HashSumToChecksum converts a 20-byte slice to a Checksum.
func HashSumToChecksum(p []byte) Checksum {
	temp := Checksum{}
	copy(temp[:], p[0:20])
	return temp
}
