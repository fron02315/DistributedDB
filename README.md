Noted that this text is informal way to describe my project. 
The white paper will be released soon.

# DistributedDB
This project is regarding from my project course. The project proposed to provide search operation on distributed file storage.
Originally, distributed database is not efficient due to communication cost and routing complexity. 

## Getting start

Start with

```Bash
$ go get github.com/fron02315/DistributedDB
```

## Project methodology
The project is divided into three part.
* Initiate the system
* Search the file address
* File retrieval -> FOCUS ON THIS 


```Bash
> Overview
src                 
├── main  # Main source code folder
│   ├── init
│   │    ├── FileEncryption.go
│   │    ├── IndexEncryption.go
│   │    ├── keygen.go
│   │    └── keyword-database.go
│   └── main-initiate.go #Set up
├── pkg        
└── bin          
 
```

### Initiate the system 
//Client side
* Generates encryption key (AES-128)
* Encrypts the file collection.
* Encrypts inverted index (In form of: "keyword", "bitmap"). Bitmap should be n-bit where n is number of document in file collection
* Send it to network (aka. nearest node) (DHT //**TODO!!** - not finished)


//Server side
* Distributed the index due to hash range (DHT //**TODO!!** - not finished)
* repeat until index is empty

### Search the file address 
(//**TODO!!** - not finished)

//Client side
* Generates token using the keyword and secret key
* send <token, gamma> to the network

//Server side
* If the token is matched the hash range, extracts it and sends it to the client.
* Otherwise, sends the token to next hop.

### File retrieval

(//**TODO!!** - not finished)

//Client side
* Decrypts the index to get the Bitmap
* Extracts the fileID from Bitmap
* Send fileID to the nearest node (aka. want manager)

//Server side
* Want manager generates wantlist and sends to the active neighbor node.
* If the node contains the target file, the storage node decides the resource allocation. Then, sends the file collection to client according to bitswap strategies.
* Otherwise, sends the wantlist to neighbor node.
* Repeat them until received the "cancle" message.


## User Guide 

Helper commands:

```Bash

Commands
  help         help guide to specify each available commands
  quit         exit the console
  keygen     	 Key generation- automatically stored on local machine

Initiate commands:
  insertI <keyword> <bitmap>  construct the inverted index
  fileEn <inputfile>      		encrypt the target file
  fileDe <inputfile>      		decrypt the target file

Search commands:
  searchKey <keyword> 	Search the keyword

Update commands:
  updatekey <keyword> <bitmap>	Update the bitmap of corresponding keyword.

 
```


