package transaction

import (
	"strconv"
)

type TransactionList []Transaction

func (tl TransactionList) ToString() string {
	var fullStr string
	for _, tx := range tl {
		str := tx.Id + tx.Sender + tx.Receiver + strconv.FormatFloat(tx.Amount, 'f', 2, 64) + strconv.FormatFloat(tx.Fee, 'f', 2, 64)
		fullStr = fullStr + str
	}
	return fullStr
}
