package bim

import "crypto/sha1"

// Blob is a bunch of data.
type Blob []byte

// Checksum calculates the SHA-1 checksum for the blob.
func (blob Blob) Checksum() Checksum {
	h := sha1.New()
	h.Write([]byte(blob))
	temp := h.Sum(nil)

	checksum := [20]byte{}
	copy(checksum[:], temp[0:sha1.Size])
	return checksum
}
