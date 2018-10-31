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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("querymeeting called")
		starttime, _ := cmd.Flags().GetString("starttime")
		endtime, _ := cmd.Flags().GetString("endtime")
		if starttime == "" || endtime == "" {
			fmt.Println("querymeeting -s [starttime] -e [endtime]")
			return
		}
		user, flag := service.GetCurUser()
		if flag != true {
			fmt.Println("Please Log in firstly")
			return
		}
		temp_mets, flag := service.QueryMeeting(user.M_name, starttime, endtime)
		fmt.Println(len(temp_mets))
		if flag == true {
			for _, m := range temp_mets {
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
			fmt.Println("Wrong Date!please input the date as yyyy-mm-dd/hh:mm and make sure that starttiem <= endtime")
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
