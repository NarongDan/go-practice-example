package exception

type PLayerCoinShowing struct{}

func (e *PLayerCoinShowing) Error() string {
	return "player coin showing failed"
}
