package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	err := errors.New("oops, I did it again!!")
	wrapped := Wrap(err, "ERR")
	require.Equal(t, "ERR: oops, I did it again!!", wrapped.Error())

	me, ok := wrapped.(*MyError)
	require.True(t, ok)
	require.Equal(t, err, me.origErr)
	require.Equal(t, "ERR: oops, I did it again!!", me.message)

	wrapped = testWrapThirdCaller()
	me = wrapped.(*MyError)
	expectedNotes := []string{
		"testWrapFirstCaller",
		"testWrapSecondCaller",
		"testWrapThirdCaller",
	}
	require.Equal(t, len(expectedNotes), len(me.notes))
	for i := range me.notes {
		require.Equal(t, expectedNotes[i], me.notes[i])
	}

	expectedNotesInString := "testWrapFirstCaller -> testWrapSecondCaller -> testWrapThirdCaller"
	require.Equal(t, expectedNotesInString, me.NotesInString())

	expectedMessage := "testWrapFirstCaller: oops, something error!!"
	require.Equal(t, expectedMessage, me.message)
}

func testWrapThirdCaller() error {
	return Wrap(testWrapSecondCaller(), "testWrapThirdCaller")
}

func testWrapSecondCaller() error {
	return Wrap(testWrapFirstCaller(), "testWrapSecondCaller")
}

func testWrapFirstCaller() error {
	return Wrap(testWrapOrigin(), "testWrapFirstCaller")
}

func testWrapOrigin() error {
	return fmt.Errorf("oops, something error!!")
}
