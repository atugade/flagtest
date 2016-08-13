package stack

import (
	"fmt"

	//github.com/atugade/flagtest/providers/cloudformation
	"github.com/atugade/flagtest/argparser"
)

func StackCommand(stack argparser.Stack) {
	fmt.Println("in StackCommand")
}
