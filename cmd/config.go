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
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gopkg.in/go-playground/validator.v9"
	"os"
)

var (
	validate *validator.Validate
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config [host] [email] [apiToken]",
	Short: "Configure credentials for jira",
	Long: `Configure credentials for jira and a user.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("host, email and apiToken are all required")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		validate = validator.New()
		errs := validate.Var(args[0], "required,url")
		if errs != nil {
			color.Red("Invalid url provided")
		}

		errs = validate.Var(args[1], "required,email")
		if errs != nil {
			color.Red("Invalid email provided")
		}

		if errs != nil {
			return
		}

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

		fmt.Println(fmt.Sprintf("Created config file %s%s", configDir, configFile))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
