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
	"Agenda/entity"
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// querymeetingCmd represents the querymeeting command
var querymeetingCmd = &cobra.Command{
	Use:   "querymeeting",
	Short: "query meetings from starttime to endtime ",
	Long: `query meetings from starttime to endtime, just like:
	Agenda querymeeting -s [yyyy-mm-dd/hh:mm] -e [yyyy-mm-dd/hh:mm]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("querymeeting called")
		starttime, _ := cmd.Flags().GetString("starttime")
		endtime, _ := cmd.Flags().GetString("endtime")
		if starttime == "" || endtime == "" {
			fmt.Println("querymeeting -s [yyyy-mm-dd/hh:mm] -e [yyyy-mm-dd/hh:mm]")
			return
		}
		user, flag := service.GetCurUser()
		if flag != true {
			fmt.Println("Please Log in firstly")
			return
		}
		meetings, flag := service.QueryMeeting(user.M_name, starttime, endtime)
		fmt.Println(len(meetings))
		if flag == true {
			for _, m := range meetings {
				fmt.Println("----------------")
				fmt.Println("Title: ", m.M_title)
				ts, _ := entity.DateToString(m.M_startDate)
				fmt.Println("Start Time", ts)
				te, _ := entity.DateToString(m.M_endDate)
				fmt.Println("End Time", te)
				fmt.Printf("Participator(s): ")
				for _, p := range m.M_participators {
					fmt.Printf(p, " ")
				}
				fmt.Printf("\n")
				fmt.Println("----------------")
			}
		} else {
			fmt.Println("Please check your input the date as yyyy-mm-dd/hh:mm and make sure that starttiem <= endtime")
		}
	},
}

func init() {
	rootCmd.AddCommand(querymeetingCmd)
	querymeetingCmd.Flags().StringP("starttime", "s", "", "starttime of time interval")
	querymeetingCmd.Flags().StringP("endtime", "e", "", "endtime of time interval")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// querymeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// querymeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
