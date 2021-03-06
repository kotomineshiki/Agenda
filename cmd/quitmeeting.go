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
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// quitmeetingCmd represents the quitmeeting command
var quitmeetingCmd = &cobra.Command{
	Use:   "quitmeeting",
	Short: "quit a meeting by its title",
	Long: `quit a meeting by its title,just like:
	Agenda quitmeeting -t [title]`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		//fmt.Println("quitmeeting called")
		if title == "" {
			fmt.Println("Please input meeting's title you want to quit")
			return
		}
		user, flag := service.GetCurUser()
		if flag != true {
			fmt.Println("[Error]")
			fmt.Println("Cmd quitmeeting failed")
			fmt.Println("Not log in yet")
			fmt.Println("Please Log in firstly!")
		} else {
			if service.QuitMeeting(user.M_name, title) {
				fmt.Println("[quit meeting] succeed!")
			} else {
				fmt.Println("quit meeting failed!")
				fmt.Println("Please read error.log for detail")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(quitmeetingCmd)
	quitmeetingCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
