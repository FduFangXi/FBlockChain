package main

import (
	"time"
)

/*
** Block 由区块头和交易两部分构成
** Timestamp，PrevBlockHash，Hash属于区块头（block header）
** Timestamp: 当前时间戳，也就是区块创建的时间
** PrevBlockHash: 前一个块的哈希
** Hash: 当前块的哈希
** Data: 区块实际存储的信息，在比特币中就是交易
** Nonce: 在工作量证明进行验证时需要用到
 */
type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
	Nonce         int
}

/*
** NewBlock 用于生成新块
** params Data: 新块生成所存储的实际数据
** params PrevBlockHash: 其所链接的前一个快的地址哈希值
** 当前块的哈希会基于Data、PrevBlockHash生成，也就是 Data & PrevBlockHash => Hash
 */
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
		Nonce:         0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	// block.SetHash()
	return block
}

/* 加入ProofOfWork后移除，哈希改为工作量证明进行生成
** SetHash 生成当前块的哈希
** Hash <= sha256(PrevBlockHash + Data + Timestamp)
 */
// func (block *Block) SetHash() {
// 	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
// 	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)

// 	block.Hash = hash[:]
// }

/*
** NewGenesisBlock 生成创世区块
 */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
