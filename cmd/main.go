package main

import (
	"os"

	"github.com/me4hit/distributed-go-cache/node"
)

func main() {
	node := node.NewNode("node1", 100)
	os.Setenv("PORT", "3000")
	node.Start()
}
