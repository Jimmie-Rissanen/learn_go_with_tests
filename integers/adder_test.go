package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	actual := Add(2, 2)
	expected := 4

	if actual != expected {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
