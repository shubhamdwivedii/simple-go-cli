package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	// go get github.com/spf13/cobra
	"github.com/spf13/cobra"
)

var (
	persistRootFlag bool
	localRootFlag   bool
	times           int
	rootCmd         = &cobra.Command{
		Use:   "example",                  // Basically name of the command
		Short: "An Example Cobra Program", // Short description
		Long: `This is a simple example of a cobra program.
It will have several subcommands and flags.`, // help message
		// using --help now will print the long description

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello From The Root Command")
		}, // run without any flag (just go run main.go) will execute this function.
		// using --help now will have addition Usage: example [flags] section
	}

	echoCmd = &cobra.Command{
		Use:   "echo [strings to echo]",
		Short: "prints given strings to stdout",
		Args:  cobra.MinimumNArgs(1), // to ensure minumum number of args
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
		// try> go run main.go echo shubham dwivedi
		// run> go run main.go echo -h (to see help for echo command only)
	}

	timesCmd = &cobra.Command{
		Use:   "times [strings to echo]",
		Short: "prints given strings to stdout multiple times",
		Args:  cobra.MinimumNArgs(1),
		// RunE is just like Run except it can return an "error"
		RunE: func(cmd *cobra.Command, args []string) error {
			if times == 0 {
				return errors.New("Times Cannot Be Zero")
			}
			for i := 0; i < times; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
			return nil
		},
	} // run > go run main.go help echo // OR go run main.go echo -help
	// this will show times under commands.
	// run > go run main.go echo times -help
	// run > go run main.go echo times -t 12 shubham dwivedi
	// run > go run main.go echo times -t 0 shubham // to force error

	// You can implement custom validator using RunE
)

/* Cobra has concept of Two Different Flags **************
1. Local flags (flags) - Only available withing the context of command its defined on.
Eg. If defined in root command its only available when root command is processed.
It won't be tossed down to any subcommands. (Not available to subcommands)

2. Persistent flags (pflags) - Stay available and relevant starting on the command they are defined on.
And any other command that is sub-command of that command.
*/

// Look up init in Go
func init() { // BoolVarP() takes: p *bool, name, shorthand, value, usage
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "A Persistent Root Flag.")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag.")
	// if you run > go run main.go echo -h // this flag won't show as its only available for root command.

	// '1' is default value of 't' ie: times
	timesCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout")
	timesCmd.MarkFlagRequired("times") // flags can be marked as required.
	rootCmd.AddCommand(echoCmd)
	echoCmd.AddCommand(timesCmd)
	// timesCmd is subcommand to echoCmd (which is itself subcommand of rootCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// If you use any invalid flag (eg. --invalid), it will print out --help
