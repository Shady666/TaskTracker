package errorList

import "errors"

var (
	ErrTaskNotExist   = errors.New("task not exist")
	ErrTasksNotExist  = errors.New("tasks not exist")
	ErrTaskNotUpdated = errors.New("task not updated")
)
