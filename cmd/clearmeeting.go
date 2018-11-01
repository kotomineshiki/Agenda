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

// clearmeetingCmd represents the clearmeeting command
var clearmeetingCmd = &cobra.Command{
	Use:   "clearmeeting",
	Short: "clear all meetings whose sponsor is the user",
	Long: `claer all meeting which belongs to the user, just like:
	agenda clearmeeting`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("clearmeeting called")
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("[Error]")
			fmt.Println("Cmd clearmeeting failed")
			fmt.Println("Not log in yet")
			fmt.Println("Please Log in firstly!")
		} else {
			cm, flag := service.ClearMeeting(user.GetName())
			if flag == true {
				fmt.Println("Successfully clear ", cm, " meeting(s)")
			} else {
				fmt.Println("ClearMeeting failed. Check error.log")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(clearmeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
