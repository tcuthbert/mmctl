// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package comp

func filterArgs(args, opts []string) []string {
	filtered := []string{}
	for _, opt := range opts {
		existing := false
		for _, arg := range args {
			if opt == arg {
				existing = true
			}
		}

		if !existing {
			filtered = append(filtered, opt)
		}
	}

	return filtered
}
