//通过UserName和Password 进行登录
package cmd

import (
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login -u [UserName] -p [PassWord]",
	Short: "通过UserName和Password 进行登录",
	Long: `使用 UserName 和 PassWord 来登录Agenda:
如果密码正确，你可以登录，否则必须登记另外一个用户才能使用Agenda`,
	Run: func(cmd *cobra.Command, args []string) {
		//entity读取当前用户？
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		fmt.Println("login called " + username + " " + password)
		login(username, password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "注册过的用户名")
	loginCmd.Flags().StringP("password", "p", "", "用于登录的用户名")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func login(username string, password string) {
	if username == "" || password == "" {
		fmt.Println("Please tell us your username[-u], password[-p]")
		return
	}
	if service.UserLogin(username, password) {
		fmt.Println("[log in] succeed!")
	} else {
		fmt.Println("[log in] Password error or user doesn't exist")
	}
}
