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

// removeparticipatorCmd represents the removeparticipator command
var removeparticipatorCmd = &cobra.Command{
	Use:   "removeparticipator",
	Short: "remove participator(s) in your meeting specified by title",
	Long: `remove participator(s) in your meeting specified by title,just like:
	Agenda removeparticipator -t [title] -p [\"name1, name2\"]`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetStringSlice("participator")
		fmt.Println("removeparticipator called")
		if title == "" || len(participator) == 0 {
			fmt.Println("Please input command like: removeparticipator -t [title] -p [\"name1, name2\"]")
			return
		}
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Please login firstly")
		} else {
			// participators := strings.Split(tmp_p, ",")
			if service.RemoveMeetingParticipator(user.M_name, title, participator) {
				fmt.Println("[remove participator] succeed!")
			} else {
				fmt.Println("[remove participator] error!. Check error.log for detail")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeparticipatorCmd)
	removeparticipatorCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	removeparticipatorCmd.Flags().StringSliceP("participator", "p", nil, "the participator(s) of the meeting, input like \"name1, name2\"")

}
