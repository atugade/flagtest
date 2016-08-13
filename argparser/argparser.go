package argparser

import "flag"
import "fmt"

type Launch struct {
	Command  *flag.FlagSet
	Manifest *string
	Keyname  *string
	Numinst  *int
}

type Stack struct {
	Command *flag.FlagSet
	Up      *bool
	Mod     *bool
	Down    *bool
	Account *string
	Layer   *string
	Env     *string
}

func LaunchSubcommand() Launch {
	launchCommand := flag.NewFlagSet("launch", flag.ExitOnError)

	manifestFlag := launchCommand.String("manifest", "", "Manifest file")
	keynameFlag := launchCommand.String("keyname", "", "Keypair name")
	numinstFlag := launchCommand.Int("numinst", 0, "Number of instances")

	launchInfo := Launch{Command: launchCommand, Manifest: manifestFlag,
		Keyname: keynameFlag, Numinst: numinstFlag}

	return launchInfo
}

func LaunchProcessArgs(launchInfo Launch) {
	fmt.Println("manifest: ", *launchInfo.Manifest)
}

func StackSubcommand() Stack {
	stackCommand := flag.NewFlagSet("stack", flag.ExitOnError)

	upFlag := stackCommand.Bool("up", false, "Stand up a stack")
	modFlag := stackCommand.Bool("mod", false, "Modify a stack")
	downFlag := stackCommand.Bool("down", false, "Destroy a stack")
	accountFlag := stackCommand.String("account", "", "Account short name")
	layerFlag := stackCommand.String("layer", "", "Stack layer")
	envFlag := stackCommand.String("env", "", "Target environment")

	stackInfo := Stack{Command: stackCommand, Up: upFlag, Mod: modFlag,
		Down: downFlag, Account: accountFlag, Layer: layerFlag,
		Env: envFlag}

	return stackInfo
}

func StackProcessArgs(stackInfo Stack) {
	activated := 0

	for _, flag := range []bool{*stackInfo.Up, *stackInfo.Mod,
		*stackInfo.Down} {
		if flag {
			activated += 1
		}
	}

	if activated > 1 {
		fmt.Println("-up, -mod and -down are mutually exclusive")
		stackInfo.Command.PrintDefaults()
	}
}

func FlagTestUsage() {
	fmt.Println("usage: tesseract <command> [<args>]")
	fmt.Println("The most commonly used tesseract commands are: ")
	fmt.Println(" launch  Build and provision an instance")
	fmt.Println(" stack   Standup, modify or bring down a stack")
}
