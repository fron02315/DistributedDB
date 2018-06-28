Noted that this text is informal way to describe my project. 
The white paper will be released soon.

# DistributedDB
This project is regarding from my project course. The project proposed to provide search operation on distributed file storage.
Originally, distributed database is not efficient due to communication cost and routing complexity. 

## Getting Started
The project is divided into three part.
* Initiate the system
* Search the file address
* File retrieval -> FOCUS ON THIS :)

### Initiate the system 
//Client side
* Generates encryption key (AES-128 yeah :v)
* Encrypts the file collection.
* Encrypts inverted index (In form of: "keyword", "bitmap"). Bitmap should be n-bit where n is number of document in file collection
* Send it to network (aka. nearest node) (DHT //**TODO!!** - not finished :-( yeah... \m/)


//Server side
* Distributed the index due to hash range (DHT //**TODO!!** - not finished :-( yeah... \m/)
* repeat until index is empty

### Search the file address
//Client side
* Generates token using the keyword and secret key
* send <token, gamma> to the network

//Server side
* If the token is matched the hash range, extracts it and sends it to the client.
* Otherwise, sends the token to next hop.




```
It doesn’t matter who you are or what you do. Just remember this one thing. 
Someone loves you more than you love yourself. I will cheer you on as well.
I love you.
#JinKi :-)💕🌞
```
