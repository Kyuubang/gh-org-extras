package member

import (
	memberInviteCmd "github.com/Kyuubang/gh-org-extras/pkg/cmd/member/invite"
	memberListCmd "github.com/Kyuubang/gh-org-extras/pkg/cmd/member/list"

	"github.com/spf13/cobra"
)

func NewCmdMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member <command>",
		Short: "Manage organization members",
	}

	cmd.AddCommand(memberInviteCmd.NewCmdInvite(nil))
	cmd.AddCommand(memberListCmd.NewListCommand(nil))

	return cmd
}
