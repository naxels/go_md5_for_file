package md5forfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const (
	testDataDir = "testdata"
)

// hashes generated from MacOS md5 app
var m = map[string]string{
	"1.txt":               "a8ae2f4a56baf78845c041c833946d00",
	"2.txt":               "45bfae8da88f82876f4da0bb6f329b28",
	"3.txt":               "973fafae06f0e850ec6017bd0f52ab83",
	"4.txt":               "edc715389af2498a623134608ba0a55b",
	"quick-brown-fox.txt": "37c4b87edffc5d198ff5a185cee7ee09",
}

func TestMD5singleFile(t *testing.T) {
	fileName := "quick-brown-fox.txt"
	fileLocation := testDataDir + string(os.PathSeparator) + fileName

	result, err := Get(fileLocation)
	if err != nil {
		t.Fatalf("Get(%q) err = %v, expected nil", fileName, err)
	}
	if fmt.Sprintf("%x", result.MD5) != m[fileName] {
		t.Fatalf("Get(%q) hash = %x, expected %v", fileName, result.MD5, m[fileName])
	}
}

func TestMD5multiFile(t *testing.T) {
	files, _ := ioutil.ReadDir(testDataDir)
	for _, f := range files {
		fileLocation := testDataDir + string(os.PathSeparator) + f.Name()
		result, err := Get(fileLocation)
		if err != nil {
			t.Fatalf("Get(%q) err = %v, expected nil", f.Name(), err)
		}
		if fmt.Sprintf("%x", result.MD5) != m[f.Name()] {
			t.Fatalf("Get(%q) hash = %x, expected %v", f.Name(), result.MD5, m[f.Name()])
		}
	}
}
