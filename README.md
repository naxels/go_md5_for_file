# MD5 grabber for file written in Golang

Written to use as a package in your app to concurrently get the MD5 hash of a file location

This package will:
* Open the file
* Grab the MD5 of the file
* Return both the filename, size and MD5 of the file

The block calculation code piece comes from Socket Loop
