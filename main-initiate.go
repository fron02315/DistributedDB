package main

import (
	initiate "./init"

	"fmt"
	//"crypto/aes"
	//"crypto/cipher"
	//"crypto/rand"
	//"encoding/hex"
	//"crypto/md5"
	//"encoding/base64"
	//"fmt"
	//"io"
	//"os"
	//"bytes"
	//"io/ioutil"
	//"errors"
)





func main() {
	keyword := "Fon"
	//bitmap := "10001000"

	//Keygenerate and collect on filename
	/*Filename := "dat2" //dat2 is key stored  file
	initiate.Keygen(Filename)*/

	//This section is file encryption and decryption
	//Requirement: file name "1"
  /*initiate.Encryptfile("1","a_aes.txt")
  initiate.Decryptfile("a_aes.txt","test.txt")*/

	//Keyword encrypytion
	/*cipher, token, gamma := initiate.Encryptkeyword(bitmap, keyword)
	fmt.Println(cipher)
	fmt.Println(keyword)
	fmt.Println(token)
	fmt.Println(gamma) */

	//insert into database
	/*db := initiate.Openconnection()
	initiate.InsertKey(db, token, cipher)*/

	//Create token for Search
	token, gamma := initiate.KeyGenBitswap(keyword)

	//search for the specific keyword
	
	cipher := "CDtXp3pU2EhRlh2NOlHJ4pmJ706qmLmK"

	//This section is keyword decryption-> take bitmap to retrieve the files
	key := []byte(gamma)
	word := initiate.Decryptkeyword(key,cipher)


}
