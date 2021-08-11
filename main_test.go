package target

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomSleep(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	d := time.Duration(10+rand.Intn(10)) * time.Second
	fmt.Println("Sleeping for", d)
	time.Sleep(d)
}

func TestMult(t *testing.T) {
	_ = mult(3, 4)
}
