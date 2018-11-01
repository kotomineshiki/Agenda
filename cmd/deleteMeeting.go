//输入一个meeting的名字，删除该meeting
package cmd

import (
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteMeetingCmd represents the deleteMeeting command
var deleteMeetingCmd = &cobra.Command{
	Use:   "deleteMeeting",
	Short: "delete a meeting by its title",
	Long: `delete a meeting just like:
	Agenda deleteMeeting -t [title]`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("Title")
		fmt.Println("deleteMeeting called" + title)
		//查看该Meeting是否存在
		user, flag := service.GetCurUser()
		if flag != true {
			fmt.Println("[Error]")
			fmt.Println("Cmd deleteMeeting failed")
			fmt.Println("Not log in yet")
			fmt.Println("Please Log in firstly!")
			return
		}
		if service.DeleteMeeting(user.M_name, title) {
			fmt.Println("[delete meeting] succeed!")
		} else {
			fmt.Println("[error] delete meeting fail!")
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteMeetingCmd)
	deleteMeetingCmd.Flags().StringP("Title", "t", "", "meeting title")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
