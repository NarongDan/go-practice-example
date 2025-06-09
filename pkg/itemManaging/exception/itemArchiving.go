package exception

type ItemArchiving struct{}

func (e *ItemArchiving) Error() string {
	return "archiving item failed"
}
