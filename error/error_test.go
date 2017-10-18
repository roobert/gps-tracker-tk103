package error

import (
	"errors"
	"testing"
)

func TestCheckErrWithNil(t *testing.T) {
	CheckErr(nil)
}

func TestCheckErrWithError(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("checkErr did not panic when passed an error")
			}
		}()

		CheckErr(errors.New("test error"))
	}()
}
