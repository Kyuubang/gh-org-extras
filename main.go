package main

import (
	"github.com/Kyuubang/gh-org-extras/pkg/cmd/member"
	"github.com/spf13/cobra"
)

// create struct to handle response from the API

func main() {
	// create a new root command
	cmd := rootCmd()

	// execute the command
	cmd.Execute()

}

func rootCmd() *cobra.Command {
	// create a new command
	cmd := &cobra.Command{
		Use:   "gh-org-extras",
		Short: "CLI tool for managing GitHub organizations",
	}

	// add subcommands
	cmd.AddCommand(member.NewCmdMember())

	return cmd
}
