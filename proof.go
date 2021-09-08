package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

/*
** targetBits 定义挖矿的难度值
** 在比特币中，当一个块被挖出来以后，”target bits“代表区块头里存储的难度，也就是开头有多少个0
** 24指的是算出来的哈希前24位必须是0，如果是用16进制表示，前6位就必须是0
 */
const targetBits = 24

// Nonce 循环的最大次数，避免溢出
const maxNonce = math.MaxInt64

/*
** 每个块必须证明的工作量结构体
** 指向Block块的指针
** target是证明的目标，最终找的哈希要小于目标，转换成大整数进行比较
 */
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

/*
** NewProofOfWork 生成新的证明
** target <= 1左移256位 - targetBits位
 */
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

/*
** prepareData 哈希之前的准备数据
** 需要用到PrevBlockHash, Data, Timestamp, targetBits, nonce
** 其中nonce就是HashCash的计数器
 */
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

/*
** Run 开始进行哈希
** 首先变量初始化，HashInt 是 hash 的整形表示；nonce是计数器
** 然后开始循环：maxNonce = math.MaxInt64，为了避免 nonce 可能出现的溢出。
** 循环中： 1、准备数据
** 2、用 SHA-256 对数据进行哈希
** 3、将哈希转换成一个大整数
** 4、将这个大整数与目标进行比较
 */
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the Block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isVlid := hashInt.Cmp(pow.target) == -1

	return isVlid
}
