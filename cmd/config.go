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
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"os"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config [host] [email] [apiToken]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("host, email and apiToken are all required")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		configDir := fmt.Sprintf("%s/.jt", home)
		configFile := fmt.Sprintf("%s%s", configDir, "/conf.yaml")

		fs := afero.NewOsFs()
		afs := &afero.Afero{
			Fs: fs,
		}

		d, err := afs.DirExists(configDir)
		if !d {
			err = fs.Mkdir(configDir, 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		f, err := afs.Exists(configFile)
		if f {
			err = fs.Remove(configFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		content := []byte(fmt.Sprintf(
			"host: %s\nemail: %s\napiToken: %s",
			args[0],
			args[1],
			args[2],
		))

		err = afs.WriteFile(configFile, content, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
