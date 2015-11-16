package md5forfile

import (
	"crypto/md5"
	"io"
	"math"
	"os"
)

const fileChunk = 8192

//FileHash struct for MD5's
type FileHash struct {
	Filename string
	Size     int64
	MD5      []byte
}

//Get takes a file and returns the FileHash struct
func Get(fileLocation string) (*FileHash, error) {
	var fileHash = FileHash{}

	file, err := os.Open(fileLocation)
	if err != nil {
		return &fileHash, err
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	fileHash.Filename = fileInfo.Name()
	fileHash.Size = fileInfo.Size()

	hash := md5.New()

	blocks := uint64(math.Ceil(float64(fileHash.Size) / float64(fileChunk)))

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(fileChunk, float64(fileHash.Size-int64(i*fileChunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	file.Close()

	fileHash.MD5 = hash.Sum(nil)

	return &fileHash, nil
}
