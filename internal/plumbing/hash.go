package plumbing

import (
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

func HashFile(filename string, write bool) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	hashHex, store := HashObject("blob", data)
	if write {
		dirname := hashHex[:2]
		hashedFilename := hashHex[2:]
		hashedPath := filepath.Join(".git.local", "objects", dirname, hashedFilename)
		os.MkdirAll(filepath.Dir(hashedPath), 0o755)
		f, _ := os.Create(hashedPath)
		defer f.Close()

		w := zlib.NewWriter(f)
		defer w.Close()
		w.Write(store)
	}
	fmt.Println(hashHex)
}

func HashObject(headerType string, data []byte) (string, []byte) {
	header := fmt.Sprintf("%s %d\u0000", headerType, len(data))
	store := append([]byte(header), data...)
	hash := sha1.Sum(store)
	hashHex := hex.EncodeToString(hash[:])

	return hashHex, store
}
