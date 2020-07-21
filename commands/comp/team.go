// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package comp

import (
	"fmt"
	"os"

	"github.com/mattermost/mmctl/client"

	"github.com/spf13/cobra"
)

func Teams(c client.Client, cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	teams, res := c.GetAllTeams("", 0, 200)
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "unable to list teams. Error: %s", res.Error)
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	names := make([]string, len(teams))
	for i, team := range teams {
		names[i] = team.Name
	}

	return filterArgs(args, names), cobra.ShellCompDirectiveNoFileComp
}
