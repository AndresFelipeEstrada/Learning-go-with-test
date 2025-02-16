package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	if Add(2, 2) != 4 {
		t.Errorf("expected '%d' but got '%d'", 4, Add(2, 2))
	}
}

func ExampleAdd() {
	fmt.Println(Add(1, 5))
	// Output: 6
}
