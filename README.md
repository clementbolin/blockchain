# Blockchain

### Description

Verify that Golang is installed on your machine.
The objective of this small project was to create a simple BlockChain in order to better understand these concepts.
in order to realize this project I read it's different topics [BlockChain Explained](https://www.investopedia.com/terms/b/blockchain.asp), [What is Blockchain Technology](https://blockgeeks.com/guides/what-is-blockchain-technology/) 

### How to use

```bash
git clone https://github.com/ClementBolin/blockchain.git && cd blockchain
go build
./blockchain
```

##### How to use BlockChain

- Send ```GET``` Request at ```http://localhost:8080/``` for watch BlockChain Status
- Send ```POST``` Request at ```http://localhost:8080/``` with body json like this
    ```json
    "bpm": 10
    ```
    you can watch after in your terminal status of BlockChain example :
    ```bash
    (main.Block) {
        index: (int) 1,
        timeStamp: (string) (len=51) "2020-11-14 18:07:25.747534 +0100 CET m=+2.759705884",
        bpm: (int) 0,
        hash: (string) (len=64) "6dea7de6593d74c217bc017a62f8bfad4846138f6db3aa73cc6e66c0b90e73a0",
        prevHash: (string) (len=64) "96a296d224f285c67bee93c30f8a309157f0daa35dc5b87e410b78630a09cfc7"
        },
        (main.Block) {
        index: (int) 2,
        timeStamp: (string) (len=52) "2020-11-14 18:07:50.521991 +0100 CET m=+27.533483335",
        bpm: (int) 500,
        hash: (string) (len=64) "0a322a6877c0de5a7432b000fd7e787bdcd43738efad5278b00809f64b82a308",
        prevHash: (string) (len=64) "6dea7de6593d74c217bc017a62f8bfad4846138f6db3aa73cc6e66c0b90e73a0"
        },
        (main.Block) {
        index: (int) 3,
        timeStamp: (string) (len=52) "2020-11-14 18:07:58.238916 +0100 CET m=+35.250196832",
        bpm: (int) 500,
        hash: (string) (len=64) "f5d7ac43eb765e7401a4c7ef81062f6bc8c2468e490536f5998ca78ac90b6403",
        prevHash: (string) (len=64) "0a322a6877c0de5a7432b000fd7e787bdcd43738efad5278b00809f64b82a308"
        },
        (main.Block) {
        index: (int) 4,
        timeStamp: (string) (len=52) "2020-11-14 18:08:05.565851 +0100 CET m=+42.576930719",
        bpm: (int) 700,
        hash: (string) (len=64) "846f7d74abda0fa36695b4127ea00dfdb149136a5c21ce34e4d5ed505416124c",
        prevHash: (string) (len=64) "f5d7ac43eb765e7401a4c7ef81062f6bc8c2468e490536f5998ca78ac90b6403"
        }
    }
    ```

