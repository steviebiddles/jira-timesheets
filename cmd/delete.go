// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/steviebiddles/jira-timesheets/clients"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [issue] [worklogId]",
	Short: "Delete a worklog entry for an issue",
	Long: `Delete a worklog entry for an issue and auto adjust the remaining estimate`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("issue and worklogId are required")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---")
		clients.DeleteIssueWorklog(args[0], args[1])
		fmt.Println("---")
		fmt.Println("Issue:", args[0])
		fmt.Println("Worklog ID:", args[1])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
