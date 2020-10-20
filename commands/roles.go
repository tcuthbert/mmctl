// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package commands

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/mattermost/mmctl/client"
	"github.com/mattermost/mmctl/printer"

	"github.com/spf13/cobra"
)

var RolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "Management of roles",
}

var RolesSystemAdminCmd = &cobra.Command{
	Use:     "system_admin [users]",
	Short:   "Set a user as system admin",
	Long:    "Make some users system admins",
	Example: "  roles system_admin user1",
	RunE:    withClient(rolesSystemAdminCmdF),
	Args:    cobra.MinimumNArgs(1),
}

var RolesMemberCmd = &cobra.Command{
	Use:     "member [users]",
	Short:   "Remove system admin privileges",
	Long:    "Remove system admin privileges from some users.",
	Example: "  roles member user1",
	RunE:    withClient(rolesMemberCmdF),
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	RolesCmd.AddCommand(
		RolesSystemAdminCmd,
		RolesMemberCmd,
	)

	RootCmd.AddCommand(RolesCmd)
}

func rolesSystemAdminCmdF(c client.Client, _ *cobra.Command, args []string) error {
	users := getUsersFromUserArgs(c, args)
	for i, user := range users {
		if user == nil {
			printer.PrintError(fmt.Sprintf("unable to find user %q", args[i]))
			continue
		}

		shouldAddSystemAdmin := true
		shouldAddSystemUser := true

		roles := strings.Fields(user.Roles)
		for _, role := range roles {
			switch role {
			case model.SYSTEM_ADMIN_ROLE_ID:
				shouldAddSystemAdmin = false
			case model.SYSTEM_USER_ROLE_ID:
				shouldAddSystemUser = false
			}
		}

		if shouldAddSystemUser {
			roles = append(roles, model.SYSTEM_USER_ROLE_ID)
		}
		if shouldAddSystemAdmin {
			roles = append(roles, model.SYSTEM_ADMIN_ROLE_ID)
		}

		_, resp := c.UpdateUserRoles(user.Id, strings.Join(roles, " "))
		if resp.Error != nil {
			printer.PrintError(fmt.Sprintf("can't update roles for user %q: %s", args[i], resp.Error))
			continue
		}

		printer.Print(fmt.Sprintf("Updated roles for user %q", args[i]))
	}

	return nil
}

func rolesMemberCmdF(c client.Client, _ *cobra.Command, args []string) error {
	users := getUsersFromUserArgs(c, args)
	for i, user := range users {
		if user == nil {
			printer.PrintError(fmt.Sprintf("unable to find user %q", args[i]))
			continue
		}

		shouldAddSystemUser := true
		var newRoles []string

		roles := strings.Fields(user.Roles)
		for _, role := range roles {
			switch role {
			case model.SYSTEM_ADMIN_ROLE_ID:
			default:
				if role == model.SYSTEM_USER_ROLE_ID {
					shouldAddSystemUser = false
				}
				newRoles = append(newRoles, role)
			}
		}

		if shouldAddSystemUser {
			newRoles = append(roles, model.SYSTEM_USER_ROLE_ID)
		}

		_, resp := c.UpdateUserRoles(user.Id, strings.Join(newRoles, " "))
		if resp.Error != nil {
			printer.PrintError(fmt.Sprintf("can't update roles for user %q: %s", args[i], resp.Error))
			continue
		}

		printer.Print(fmt.Sprintf("Updated roles for user %q", args[i]))
	}

	return nil
}
