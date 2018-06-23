package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
  "os"
  "bytes"
  "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	// Load your secret key from a safe place and reuse it across multiple
  f, err := os.Open("dat2")
  check(err)

  o3, err := f.Seek(66, 0)
  check(err)
  b3 := make([]byte, 33)
  n3, err := f.Read(b3)
  check(err)
  fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

  _, err = f.Seek(0, 0)
  check(err)

  f.Close()

	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString(string(b3))
	plaintext, err := ioutil.ReadFile("1")
  check(err)

	block, err := aes.NewCipher(key)
	check(err)

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

  //Create file to store SSE file
  g, err := os.Create("a_aes.txt")
  check(err)

  _, err = io.Copy(g, bytes.NewReader(ciphertext))
  check(err)
	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
}
