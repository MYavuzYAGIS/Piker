package digest5

// TODO -> Read file path from stdin
// TODO -> Read string from stdin
// TODO -> Check file extension
// TODO -> maybe a better error handling
// TODO -> enable Flags and coloring schemes
// TODO -> enable online lookup

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Md5StringHasher creates a md5 hash of a given string as argument
func Md5StringHasher(input string) string {
	hash := md5.Sum([]byte(input))
	hashed := hex.EncodeToString(hash[:])
	fmt.Println(hashed)
	return hashed
}

// take a file and hash it completely, not EACH Line!
func Md5FileHasher() {
}
