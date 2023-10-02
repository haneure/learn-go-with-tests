package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestContains(t *testing.T) {
	haysack := Contains("Sword Art Online", "Online")
	needle := "Online"

	if !haysack {
		t.Errorf("The string does not containing %q", needle)
	}
}

func TestSplit(t *testing.T) {
	split := Split("Getsuyoubi,Kayoubi,Suiyoubi,Mokuyoubi,Kinyoubi,Doyoubi,Nichiyoubi", ",")

	expected := make([]string, len(split))
	expected[0] = "Getsuyoubi"
	expected[1] = "Kayoubi"
	expected[2] = "Suiyoubi"
	expected[3] = "Mokuyoubi"
	expected[4] = "Kinyoubi"
	expected[5] = "Doyoubi"
	expected[6] = "Nichiyoubi"

	if !isSlicesEqual(split, expected) {
		t.Errorf("String %q splitted is not the same as %q", &split, expected)
	}
}

func TestIsSlicesEqual(t *testing.T) {
	slice1 := make([]string, 2)
	slice1[0] = "1"
	slice1[1] = "2"

	slice2 := make([]string, 2)
	slice2[0] = "1"
	slice2[1] = "2"

	if !isSlicesEqual(slice1, slice2) {
		t.Errorf("String slices %q is not the same as %q", &slice1, slice2)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
