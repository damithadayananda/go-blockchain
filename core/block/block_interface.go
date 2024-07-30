package block

type IBlock interface {
	mine() []byte
}
