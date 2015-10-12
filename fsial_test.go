package fsial
import (
	"testing"
)

func TestGen(t *testing.T) {
	val, err := Gen(999999999999999999, 15)
	t.Logf("%v, %v\n", val, err)
	val, err = Gen(18620897889, 15)
	t.Logf("%v, %v\n", val, err)
}