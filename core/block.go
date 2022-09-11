package core

import "time"

type block struct {
	Data         any
	Hash         []byte
	PreviousHash []byte
	Timestamp    time.Time
	Nonce        int32
}
