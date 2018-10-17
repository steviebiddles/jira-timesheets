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
	"github.com/spf13/viper"
	"github.com/steviebiddles/jira-timesheets/clients"
	"github.com/steviebiddles/jira-timesheets/models"
	"log"
	"time"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [issue]",
	Short: "Add a worklog entry for an issue",
	Long: `Add a worklog entry for an issue and auto adjust the remaining estimate`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("issue must be provided")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var started string
		var timeSpentSeconds int

		issue := args[0]

		fmt.Println("---")

		if viper.GetString("date") != "" {
			t, err := time.Parse("2006-01-02 15:04:05",  viper.GetString("date") + " 12:01:00")
			if err != nil {
				log.Fatal(err)
			}

			started = t.Format("2006-01-02T15:04:05.999-0700")
		} else {
			t := time.Now().UTC()
			started = t.Format("2006-01-02T15:04:05.999-0700")
		}

		timeSpentSeconds = getTimeSpentInSeconds(viper.GetFloat64("hours"), viper.GetFloat64("minutes"))
		payload := models.Worklog{
			Comment: models.Comment{
				Content: []models.Content{
					{
						Type: "paragraph",
						Content: []models.Content{
							{
								Type: "text",
								Text: viper.GetString("comment"),
							},
						},
					},
				},
			},
			Started: started,
			TimeSpentSeconds: timeSpentSeconds,
		}

		clients.PostIssueWorklog(issue, payload)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("date","d", "", "start date for worklog entry e.g. \"2018-01-01\". defaults to current date")
	addCmd.Flags().Float64P("hours","H", 0, "hours spent on issue")
	addCmd.Flags().Float64P("minutes","m", 0, "minutes spent on issue")
	addCmd.Flags().StringP("comment","c", "", "comments")

	_ = viper.BindPFlag("date", addCmd.Flags().Lookup("date"))
	_ = viper.BindPFlag("hours", addCmd.Flags().Lookup("hours"))
	_ = viper.BindPFlag("minutes", addCmd.Flags().Lookup("minutes"))
	_ = viper.BindPFlag("comment", addCmd.Flags().Lookup("comment"))
}

func getTimeSpentInSeconds(h float64, m float64) (s int) {
	if h > 0 {
		s += int(h * 60 * 60)
	}

	if m > 0 {
		s += int(m * 60)
	}

	return
}