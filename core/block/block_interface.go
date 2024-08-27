package block

type IBlock interface {
	Mine(stop <-chan bool, done chan<- bool) (interrupted bool)
}
