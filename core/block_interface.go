package core

type IBlock interface {
	mine() []byte
}
