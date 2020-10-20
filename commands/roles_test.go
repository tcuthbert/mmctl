// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package commands

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/mattermost/mmctl/printer"

	"github.com/spf13/cobra"
)

func (s *MmctlUnitTestSuite) TestRolesSystemAdminCmd() {
	s.Run("Make some users sysadmin, some existing and some not", func() {
		printer.Clean()

		user1 := &model.User{Id: "1", Email: "u1@example.com", Roles: "system_user"}
		nonexistentUser2 := "nonexistent"
		user3 := &model.User{Id: "3", Email: "u3@example.com", Roles: "system_user"}
		newRoles := "system_user system_admin"

		s.client.
			EXPECT().
			GetUserByEmail(user1.Email, "").
			Return(user1, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			GetUserByEmail(nonexistentUser2, "").
			Return(nil, &model.Response{Error: &model.AppError{Id: "Mock Error"}}).
			Times(1)

		s.client.
			EXPECT().
			GetUserByUsername(nonexistentUser2, "").
			Return(nil, &model.Response{Error: &model.AppError{Id: "Mock Error"}}).
			Times(1)

		s.client.
			EXPECT().
			GetUser(nonexistentUser2, "").
			Return(nil, &model.Response{Error: &model.AppError{Id: "Mock Error"}}).
			Times(1)

		s.client.
			EXPECT().
			GetUserByEmail(user3.Email, "").
			Return(user3, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			UpdateUserRoles(user1.Id, newRoles).
			Return(true, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			UpdateUserRoles(user3.Id, newRoles).
			Return(true, &model.Response{Error: nil}).
			Times(1)

		args := []string{user1.Email, nonexistentUser2, user3.Email}
		err := rolesSystemAdminCmdF(s.client, &cobra.Command{}, args)
		s.Require().Nil(err)

		s.Require().Len(printer.GetLines(), 2)
		s.Require().Equal(fmt.Sprintf("Updated roles for user %q", user1.Email), printer.GetLines()[0])
		s.Require().Equal(fmt.Sprintf("Updated roles for user %q", user3.Email), printer.GetLines()[1])
		s.Require().Len(printer.GetErrorLines(), 1)
		s.Require().Equal(fmt.Sprintf("unable to find user %q", nonexistentUser2), printer.GetErrorLines()[0])
	})

	s.Run("Make sysadmin a user that already is a sysadmin", func() {
		printer.Clean()

		roles := "system_user system_admin"
		user := &model.User{Id: "1", Email: "u1@example.com", Roles: roles}

		s.client.
			EXPECT().
			GetUserByEmail(user.Email, "").
			Return(user, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			UpdateUserRoles(user.Id, roles).
			Return(true, &model.Response{Error: nil}).
			Times(1)

		err := rolesSystemAdminCmdF(s.client, &cobra.Command{}, []string{user.Email})
		s.Require().Nil(err)

		s.Require().Len(printer.GetLines(), 1)
		s.Require().Equal(fmt.Sprintf("Updated roles for user %q", user.Email), printer.GetLines()[0])
		s.Require().Len(printer.GetErrorLines(), 0)
	})

	s.Run("The update request fails", func() {
		printer.Clean()

		roles := "system_user system_admin"
		user := &model.User{Id: "1", Email: "u1@example.com", Roles: roles}

		s.client.
			EXPECT().
			GetUserByEmail(user.Email, "").
			Return(user, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			UpdateUserRoles(user.Id, roles).
			Return(false, &model.Response{Error: &model.AppError{Id: "Mock Error"}}).
			Times(1)

		err := rolesSystemAdminCmdF(s.client, &cobra.Command{}, []string{user.Email})
		s.Require().Nil(err)

		s.Require().Len(printer.GetLines(), 0)
		s.Require().Len(printer.GetErrorLines(), 1)
		s.Require().Equal(fmt.Sprintf("can't update roles for user %q: : , ", user.Email), printer.GetErrorLines()[0])
	})
}

func (s *MmctlUnitTestSuite) TestRolesMemberCmd() {
	s.Run("Make some users members, some existing and some not", func() {
		printer.Clean()

		user1 := &model.User{Id: "1", Email: "u1@example.com", Roles: "system_user system_admin"}
		nonexistentUser2 := "nonexistent"
		user3 := &model.User{Id: "3", Email: "u3@example.com", Roles: "system_user system_admin"}
		newRoles := "system_user"

		s.client.
			EXPECT().
			GetUserByEmail(user1.Email, "").
			Return(user1, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			GetUserByEmail(nonexistentUser2, "").
			Return(nil, &model.Response{Error: &model.AppError{Id: "Mock Error"}}).
			Times(1)

		s.client.
			EXPECT().
			GetUserByUsername(nonexistentUser2, "").
			Return(nil, &model.Response{Error: &model.AppError{Id: "Mock Error"}}).
			Times(1)

		s.client.
			EXPECT().
			GetUser(nonexistentUser2, "").
			Return(nil, &model.Response{Error: &model.AppError{Id: "Mock Error"}}).
			Times(1)

		s.client.
			EXPECT().
			GetUserByEmail(user3.Email, "").
			Return(user3, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			UpdateUserRoles(user1.Id, newRoles).
			Return(true, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			UpdateUserRoles(user3.Id, newRoles).
			Return(true, &model.Response{Error: nil}).
			Times(1)

		args := []string{user1.Email, nonexistentUser2, user3.Email}
		err := rolesMemberCmdF(s.client, &cobra.Command{}, args)
		s.Require().Nil(err)

		s.Require().Len(printer.GetLines(), 2)
		s.Require().Equal(fmt.Sprintf("Updated roles for user %q", user1.Email), printer.GetLines()[0])
		s.Require().Equal(fmt.Sprintf("Updated roles for user %q", user3.Email), printer.GetLines()[1])
		s.Require().Len(printer.GetErrorLines(), 1)
		s.Require().Equal(fmt.Sprintf("unable to find user %q", nonexistentUser2), printer.GetErrorLines()[0])
	})

	s.Run("Make member a user that it's only a system_user", func() {
		printer.Clean()

		roles := "system_user"
		user := &model.User{Id: "1", Email: "u1@example.com", Roles: roles}

		s.client.
			EXPECT().
			GetUserByEmail(user.Email, "").
			Return(user, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			UpdateUserRoles(user.Id, roles).
			Return(true, &model.Response{Error: nil}).
			Times(1)

		err := rolesMemberCmdF(s.client, &cobra.Command{}, []string{user.Email})
		s.Require().Nil(err)

		s.Require().Len(printer.GetLines(), 1)
		s.Require().Equal(fmt.Sprintf("Updated roles for user %q", user.Email), printer.GetLines()[0])
		s.Require().Len(printer.GetErrorLines(), 0)
	})

	s.Run("The update request fails", func() {
		printer.Clean()

		roles := "system_user"
		user := &model.User{Id: "1", Email: "u1@example.com", Roles: roles}

		s.client.
			EXPECT().
			GetUserByEmail(user.Email, "").
			Return(user, &model.Response{Error: nil}).
			Times(1)

		s.client.
			EXPECT().
			UpdateUserRoles(user.Id, roles).
			Return(false, &model.Response{Error: &model.AppError{Id: "Mock Error"}}).
			Times(1)

		err := rolesMemberCmdF(s.client, &cobra.Command{}, []string{user.Email})
		s.Require().Nil(err)

		s.Require().Len(printer.GetLines(), 0)
		s.Require().Len(printer.GetErrorLines(), 1)
		s.Require().Equal(fmt.Sprintf("can't update roles for user %q: : , ", user.Email), printer.GetErrorLines()[0])
	})
}
