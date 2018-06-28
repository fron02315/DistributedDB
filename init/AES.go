package init

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
  "os"
  "bytes"
  "io/ioutil"
	"errors"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readkey()([]byte){
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

  key, _ := hex.DecodeString(string(b3))

  return key
}

func Encryptfile(infile string, outfile string){
  key := readkey()
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	plaintext, err := ioutil.ReadFile(infile)
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
  g, err := os.Create(outfile)
  check(err)

  _, err = io.Copy(g, bytes.NewReader(ciphertext))
  check(err)
	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
}

func Decryptfile(infile string, outfile string){
  key := readkey()

  ciphertext, err  := ioutil.ReadFile(infile)
  check(err)

	block, err := aes.NewCipher(key)
  check(err)

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

  //Create file to store SSE file
  g, err := os.Create(outfile)
  check(err)

  _, err = io.Copy(g, bytes.NewReader(ciphertext))
  check(err)
}

func keykeyword()([]byte,[]byte){
  // Load your secret key from a safe place and reuse it across multiple
  f, err := os.Open("dat2")
  check(err)

	//read key for keyword1
  o3, err := f.Seek(0, 0)
  check(err)
  b3 := make([]byte, 33)
  n3, err := f.Read(b3)
  check(err)
  fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

  _, err = f.Seek(0, 0)
  check(err)

	key1, _ := hex.DecodeString(string(b3))

	//read key for keyword2
	o2, err := f.Seek(33, 0)
	check(err)
	b2 := make([]byte, 33)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	_, err = f.Seek(0, 0)
	check(err)

	 key2, _ := hex.DecodeString(string(b2))

  f.Close()
  return key1, key2
}

func KeyGenBitswap(keyword string)(string, string){
	//retrieve key
	key1, key2 := keykeyword()

	keywordkey1 := string(key1)
	keywordkey2 := string(key2)

	//Hash for token
	output := fmt.Sprintf(keywordkey1 , keyword)
	hasher := md5.New()
  hasher.Write([]byte(output))
  token := hex.EncodeToString(hasher.Sum(nil))
	//fmt.Println(token)

	//Hash for gamma
	output2 := fmt.Sprintf(keywordkey2, keyword)
	hasher2 := md5.New()
	hasher2.Write([]byte(output2))
	gamma := hex.EncodeToString(hasher2.Sum(nil))
	//fmt.Println(gamma)

	// Load your secret key from a safe place and reuse it across multiple
  // NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
 	// package like bcrypt or scrypt.

	return token, gamma
}

func Encryptkeyword(message string, keyword string)(string, string, string) {
	token, gamma := KeyGenBitswap(keyword)

	key := []byte(gamma)
	plainText := []byte(message)
	block, err := aes.NewCipher(key)
	if err != nil {

	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {

	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess := base64.URLEncoding.EncodeToString(cipherText)
	return encmess , token, gamma
}

func Decryptkeyword(key []byte, securemess string) (string) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {

	}

	block, err := aes.NewCipher(key)
	if err != nil {

	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")

	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	decodedmess := string(cipherText)
	return decodedmess
}

//func main() {

	//Call function
	//Filename := "dat2"
	//initiate.Keygen(Filename)

	//Encryptfile("1","a_aes.txt")
	//Decryptfile("a_aes.txt","test.txt")

	//Keyword encrypytion
	/*cipher, token, gamma := Encryptkeyword(bitmap, keyword)
	fmt.Println(cipher)
	fmt.Println(keyword)
	fmt.Println(token)
	fmt.Println(gamma) */

	//This section is keyword decryption-> take bitmap to retrieve the files
	/*key := []byte(gamma)
	word := Decryptkeyword(key,cipher)
	fmt.Println(word)*/

//}
