package core

type BlockChainInterface interface {
	AddBlock(block *Block) error
	ValidateBlock(block *Block) error
}
