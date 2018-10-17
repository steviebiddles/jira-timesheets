// Copyright © 2018 Stephen McAuley <steviebiddles@gmail.com>
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
	tm "github.com/buger/goterm"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/steviebiddles/jira-timesheets/clients"
	"github.com/steviebiddles/jira-timesheets/models"
	"strings"
	"time"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [issue]",
	Short: "List worklog for an issue",
	Long:  `Complete list of all worklogs for an issue`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("issue must be provided")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		issue := args[0]

		fmt.Println("---")
		worklogs := clients.GetIssueWorklogs(issue)
		fmt.Println("---")
		fmt.Println("Issue:", issue)
		fmt.Println()

		displayList(*worklogs)
		fmt.Println("Total:", worklogs.Total)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func displayList(worklogs models.Worklogs) {
	totals := tm.NewTable(0, 10, 5, ' ', 0)

	_, _ = fmt.Fprintf(totals, "Id\tAuthor\tDate\tTime Spent\tComments\n")
	_, _ = fmt.Fprintf(totals, "---\t---\t---\t---\t---\n")

	for _, w := range worklogs.Worklogs {
		_, _ = fmt.Fprintf(
			totals,
			"%s\t%s\t%s\t%s\t%s\n",
			w.Id,
			w.Author.DisplayName,
			issueDate(w.Started),
			w.TimeSpent,
			issueComments(w.Comment),
		)
	}

	_, _ = tm.Print(totals)
	tm.Flush()
}

func issueDate(startedDate string) string {
	d, _ := time.Parse("2006-01-02T15:04:05.999-0700", startedDate)

	return d.Format("Mon, 02 Jan 2006")
}

func issueComments(comment models.Comment) string {
	comments := ""

	for _, contents := range comment.Content {
		for _, content := range contents.Content {
			comments += strings.TrimRight(fmt.Sprintf("%s ", content.Text), "")
		}
	}

	return comments
}
