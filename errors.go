package errors

import (
	"fmt"
	"strings"
)

type MyError struct {
	origErr error
	notes   []string
	message string
}

func (me *MyError) Error() string {
	return me.message
}

func (me *MyError) NotesInString() string {
	var result string
	for _, val := range me.notes {
		result += val + " -> "
	}
	return strings.TrimSuffix(result, " -> ")
}

func newMyError(err error) *MyError {
	return &MyError{
		origErr: err,
		notes:   make([]string, 0),
	}
}

func Wrap(err error, note string) error {
	me, ok := err.(*MyError)
	if !ok {
		me = newMyError(err)
		me.message = fmt.Sprintf("%s: %v", note, err)
	}
	me.notes = append(me.notes, note)
	return me
}
