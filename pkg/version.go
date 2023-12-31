// Copyright 2020-2021 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package pkg

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/community-edition/cli/cmd/plugin"
	"github.com/vmware-tanzu/sonobuoy/pkg/buildinfo"
)

// VersionCmd implements a custom version command because we cannot expose both
// TCE and Sonobuoy via a PluginDescriptor's `Version` field, which
// needs a strict semantic version-compatible string.
var VersionCmd = NewCmdVersion()

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print Sonobuoy and Tanzu build information",
		Run:   runVersion(),
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}

func runVersion() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Printf("Sonobuoy Version: %s\n", buildinfo.Version)
		fmt.Printf("TCE Version: %s\n", plugin.Version)
		fmt.Printf("TCE SHA %s\n", plugin.SHA)
	}
}
