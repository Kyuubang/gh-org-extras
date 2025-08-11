package bulk

import (
	bulkRemoveCmd "github.com/Kyuubang/gh-org-extras/pkg/cmd/bulk/remove"

	"github.com/spf13/cobra"
)

func NewCmdBulk() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bulk <command>",
		Short: "Manage bulk operations",
	}

	cmd.AddCommand(bulkRemoveCmd.NewCmdBulkRemove(nil))

	return cmd
}
