package test

import (
	"fmt"
	"parcel-management/util"
	"testing"
	"time"
)

func TestRandomEmail(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(util.RandomEmail())
	}
}

func TestRandomUnitNo(t *testing.T) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("randomUnit %v\n", util.RandomUnitNo())
	}
}
