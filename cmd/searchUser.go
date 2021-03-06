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
	//"fmt"

	"agenda/entity"
	"log"

	"github.com/spf13/cobra"
)

// searchUserCmd represents the searchUser command
var searchUserCmd = &cobra.Command{
	Use:   "searchUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_userName_, _ := cmd.Flags().GetString("user")
		users := entity.READUSERS()
		current := entity.GetCurrentUserName()
		if current == "" {
			log.Println("Please log in!")
			return
		}
		//显示所有用户信息
		if _userName_ == "_ALL_" {
			for _, user := range users {
				log.Println("NAME: " + user.Username + "   " + "EMAIL: " + user.Email + "   " + "TEL: " + user.Phone)
			}
		} else { //搜索特定用户
			for _, user := range users {
				if user.Username == _userName_ {
					log.Println("NAME: " + user.Username + "   " + "EMAIL: " + user.Email + "   " + "TEL: " + user.Phone)
					return
				}
			}
			log.Println("User does not exist")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//得到用户名称[-user userName/_ALL_] _ALL_显示所有用户，否则搜索特定用户
	searchUserCmd.Flags().StringP("user", "u", "", "search user participants")
}
