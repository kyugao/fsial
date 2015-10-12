package fsial
import (
	"errors"
	"fmt"
)

const (
	r = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"
	l = uint64(len(r))
	MaxLen = 20
)

func Gen(num uint64, length int) (result string, err error){
	rb := r[:]

	if length > MaxLen {
		return "", errors.New(fmt.Sprintf("Require length %d exceed %d", length, MaxLen))
	}

	re := make([]byte, 0, MaxLen)

	index := 0
	for num > l {
		dig := num % l
		re = append(re, rb[dig])
		num /= l

		index ++
	}

	result = string(re[:len(re)])
	return
}