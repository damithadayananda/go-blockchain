package minor

import (
	"go-blockchain/app"
	"go-blockchain/core/block"
	"go-blockchain/core/blockchain"
	"go-blockchain/core/mempool"
	"go-blockchain/core/transaction"
	"sync"
	"time"
)

type minorImp struct {
	isInMining          bool
	memPool             mempool.MemPool
	transactionOnMining []transaction.Transaction
	mineLock            sync.Mutex
	stopChan            chan bool
	doneChan            chan bool
	chain               blockchain.BlockChainInterface
}

func NewMinor(memPool *mempool.MemPool) {
	minor := minorImp{
		memPool:  *memPool,
		stopChan: make(chan bool),
		doneChan: make(chan bool),
		chain:    blockchain.Chain,
	}
	// start mining
	go minor.Start()
	// stop mining if transactions are removed from mem pool
	go minor.stop()
}

func (minor *minorImp) Start() {
	for {
		select {
		case <-minor.memPool.Receiver:
			// start mining once transaction added to mem pool
			app.Logger.Info.Log("Received signal on memPool.Receiver, starting mining...")
			go minor.mine()
		case <-minor.doneChan:
			// start new mining once previous one is done
			go minor.mine()
		}
	}
}

func (minor *minorImp) mine() {
	// is to avoid two parallel mining operations
	// lock is being placed
	// -----why both lock and flag-----
	// if only lock, for each transaction there will be queue to acquire a lock
	if minor.isInMining {
		return
	}
	minor.mineLock.Lock()
	defer minor.mineLock.Unlock()
	minor.isInMining = true
	txn, err := minor.memPool.Get()
	if err != nil || txn == nil {
		minor.isInMining = false
		return
	}
	//Adding self transaction, summation of fee
	var fee = float64(0)
	for _, tx := range txn {
		fee = fee + tx.Fee
	}
	txn = append(txn, transaction.NewTransaction(transaction.Transaction{
		Amount:   fee,
		Receiver: app.App.Address,
	}))

	pb, _ := minor.chain.GetLastBlock()
	b := block.Block{
		Data:         txn,
		PreviousHash: pb.Hash,
		Timestamp:    time.Now(),
	}
	b.CalculateMerkleRoot()
	interrupted := b.Mine(minor.stopChan, minor.doneChan)
	minor.chain.AddBlock([]block.Block{b})
	// since mine function is thread safe
	// calling marking function here should be fine
	if !interrupted {
		minor.memPool.Mark(txn)
	}
	minor.isInMining = false
}

func (minor *minorImp) stop() {
	for {
		select {
		case <-minor.memPool.Remover:
			// stop on going mining
			minor.doneChan <- true
		}
	}
}

func (minor *minorImp) Stop() {

}
