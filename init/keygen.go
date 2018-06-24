package init

import (
	"crypto/rand"
	"fmt"
  "os"
)

//Key generation here
func genkey() (string){
  Length := 16 //Key length in Byte
	Key := make([]byte, Length)
	_, err := rand.Read(Key)

  if err != nil {
		fmt.Println("error:", err)
	}
  return fmt.Sprintf("%x\n",Key)
}

func Keygen(path string){
	Key1 := genkey()
	Key2 := genkey()
  Key3 := genkey()

	//File management
  f, err := os.Create(path)

  if err != nil {
      panic(err)
  }

  defer f.Close()

  n1, err := f.WriteString(Key1)
  fmt.Printf("wrote %d bytes\n", n1)

  n2, err := f.WriteString(Key2)
  fmt.Printf("wrote %d bytes\n", n2)

  n3, err := f.WriteString(Key3)
  fmt.Printf("wrote %d bytes\n", n3)

  f.Sync()
}

//func main() {
	//Keygen("dat2")

//}
