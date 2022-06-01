package rangelist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tmpString string

func init() {
	print = func(a ...any) (n int, err error) {
		tmpString = fmt.Sprint(a)
		b := []byte(tmpString)
		tmpString = string(b[1 : len(b)-1])
		return 0, nil
	}
}

func TestRangelist(t *testing.T) {

	rl := New()

	err := rl.Add([2]int{5, 1})
	assert.EqualError(t, err, "invalid range")

	err = rl.Add([2]int{1, 5})
	assert.Nilf(t, err, "Add should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 5)", tmpString, "The two should be equal")

	err = rl.Add([2]int{10, 20})
	assert.Nilf(t, err, "Add should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 5) [10, 20)", tmpString, "The two should be equal")

	err = rl.Add([2]int{20, 20})
	assert.Nilf(t, err, "Add should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 5) [10, 20)", tmpString, "The two should be equal")

	err = rl.Add([2]int{20, 21})
	assert.Nilf(t, err, "Add should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 5) [10, 21)", tmpString, "The two should be equal")

	err = rl.Add([2]int{2, 4})
	assert.Nilf(t, err, "Add should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 5) [10, 21)", tmpString, "The two should be equal")

	err = rl.Add([2]int{3, 8})
	assert.Nilf(t, err, "Add should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 8) [10, 21)", tmpString, "The two should be equal")

	err = rl.Remove([2]int{10, 10})
	assert.Nilf(t, err, "Remove should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 8) [10, 21)", tmpString, "The two should be equal")

	err = rl.Remove([2]int{10, 11})
	assert.Nilf(t, err, "Remove should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 8) [11, 21)", tmpString, "The two should be equal")

	err = rl.Remove([2]int{15, 17})
	assert.Nilf(t, err, "Remove should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 8) [11, 15) [17, 21)", tmpString, "The two should be equal")

	err = rl.Remove([2]int{3, 19})
	assert.Nilf(t, err, "Remove should succeed")
	err = rl.Print()
	assert.Nilf(t, err, "Print should succeed")
	assert.Equal(t, "[1, 3) [19, 21)", tmpString, "The two should be equal")

}
