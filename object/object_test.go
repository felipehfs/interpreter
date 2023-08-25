package object

import "testing"

func TestStringHashkey(t *testing.T) {
	hello1 := &String{Value: "Hello world"}
	hello2 := &String{Value: "Hello world"}
	diff1 := &String{Value: "My name is Johnny"}
	diff2 := &String{Value: "My name is Johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
