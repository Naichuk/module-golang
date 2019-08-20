package main

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(database *sql.DB, data string) *Block {
	rows, err := database.Query("SELECT * FROM Block WHERE id = (SELECT MAX(id)  FROM Block)")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var id int
	p := Block{}
	rows.Next()
	g := rows.Scan(&id, &p.Timestamp, &p.Data, &p.Hash, &p.PrevBlockHash)
	if g != nil {
		fmt.Println(g)
	}
	newBlock := NewBlock(data, p.Hash)
	bc.blocks = append(bc.blocks, newBlock)
	return newBlock
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func db_Block(block *Block, database *sql.DB) *sql.Stmt {

	statement, err := database.Prepare("INSERT INTO Block (timestamp, data, hash, prevHash) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(block.Timestamp, block.Data, block.Hash,
		block.PrevBlockHash)
	return statement
}

func isGenesis(database *sql.DB) bool {
	rows, _ := database.Query("SELECT data FROM Block WHERE id = 1")
	var data string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data)
		return false
	}
	return true
}

func main() {
	bc := NewBlockchain()
	db, err := sql.Open("sqlite3", "blockchain_db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("CREATE TABLE IF NOT EXISTS Block (id INTEGER PRIMARY KEY, timestamp INTEGER, data BLOB, hash TEXT, prevHash TEXT)")
	if err != nil {
		panic(err)
	}
	if isGenesis(db) {
		bc0 := NewGenesisBlock()
		db_Block(bc0, db)
	}

	if os.Args[1] == "add" {
		bc2 := bc.AddBlock(db, os.Args[2])
		db_Block(bc2, db)
		fmt.Println("Transaction was added to the blockchain")
	}
	if os.Args[1] == "list" {
		rows, err := db.Query("select * from Block")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		products := []Block{}
		var id int
		for rows.Next() {
			p := Block{}
			err := rows.Scan(&id, &p.Timestamp, &p.Data, &p.Hash, &p.PrevBlockHash)
			if err != nil {
				fmt.Println(err)
				continue
			}
			products = append(products, p)
		}
		for _, p := range products {
			fmt.Printf("Prev Hash: %x\n", p.PrevBlockHash)
			fmt.Printf("Data: %s\n", p.Data)
			fmt.Printf("Hash: %x\n", p.Hash)
			fmt.Println()
		}
	}

}
