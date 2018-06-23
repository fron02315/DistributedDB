package main

import (
  //Keygen and print
	"crypto/rand"
	"fmt"
  //Save file
  //"bufio"
//  "io/ioutil"
  "os"
)

//Key generation here
func genkey() ([]byte){
  Length := 16 //Key length in Byte
	Key := make([]byte, Length)
	_, err := rand.Read(Key)

  if err != nil {
		fmt.Println("error:", err)
	}
  return Key
}

//Save key value in string format
func saveprint(Key []byte) (string){
  return fmt.Sprintf("%x\n",Key)
}

func main() {

  Key1 := genkey()
  Print1 := saveprint(Key1)

  Key2 := genkey()
  Print2 := saveprint(Key2)

  Key3 := genkey()
  Print3 := saveprint(Key3)

  //File management
  f, err := os.Create("dat2")

  if err != nil {
      panic(err)
  }

  defer f.Close()

  n1, err := f.WriteString(Print1)
  fmt.Printf("wrote %d bytes\n", n1)

  n2, err := f.WriteString(Print2)
  fmt.Printf("wrote %d bytes\n", n2)

  n3, err := f.WriteString(Print3)
  fmt.Printf("wrote %d bytes\n", n3)

  f.Sync()
}
