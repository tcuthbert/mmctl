// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package comp

import (
	"fmt"
	"os"

	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/mattermost/mmctl/client"

	"github.com/spf13/cobra"
)

func InstalledPlugins(c client.Client, cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	pluginsResp, res := c.GetPlugins()
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "unable to list plugins. Error: %s", res.Error)
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	ids := []string{}
	for _, plugin := range pluginsResp.Active {
		ids = append(ids, plugin.Manifest.Id)
	}
	for _, plugin := range pluginsResp.Inactive {
		ids = append(ids, plugin.Manifest.Id)
	}

	return filterArgs(args, ids), cobra.ShellCompDirectiveNoFileComp
}

func EnabledPlugins(c client.Client, cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	pluginsResp, res := c.GetPlugins()
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "unable to list plugins. Error: %s", res.Error)
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	ids := make([]string, len(pluginsResp.Active))
	for i, plugin := range pluginsResp.Active {
		ids[i] = plugin.Manifest.Id
	}

	return filterArgs(args, ids), cobra.ShellCompDirectiveNoFileComp
}

func DisabledPlugins(c client.Client, cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	pluginsResp, res := c.GetPlugins()
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "unable to list plugins. Error: %s", res.Error)
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	ids := make([]string, len(pluginsResp.Inactive))
	for i, plugin := range pluginsResp.Inactive {
		ids[i] = plugin.Manifest.Id
	}

	return filterArgs(args, ids), cobra.ShellCompDirectiveNoFileComp
}

func MarketplacePlugins(c client.Client, cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		pluginFilter := &model.MarketplacePluginFilter{PerPage: 200, Filter: toComplete}
		plugins, res := c.GetMarketplacePlugins(pluginFilter)
		if res.Error != nil {
			fmt.Fprintf(os.Stderr, "unable to list marketplace plugins. Error: %s", res.Error)
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		ids := make([]string, len(plugins))
		for i, plugin := range plugins {
			ids[i] = plugin.Manifest.Id
		}

		return ids, cobra.ShellCompDirectiveNoFileComp
	}

	return nil, cobra.ShellCompDirectiveNoFileComp
}
