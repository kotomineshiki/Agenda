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

// addparticipatorCmd represents the addparticipator command
var addparticipatorCmd = &cobra.Command{
	Use:   "addparticipator",
	Short: "add participators",
	Long: `add participators to a meeting by meeting's name ,just like:
	agenda addparticipator -t meetingtitle -p "a, b"`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		meetingtitle, _ := cmd.Flags().GetString("title")
		username, _ := cmd.Flags().GetStringSlice("participator")
		if len(username) == 0 || meetingtitle == "" {
			fmt.Println("Please input commend like, aaddparticipator -t meetingtitle -p \"name1, name2\")")
			return
		}
		//fmt.Println("addparticipator called")
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("[Error]")
			fmt.Println("Cmd addparticipator failed")
			fmt.Println("Not log in yet")
			fmt.Println("Please Log in firstly!")
		} else {
			// participators := strings.Split(tmp_p,",")
			flag := service.AddMeetingParticipator(user.GetName(), meetingtitle, username)
			if flag != true {
				fmt.Println("Unexpected error. Check error.log for detail")
			} else {
				fmt.Println("Successfully add")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addparticipatorCmd)
	addparticipatorCmd.Flags().StringP("title", "t", "", "the title of meeting")
	addparticipatorCmd.Flags().StringSliceP("participator", "p", nil, "participator(s) you want to add, input like \"name1, name2\"")
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addparticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addparticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
func isParticipatorinList(name string, participators []string) bool {
	for _, j := range participators {
		if name == j {
			return true
		}
	}
	return false
}
func isParticipatorExist(name string, participators []entity.User) bool {
	for _, j := range participators {
		if name == j.Name {
			return true
		}
	}
	return false
}
func isParticipatorExistinMeeting(name string, meeting entity.Meeting) bool {
	if name == meeting.Sponsor {
		return true
	}
	for _, j := range meeting.Participators {
		if name == j {
			return true
		}
	}
	return false
}*/
