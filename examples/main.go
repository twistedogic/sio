package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/twistedogic/sio"
)

type defaultSource struct{}

func (defaultSource) Name() string { return "System" }

func (d defaultSource) Response(ctx context.Context, input string) (string, error) {
	if strings.ToLower(input) == "bye" {
		return "bye", nil
	}
	n := rand.Intn(5)
	time.Sleep(time.Duration(n) * time.Second)
	return fmt.Sprintf("received - %q", input), nil
}

func main() {
	if err := sio.Start(context.Background(), "system", defaultSource{}); err != nil {
		log.Fatal(err)
	}
}
