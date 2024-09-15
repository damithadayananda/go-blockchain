package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-blockchain/app"
	"go-blockchain/config"
	"go-blockchain/core/transaction"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        int64
	Data         interface{}
	MerkleRoot   string
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Nonce        int32
}

// Mine once the mining is completed we are returning
// and sending signal to done channel
// both the terminal operations are important if we just return only
// next mining cycle will start once we receive a new transaction only
func (b *Block) Mine(stop <-chan bool, done chan<- bool) (interrupted bool) {
	for {
		select {
		case <-stop:
			app.Logger.Info.Log("Mining is interrupted")
			// stop mining triggering new one
			done <- true
			return true
		default:
			for {
				b.Nonce++
				sha := sha256.New()
				valueString := fmt.Sprintf("merkleRoot: %v, previousHash: %v, nonce: %v ", b.MerkleRoot, b.PreviousHash, b.Nonce)
				sha.Write([]byte(valueString))
				bytString := hex.EncodeToString(sha.Sum(nil))
				subString := bytString[:3]
				numOfZeroes := strings.Count(subString, "0")
				if numOfZeroes == int(config.AppConfig.Complexity) {
					b.Hash = bytString
					// ending mining operation
					app.Logger.Info.Log("Mining is complete")
					done <- true
					return false
				}
			}
		}

	}
}
func (b *Block) CalculateMerkleRoot() {
	switch v := b.Data.(type) {
	case []transaction.Transaction:
		b.MerkleRoot = calculateMerkleRootForSliceOfTransaction(v)
	case string:
		b.MerkleRoot = calculateMerkleRootForString(v)
	}
}

func calculateMerkleRootForString(s string) string {
	sha := sha256.New()
	byteArray, _ := json.Marshal(s)
	sha.Write(byteArray)
	return hex.EncodeToString(sha.Sum(nil))
}

func calculateMerkleRootForSliceOfTransaction(txn []transaction.Transaction) string {
	hashes := calculateHashOfTransactions(txn)
	return MerkleRootFromTransactionHashes(hashes)
}
func calculateHashOfTransactions(transactions []transaction.Transaction) []string {
	var hashes []string
	for _, transaction := range transactions {
		str := transaction.Id + transaction.Receiver + transaction.Sender +
			strconv.FormatFloat(transaction.Amount, 'f', -1, 64) +
			strconv.FormatFloat(transaction.Fee, 'f', -1, 64)
		sha := sha256.New()
		byteArray, _ := json.Marshal(str)
		sha.Write(byteArray)
		hashes = append(hashes, hex.EncodeToString(sha.Sum(nil)))
	}
	return hashes
}

func MerkleRootFromTransactionHashes(hashes []string) string {
	if len(hashes) == 0 {
		return ""
	}
	if len(hashes) == 1 {
		return hashes[0]
	}
	// Process the list by combining each pair of hashes.
	var newLevel []string
	for i := 0; i < len(hashes); i += 2 {
		// If there's an odd number of hashes, duplicate the last one.
		if i+1 == len(hashes) {
			newLevel = append(newLevel, CalculateHash(hashes[i], hashes[i]))
		} else {
			newLevel = append(newLevel, CalculateHash(hashes[i], hashes[i+1]))
		}
	}
	// Recursively calculate the Merkle root until we reach the top level.
	return MerkleRootFromTransactionHashes(newLevel)
}

func CalculateHash(left, right string) string {
	hash := sha256.Sum256([]byte(left + right))
	return hex.EncodeToString(hash[:])
}
