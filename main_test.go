package target

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRandomSleep(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	d := time.Duration(1+rand.Intn(3)) * time.Second
	fmt.Println("Sleeping for", d)
	time.Sleep(d)
}

func TestMult(t *testing.T) {
	require.Equal(t, mult(2, 3), 6)
}

func TestAdd(t *testing.T) {
	require.Equal(t, Add(3, 4, 5), 12)
}
