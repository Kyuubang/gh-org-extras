package team

import (
	"github.com/spf13/cobra"

	teamListCmd "github.com/Kyuubang/gh-org-extras/pkg/cmd/team/list"
)

func NewCmdTeam() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "team <command>",
		Short: "Manage organization teams",
	}

	cmd.AddCommand(teamListCmd.NewListCommand(nil))

	return cmd
}
