package remove

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

type BulkOptions struct {
	Filenames    string
	Teams        string
	Organization string
}

type TeamMember struct {
	Username string `json:"login,omitempty"`
	IsAdmin  bool   `json:"site_admin,omitempty"`
}

func NewCmdBulkRemove(runF func(*BulkOptions) error) *cobra.Command {
	var opts BulkOptions

	cmd := &cobra.Command{
		Use:   "remove [organization name]",
		Short: "Bulk remove users from an organization",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("organization name is required")
			}

			opts.Organization = args[0]

			if runF != nil {
				return runF(&opts)
			}

			return bulkDeleteRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Filenames, "file", "f", "", "Filename containing list of users to invite")
	cmd.Flags().StringVarP(&opts.Teams, "teams", "t", "", "Teams to add the user to")

	return cmd
}

func bulkDeleteRun(opts *BulkOptions) error {
	// Implement the logic to bulk delete users from an organization
	// This function will read the file specified in opts.Filenames,
	// parse the user data, and then remove them from the organization
	// specified in opts.Organization.

	if opts.Teams != "" {
		// If teams are specified, get the team member by slug
		teamMembers, err := getTeamMemberBySlug(opts)
		if err != nil {
			return err
		}

		// get confirmation before proceeding with deletion
		// and print list of users to be deleted
		fmt.Println("The following users will be deleted from the organization:")
		for _, member := range teamMembers {
			if !member.IsAdmin {
				fmt.Println("-", member.Username)
			}
		}

		// get confirmation from the user
		var confirmation string
		fmt.Print("Are you sure you want to delete these users? (yes/no): ")
		fmt.Scanln(&confirmation)
		if confirmation != "yes" {
			fmt.Println("Operation cancelled.")
			return nil
		}

		for _, member := range teamMembers {
			// Delete each user from the organization
			if !member.IsAdmin {
				fmt.Println("Deleting user:", member.Username)
				err := deleteUserFromOrg(opts.Organization, member.Username)
				if err != nil {
					return err
				}
				fmt.Println("User deleted:", member.Username)
			}
		}
	} else if opts.Filenames != "" {
		// If filenames are specified, read the file and delete each user
		// Implement the logic to read the file and delete each user
		// read the file specified in opts.Filenames
		// parse the user data, and then remove them from the organization

		// warning this destructive operation will delete users from the organization may even admin
		fmt.Println("Warning: This operation will delete users from the organization, including admins. Proceed with caution.")

		file, err := os.Open(opts.Filenames)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", opts.Filenames, err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			username := scanner.Text()
			if username == "" {
				continue // Skip empty lines
			}

			// Delete each user from the organization
			err := deleteUserFromOrg(opts.Organization, username)
			if err != nil {
				return fmt.Errorf("failed to delete user %s: %w", username, err)
			}
		}
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading file %s: %w", opts.Filenames, err)
		}
	} else {
		return errors.New("either teams or filenames must be specified")
	}

	return nil
}

func deleteUserFromOrg(org string, username string) error {
	// Implement the logic to delete a user from an organization
	// This function will interact with the GitHub API to remove the user from the organization.
	var resp interface{}

	httpClient, err := api.DefaultRESTClient()
	if err != nil {
		return err
	}

	err = httpClient.Delete(fmt.Sprintf("orgs/%s/members/%s", org, username), &resp)
	if err != nil {
		return fmt.Errorf("failed to delete user %s from organization %s: %w", username, org, err)
	}

	// Successfully deleted the user
	// Return nil to indicate success

	return nil
}

func getTeamMemberBySlug(opts *BulkOptions) (teamMembers []TeamMember, err error) {
	httpClient, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}

	// Implement the logic to get a team member by their slug
	// This function will interact with the GitHub API to retrieve the team member's details.

	err = httpClient.Get(fmt.Sprintf("orgs/%s/teams/%s/members?role=member", opts.Organization, opts.Teams), &teamMembers)
	if err != nil {
		return nil, err
	}

	return teamMembers, nil
}
