package slicesext

import (
	"fmt"
	"testing"

	"github.com/fluentassert/verify"
)

func TestMap(t *testing.T) {
	in := []int{1, 2, 3}
	out := Map(in, func(e int) string { return fmt.Sprint(e) })
	verify.Slice(out).Equivalent([]string{"1", "2", "3"})
}

func TestReduce(t *testing.T) {
	in := []int{1, 2, 3}
	out := Reduce(in, func(acc int, e int) int { return acc + e }, 0)
	verify.Number(out).Equal(6)
}
