/*package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block { //creating a new block
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

type Blockchain struct { //blockchain is created
	blocks []*Block //array of blocks
}

func (bc *Blockchain) AddBlock(data string) { //adding blocks into blockchain
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}



type Item struct {

	Id int
	Data string
	Hash []byte
	prevHash []byte
}



func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}


func CreateTable(db *sql.DB) {
	sql_table := `
	CREATE TABLE IF NOT EXISTS block(
		Id INTEGER PRIMARY KEY,
		timestamp INTEGER,
		data BLOB,
		hash TEXT
		prevHash TEXT
	);
	`
	_, err := db.Exec(sql_table)
	if err != nil { panic(err) }
}



func StoreItem(db *sql.DB, items []Item) {
	sql_additem := `
	INSERT OR REPLACE INTO items(
		Id,
		timestamp
		data,
		hash,
		prevHash
	) values(?, CURRENT_TIMESTAMP, ?, ?, ?)
	`

	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(item.Id, item.Name, item.Phone)
		if err2 != nil { panic(err2) }
	}
}





func main() {
		db, err:= sql.Open("sqlite3", "./blockchain.db")
		if err != nil {
		    panic(err)
		}
		defer db.Close()
		result, err := db.Exec("insert into products (product, price )values (os.Args[2], os.Args[3])",
		if err != nil{
		    panic(err)
		    }
		}
		bc := NewBlockchain()
		switch os.Args[1] {
		case "add":
			bc.AddBlock(os.Args[2])
			fmt.Println("Transaction was added to the blockchain")
		case "list":
		for _, block := range bc.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
	}
	}*/
