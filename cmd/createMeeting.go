//输入title participator、starttime、endtime创建一个meeting

package cmd

import (
	"Agenda/service"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "create a new meeting",
	Long: `create a meeting just like:
	agenda createMeeting -t [Title] -p [\"name1, name2\"] -s [yyyy-mm-dd/hh:mm] -e [yyyy-mm-dd/hh:mm]`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("Title")
		participators, _ := cmd.Flags().GetStringSlice("Participator")
		startTime, _ := cmd.Flags().GetString("StartTime")
		endTime, _ := cmd.Flags().GetString("EndTime")
		if title == "" || len(participators) == 0 || startTime == "" || endTime == "" {
			fmt.Println("createMeeting -t [Title] -p [\"name1, name2\"] -s [yyyy-mm-dd/hh:mm] -e [yyyy-mm-dd/hh:mm]")
			return
		}
		//fmt.Println("createMeeting called" + title + startTime + endTime)
		//判断是否合法
		user, flag := service.GetCurUser()
		if flag != true {
			fmt.Println("[Error]")
			fmt.Println("Cmd createMeeting failed")
			fmt.Println("Not log in yet")
			fmt.Println("Please Log in firstly!")
			return
		}

		if isValidDate(startTime) && isValidDate(endTime) {
			//在CreateMeeting函数中包含对于日期合法性的检测
			if service.CreateMeeting(user.M_name, title, startTime, endTime, participators) {
				fmt.Println("[create meeting] succeed!")
				fmt.Println("Current user: ", user.GetName())
				return
			} else {
				fmt.Println("[create meeting] error!")
			}
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

func isValidDate(e string) bool {
	b := []byte(e)
	val, _ := regexp.Match("^[1-2][0-9][0-9][0-9]-([1][0-2]|0?[1-9])-([12][0-9]|3[01]|0?[1-9])/([01][0-9]|[2][0-3]):[0-5][0-9]$", b)

	if !val {
		fmt.Println("data is invaild")
	}
	return val
}
