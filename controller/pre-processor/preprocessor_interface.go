package pre_processor

import "go-blockchain/core/block"

type IPreProcessor interface {
	ProcessBlock(block *block.Block, callerAddress string) (error, []block.Block)
}
