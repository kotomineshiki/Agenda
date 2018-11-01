//此命令无参数

package cmd

import (
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "登出当前用户",
	Long:  `使用此指令可以退出当前账户`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")
		_, flag := service.GetCurUser()
		if flag != true {
			fmt.Println("[Error]")
			fmt.Println("Cmd logout failed")
			fmt.Println("Not log in yet")
			fmt.Println("Please Log in firstly!")
			return
		}
		if service.UserLogout() {
			fmt.Println("[log out] succeed!")
		} else {
			fmt.Println("[log out] wrong!")
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
