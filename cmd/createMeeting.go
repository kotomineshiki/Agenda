//输入title participator、starttime、endtime创建一个meeting

package cmd

import (
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting -t [Title] -p [Participator] -s [StartTime] -e [EndTime]",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("Title")
		participators, _ := cmd.Flags().GetStringSlice("Participator")
		startTime, _ := cmd.Flags().GetString("StartTime")
		endTime, _ := cmd.Flags().GetString("EndTime")
		if title == "" || len(participators) == 0 || startTime == "" || endTime == "" {
			fmt.Println("createMeeting -t [Title] -p [Participator] -s [StartTime] -e [EndTime]")
			return
		}
		fmt.Println("createMeeting called" + title + startTime + endTime)
		//判断是否合法
		user, flag := service.GetCurUser()
		if flag != true {
			fmt.Println("Error: please login firstly")
			return
		}
		if service.CreateMeeting(user.M_name, title, startTime, endTime, participators) {
			fmt.Println("[create meeting] succeed!")
			return
		} else {
			fmt.Println("[create meeting] error!")
		}
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)

	createMeetingCmd.Flags().StringP("Title", "t", "", "meeting title")
	createMeetingCmd.Flags().StringSliceP("Participator", "p", nil, "meeting's participator")
	createMeetingCmd.Flags().StringP("StartTime", "s", "", "meeting's startTime")
	createMeetingCmd.Flags().StringP("EndTime", "e", "", "meeting's endTime")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
