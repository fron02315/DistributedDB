# DistributedDB
This project is regarding from my project course. The project proposed to provide search operation on distributed file storage.
Originally, distributed database is not efficient due to communication cost and routing complexity. 

## Getting Started
The project is divided into three part.
* Initiate the system
* Search the file address
* File retrieval -> FOCUS ON THIS :)

### Initiate the system 
* Generates encryption key (AES-128 yeah :v)
* Encrypts the file collection.
* Encrypts inverted index (In form of: "keyword", "bitmap"). Bitmap should be n-bit where n is number of document in file collection
* Send it to network (aka. nearest node)
//Progress HERE: store them on the database

### Search the file system
