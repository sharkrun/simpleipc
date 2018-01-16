package test

import (
	"testing"
)


func TestTest1(t *testing.T) {
	r := Test1(9)

	if r[0] != 9 {
		t.Errorf("Test1(9) failed. Expected [9]. Got:%v", r)
	}

	}
