package main

import (
	"fmt"
	"os"

	"github.com/atugade/flagtest/argparser"
	"github.com/atugade/flagtest/launch"
	"github.com/atugade/flagtest/stack"
)

func main() {
	//argparser.GetParsedArgs()

	if len(os.Args) == 1 {
		argparser.FlagTestUsage()
		return
	}

	switch os.Args[1] {
	case "launch":
		launch_args := argparser.LaunchSubcommand()
		launch_args.Command.Parse(os.Args[2:])

		if launch_args.Command.Parsed() {
			argparser.LaunchProcessArgs(launch_args)
			launch.LaunchCommand(launch_args)
		}
	case "stack":
		stack_args := argparser.StackSubcommand()
		stack_args.Command.Parse(os.Args[2:])

		if stack_args.Command.Parsed() {
			argparser.StackProcessArgs(stack_args)
			stack.StackCommand(stack_args)
		}
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
