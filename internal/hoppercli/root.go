// Copyright 2024 The Contributors to go-eth-app
// This file is part of the go-eth-app project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the License.
//
// For a list of contributors, refer to the CONTRIBUTORS file or the
// repository's commit history.

package hoppercli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	solBinPath    string   // Path to compiled solidity
	solABIPath    string   // Path to ABI content
	constructArgs []string // Arguments for constructor
)

var (
	devDeployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "deploy is a command to deploy contract",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("bin", solBinPath)
			fmt.Println("abi", solABIPath)
			fmt.Println("constructor", constructArgs)
		},
	}

	devCmd = &cobra.Command{
		Use:   "dev",
		Short: "dev is a command to support interactions with Geth dev node",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func setupDevCmd() {
	devCmd.AddCommand(devDeployCmd)
	devDeployCmd.Flags().StringVarP(&solBinPath, "binary", "b", "", "path to solidity binary")
	devDeployCmd.MarkFlagRequired("binary")
	devDeployCmd.Flags().StringVarP(&solABIPath, "abi", "a", "", "path to ABI file")
	devDeployCmd.MarkFlagRequired("abi")
	devDeployCmd.Flags().StringArrayVarP(&constructArgs, "cargs", "c", nil, "array of constructor argments in string")
}

var rootCmd = &cobra.Command{
	Use:   "hopper",
	Short: "hopper is a cli app to help you deploy a Solidity contract.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func setUpRootCmd() {
	rootCmd.AddCommand(devCmd)
}

func init() {
	setupDevCmd()
	setUpRootCmd()
}

func Execute() error {
	return rootCmd.Execute()
}
