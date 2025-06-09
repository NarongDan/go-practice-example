package exception

type ItemEditing struct{}

func (e *ItemEditing) Error() string {
	return "editing item failed"
}
