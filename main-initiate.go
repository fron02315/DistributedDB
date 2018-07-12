package main

import (
	initiate "./init"
	block "./blockchain"

	"fmt"
	"bufio"
	"os"
	"strings"
	"unsafe"
)



func keygen(){
	//If fisrt time == true
	//Keygenerate and collect on filename
	Filename := "dat2" //dat2 is key stored  file
	initiate.Keygen(Filename)
}

func insertIndex(keyword string, bitmap string, bc *block.Blockchain){
	cipher, token, _ := initiate.Encryptkeyword(bitmap, keyword)

	//Connect to database
	db := initiate.Openconnection()

	//insert into database
	initiate.InsertKey(db, token, cipher)

	fmt.Println("Successfully Insert keyword",keyword)

	db.Close()

	//size of data exchange
	msgsize := fmt.Sprint(unsafe.Sizeof(keyword))

	block.AddBlock(bc,"client", "peer", msgsize)
}

func fileDecryption(inFile string, outFile string){
	initiate.Decryptfile(inFile,outFile)
}

func searchKeyword(keyword string){
	//Create token for Search
	token1, gamma1 := initiate.KeyGenBitswap(keyword)

	//Connect to database
	db := initiate.Openconnection()

	//search for the specific keyword
	EnBitmap := initiate.Search(db, token1)

	//This section is keyword decryption-> take bitmap to retrieve the files
	key := []byte(gamma1)
	Bitmap := initiate.Decryptkeyword(key,EnBitmap)
	fmt.Println(Bitmap)
	fmt.Println(unsafe.Sizeof(Bitmap)) //Print the message size

	db.Close()
}

func updateKeyword(keyword string, bitmap string){
	db := initiate.Openconnection()

	initiate.UpdateTable(db, "Egyptian", "00000001")

	db.Close()

}

//CMD function \m/
func getInput() ([]string, error) {
	b := bufio.NewReader(os.Stdin)
	line, err := b.ReadString('\n')
	if err != nil {
		return []string{}, err
	}
	return strings.Split(strings.TrimSpace(line), " "), nil
}

func helpFn(){
	fmt.Println(`Available commands:
Commands
  help         help guide to specify each available commands
  quit         exit the console
  keygen         Key generation- automatically stored on local machine
  blockchain         Blockchain lookup

Initiate commands:
  insertI <keyword,bitmap>  				construct the inverted index
  fileEn <inputfile,outputfile>    	encrypt the target file
  fileDe <inputfile,outputfile>     decrypt the target file

Search commands:
  searchKey <keyword> 	Search the keyword

Update commands:
  updateKey <keyword,bitmap>	Update the bitmap of corresponding keyword.
`)
}

func Call(function string, bc *block.Blockchain) {
	switch {
		case strings.TrimRight(function, "\n") ==  "help": helpFn()
		case strings.TrimRight(function, "\n") ==  "quit": os.Exit(1)
		case strings.TrimRight(function, "\n") ==  "keygen": keygen()
		case strings.TrimRight(function, "\n") ==  "blockchain":
			block.PrintBC(bc)
		case strings.TrimRight(function, "\n") ==  "insertI":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("keyword: ")
			keyword1, _ := reader.ReadString('\n')
			keyword := strings.TrimRight(keyword1, "\n")
			fmt.Print("bitmap: ")
			bitmap1, _ := reader.ReadString('\n')
			bitmap := strings.TrimRight(bitmap1, "\n")
			insertIndex(keyword, bitmap, bc)
		case strings.TrimRight(function, "\n") ==  "fileEn":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Input file: ")
			input1, _ := reader.ReadString('\n')
			input := strings.TrimRight(input1, "\n")
			fmt.Print("Output file: ")
			output1, _ := reader.ReadString('\n')
			output := strings.TrimRight(output1, "\n")
			initiate.Encryptfile(input, output)
		case strings.TrimRight(function, "\n") ==  "fileDe":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Input file: ")
			input1, _ := reader.ReadString('\n')
			input := strings.TrimRight(input1, "\n")
			fmt.Print("Output file: ")
			output1, _ := reader.ReadString('\n')
			output := strings.TrimRight(output1, "\n")
			initiate.Decryptfile(input, output)
		case strings.TrimRight(function, "\n") ==  "searchKey":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("keyword: ")
			keyword1, _ := reader.ReadString('\n')
			keyword := strings.TrimRight(keyword1, "\n")
			searchKeyword(keyword)
		case strings.TrimRight(function, "\n") ==  "updateKey":
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Update Bitmap where Keyword is: ")
			input1, _ := reader.ReadString('\n')
			input := strings.TrimRight(input1, "\n")
			fmt.Print("New Bitmap: ")
			output1, _ := reader.ReadString('\n')
			output := strings.TrimRight(output1, "\n")
			initiate.Decryptfile(input, output)
		default:helpFn()
		}
}

func main() {
	cmd := "DFS-Console> "

	bc := block.NewBlockchain()

	fmt.Println(`Welcome to Our Distributed file storage demonstration. `)
	fmt.Println(`Type "help" to learn the available commands`)
	//helpFn()
	for {
		fmt.Printf(cmd)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
    text, _ := reader.ReadString('\n')

    Call(text,bc)
	}
}
