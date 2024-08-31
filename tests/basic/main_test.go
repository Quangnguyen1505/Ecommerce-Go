package basic

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestTotal(t *testing.T) {
	var (
		intput = 1
		ouput  = 2
	)
	// actual := totalNumber(intput)
	assert.Equal(t, totalNumber(intput), ouput)
	// if actual != ouput {
	// 	t.Errorf("Expected %d but got %d", ouput, actual)
	// }

}

// func TestRequire(t *testing.T) {
// 	require.Equal(t, totalNumber(1), 3)
// 	fmt.Println("TestRequire aa")
// }

// func TestAssert(t *testing.T) {
// 	assert.Equal(t, totalNumber(1), 3)
// 	fmt.Println("TestAssert bb")
// }
