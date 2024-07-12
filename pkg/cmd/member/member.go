package member

import (
	memberInviteCmd "github.com/Kyuubang/gh-org-extras/pkg/cmd/member/invite"
	"github.com/spf13/cobra"
)

func NewCmdMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member <command>",
		Short: "Manage organization members",
	}

	cmd.AddCommand(memberInviteCmd.NewCmdInvite(nil))

	return cmd
}
