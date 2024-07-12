package invite

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/mail"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

type InviteOptions struct {
	Filenames    string
	Teams        []int
	Role         string
	Organization string
	Email        string
	Username     string
}

type InviteBody struct {
	InviteeID string `json:"invitee_id,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	TeamIDs   []int  `json:"team_ids,omitempty"`
}

// TODO: improve ux by adding consise help message and required flags
func NewCmdInvite(runF func(*InviteOptions) error) *cobra.Command {
	var opts InviteOptions

	cmd := &cobra.Command{
		Use:   "invite",
		Short: "Invite a user to an organization",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("organization name is required")
			}

			opts.Organization = args[0]

			if runF != nil {
				return runF(&opts)
			}

			return inviteRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Filenames, "file", "f", "", "Filename containing list of users to invite")
	cmd.Flags().IntSliceVarP(&opts.Teams, "teams", "t", nil, "Teams to add the user to")
	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role to assign the user")
	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "Username of the user to invite")
	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email of the user to invite")

	return cmd
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func inviteRun(opts *InviteOptions) error {
	if opts.Filenames != "" {
		lines, err := readLines(opts.Filenames)
		if err != nil {
			return err
		}

		for _, line := range lines {
			if isValidEmail(line) {
				opts.Email = line
			} else {
				opts.Username = line
			}

			err := invite(opts)
			if err != nil {
				return err
			}
		}

		return nil
	} else {
		return invite(opts)
	}
}

func invite(opts *InviteOptions) error {
	httpClient, err := api.DefaultRESTClient()
	if err != nil {
		return err
	}

	var response interface{}

	var body = InviteBody{
		Email:     opts.Email,
		Role:      opts.Role,
		InviteeID: opts.Username,
		TeamIDs:   opts.Teams,
	}

	// convert body to json
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// print json body
	fmt.Println(string(bodyJSON))

	// convert body to json
	err = httpClient.Post(fmt.Sprintf("orgs/%s/invitations", opts.Organization), bytes.NewBuffer(bodyJSON), &response)
	if err != nil {
		return err
	}

	if opts.Username != "" {
		fmt.Printf("✓ Invitation sent to %s\n", opts.Username)

	} else if opts.Email != "" {
		fmt.Printf("✓ Invitation sent to %s\n", opts.Email)
	}

	return nil
}

// create function to read file and return slice of string
func readLines(filename string) ([]string, error) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create slice of string
	var lines []string

	// create scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// return slice of string
	return lines, scanner.Err()
}
