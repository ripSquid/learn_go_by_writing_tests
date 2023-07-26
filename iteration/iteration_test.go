package iteration

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	hello = os.Getenv("HELLO")
	world = os.Getenv("WORLD")
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 4)
	expected := "aaaa"

	if got != expected {
		t.Errorf("got %q but expected %q", got, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 512)
	}
}

//go:noinline
func nop() {}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := strings.Compare(hello, world)
		switch {
		case x == 0:
			nop()
		case x > 0:
			nop()
		case x < 0:
			nop()
		}
	}
}

func BenchmarkEqOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch {
		case hello == world:
			nop()
		case hello < world:
			nop()
		case hello > world:
			nop()
		}
	}
}

func ExampleRepeat() {
	tmp := Repeat("Hello, ", 2)
	fmt.Println(tmp)
	// Output: Hello, Hello,
}
