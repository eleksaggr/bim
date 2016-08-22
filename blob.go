package bim

import "crypto/sha1"

// Blob is a bunch of data.
type Blob []byte

// Checksum calculates the checksum for the blob.
func (blob Blob) Checksum() []byte {
	return sha1.New().Sum([]byte(blob))
}
