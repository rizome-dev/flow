package main

import (
	"github.com/rizome-dev/flow"
)

func main() {
	f := flow.NewFlow()
	f.AddDeepseekAgent("master")
}
