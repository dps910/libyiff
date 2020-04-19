/*
/	libyiff
/
/	Go library for identifying and sorting furry porn
/   written by heyitspuggo with help from spotlight
/   licensed under GNU Public License version 3
/
/   ....don't ask......
*/

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// Options struct forms the options for Google Cloud AI and the e621 API key
type Options struct {
	GCPToken string
	e621APIkey string
	species []string
}

// Identify takes in an instance of Options and a path to a file and returns 'true' if operation to identify
// character and move image to proper folder is successful. Otherwise, returns the error it ran into.
func Identify(o Options, filePath string) string {

}

// md5Hash does *exactly* what you think it does.
func md5Hash(filePath string) string {
	// file hashing
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		// every day i imagine a future where Go had universal error handlers that didn't require
		// this snippet after every function; but alas, it is the meme of the language
		return returnMD5String
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String
	}

	// type correction memes
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	fmt.Println("Got hash for" + filePath + ": " + returnMD5String)
	return returnMD5String
}
