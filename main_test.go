package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomSleep(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	d := time.Duration(10+rand.Intn(20)) * time.Second
	fmt.Println("Sleeping for", d)
	time.Sleep(d)
}

func TestAdd(t *testing.T) {
	add(1, 2)
}
