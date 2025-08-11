package list

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Organization string

	Pending bool
	Failed  bool
	Public  bool
}

type MemberList struct {
	Username string `json:"login,omitempty"`
	Email    string `json:"email,omitempty"`
}

func NewListCommand(runF func(*ListOptions) error) *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list [flags]",
		Short: "List members of an organization",
		RunE: func(cmd *cobra.Command, args []string) error {
			if runF != nil {
				return runF(&opts)
			}

			return listRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Pending, "pending", "p", false, "List pending members")
	cmd.Flags().BoolVarP(&opts.Failed, "failed", "f", false, "List failed members")
	cmd.Flags().StringVarP(&opts.Organization, "name", "n", "", "Name of the organization")
	cmd.Flags().BoolVarP(&opts.Public, "public", "u", false, "List public members")

	return cmd
}

func listRun(opts *ListOptions) error {
	var members []MemberList
	var err error

	if opts.Pending {
		members, err = getPendingMembers(opts)
	} else if opts.Failed {
		members, err = getFailedMembers(opts)
	} else if opts.Public {
		members, err = getPublicMembers(opts)
	} else {
		members, err = getMembers(opts)
	}

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Username", "Email"})

	for _, member := range members {
		if member.Username != "" && member.Email != "" {
			table.Append([]string{member.Username, member.Email})
		} else if member.Username != "" {
			table.Append([]string{member.Username, ""})
		} else if member.Email != "" {
			table.Append([]string{"", member.Email})
		}
	}

	table.Render()

	return nil
}

func getMembers(opts *ListOptions) ([]MemberList, error) {
	httpClient, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}

	var response []MemberList

	err = httpClient.Get(fmt.Sprintf("orgs/%s/members?per_page=100&role=member", opts.Organization), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// get pending members
func getPendingMembers(opts *ListOptions) ([]MemberList, error) {
	httpClient, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}

	var response []MemberList

	err = httpClient.Get(fmt.Sprintf("orgs/%s/invitations", opts.Organization), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// get failed members
func getFailedMembers(opts *ListOptions) ([]MemberList, error) {
	httpClient, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}

	var response []MemberList

	err = httpClient.Get(fmt.Sprintf("orgs/%s/failed_invitations", opts.Organization), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// get public members
func getPublicMembers(opts *ListOptions) ([]MemberList, error) {
	httpClient, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}

	var response []MemberList

	err = httpClient.Get(fmt.Sprintf("orgs/%s/public_members", opts.Organization), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
