//通过UserName和Password 进行登录
package cmd

import (
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "log in a user ",
	Long: `use username and passward to login ,just like:
	Agenda login -u [UserName] -p [PassWord]`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		//fmt.Println("login called " + username + " " + password)
		if username == "" || password == "" {
			fmt.Println("Please input like,login -u [UserName] -p [PassWord]")
			return
		}
		if isValidName(username) && isValidPassword(password) {
			if service.UserLogin(username, password) {
				fmt.Println("[log in] succeed!")
				fmt.Println("Current user:", username)
			} else {
				fmt.Println("login failed, may be Password error or user doesn't exist")
			}
		}
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

/*func isValidName(n string) bool {
	b := []byte(n)
	val, _ := regexp.Match(".+", b)
	if !val {
		fmt.Println("flag -n ,name is invaild")
	}
	return val
}

func isValidPassword(p string) bool {
	b := []byte(p)
	val, _ := regexp.Match(".+", b)
	if len(p) < 8 {
		fmt.Println("The password must be longer than 8 digits")
		val = false
	}
	if !val {
		fmt.Println("flag -p ,password is invaild")
	}
	return val
}*/
