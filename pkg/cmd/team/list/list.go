package list

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	HttpClient *api.RESTClient

	Organization string
	Team         string
}

type Team struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

func NewListCommand(runF func(*ListOptions) error) *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list [flags]",
		Short: "List teams of an organization",
		RunE: func(cmd *cobra.Command, args []string) error {
			if runF != nil {
				return runF(&opts)
			}

			return listRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Organization, "name", "n", "", "Name of the organization")
	cmd.Flags().StringVarP(&opts.Team, "team", "t", "", "Name of the team")

	return cmd
}

func listRun(opts *ListOptions) error {
	if opts.Organization == "" {
		return fmt.Errorf("organization name is required")
	}

	teams, err := getTeams(opts)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Slug"})

	for _, team := range teams {
		table.Append([]string{fmt.Sprintf("%d", team.ID), team.Name, team.Slug})
	}

	table.Render()

	return nil
}

func getTeams(opts *ListOptions) ([]Team, error) {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}

	var teams []Team

	err = client.Get("orgs/"+opts.Organization+"/teams", &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}
