# MD5 grabber for file written in Golang

Written to use as a package in your app to concurrently get the MD5 hash of a file location

This package will:
* Open the file
* Grab the MD5 of the file
* Return both the filename, size and MD5 of the file

The block calculation code piece comes from Socket Loop

## Installation
```
go get github.com/naxels/go_md5_for_file
```
## Basic Usage to md5 folder (nonconcurrent)
```
import "github.com/naxels/go_md5_for_file"

const fileFolder = "(folder)"

func main() {
	files, _ := ioutil.ReadDir(fileFolder)

	for _, f := range files {
		result, err := md5forfile.Get(fileFolder + string(os.PathSeparator) + f.Name())
		if err != nil {
			fmt.Print("Error")
		}

		fmt.Printf("%v, %x\n", result.Filename, result.MD5)
	}
}
```
