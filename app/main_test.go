package main

import "testing"

func TestSum(t *testing.T) {
	if add(2, 5) != 7 {
		t.Fail()
	}
	if add(2, 100) != 102 {
		t.Fail()
	}
	if add(222, 100) != 322 {
		t.Fail()
	}
}

func TestProduct(t *testing.T) {
	if multiply(2, 5) != 10 {
		t.Fail()
	}
	if multiply(2, 100) != 200 {
		t.Fail()
	}
	if multiply(111, 3) != 333 {
		t.Fail()
	}
}
