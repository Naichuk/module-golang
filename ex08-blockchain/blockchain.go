package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	//	"fmt"
	"io"
	//	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const filename = "blockchain.json"

type Block struct {
	Timestamp     int64  `json:"timestamp"`
	Data          []byte `json:"data"`
	PrevBlockHash []byte `json:"prevblockhash"`
	Hash          []byte `json:"hash "`
}

type Blockchain struct {
	blocks []*Block `json:"blocks"`
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

func (bc *Blockchain) AddBlock(data string) {
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

// Work with files
func CreateFile() {
	block1 := NewBlockchain()

	jsonFile, err := os.Create(filename)
	must(err)

	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	err = encoder.Encode(&block1)
	must(err)
}

/*func Read() {
	content, err := ioutil.Readfile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
}
*/
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func Add() {

	CreateFile()
}
func main() {
	switch os.Args[1] {
	case "add":
		Add()
		//	case "list":
	}
}
