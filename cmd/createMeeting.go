//输入title participator、starttime、endtime创建一个meeting

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting -t [Title] -p [Participator] -s [StartTime] -e [EndTime]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		title, _ := cmd.Flags().GetString("Title")

		participators, _ := cmd.Flags().GetStringSlice("Participator")

		startTime, _ := cmd.Flags().GetString("StartTime")

		endTime, _ := cmd.Flags().GetString("EndTime")
		fmt.Println("createMeeting called"+title+participators[0]+startTime+endTime)
		//判断是否合法
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)

	createMeetingCmd.Flags().StringP("Title", "t", "", "meeting title")

	createMeetingCmd.Flags().StringSliceP("Participator", "p", []string{}, "meeting's participator")

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
