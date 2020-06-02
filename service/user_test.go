package service

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	s := "Hello,我在"
	bytes := []byte(s)
	fmt.Println("bytes", bytes)
	s1 := string(bytes)
	fmt.Printf("s1:%q", s1)
	if s != s1 {
		t.Errorf("s1:%q与s不相等", s1)
	}
}
