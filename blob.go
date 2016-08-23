package bim

import "crypto/sha1"

// Blob is a bunch of data.
type Blob []byte

// Checksum calculates the SHA-1 checksum for the blob.
func (blob Blob) Checksum() [20]byte {
	h := sha1.New()
	h.Write([]byte(blob))
	return sha1.Sum(nil)
}
