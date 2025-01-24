package flow_test

import (
	"testing"

	"github.com/rizome-dev/flow"
)

func Test(t *testing.T) {
	f := flow.NewFlow()
	f.AddDeepseekAgent("master")
}
