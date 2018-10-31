//注册一个账户

package cmd

import (
	"Agenda/service"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new User",
	Long: `register:register a new User,just like
	register a new user,with name:A, password:12345678
	agenda register -n=A -p=12345678 `,

	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phonenumber, _ := cmd.Flags().GetString("telephone")
		//查看本账户是否已经被注册过
		fmt.Println("register called by " + username + password + email + phonenumber)
		register(username, password, email, phonenumber)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "新用户名")
	registerCmd.Flags().StringP("password", "p", "", "新用户密码")
	registerCmd.Flags().StringP("email", "e", "", "新用户邮箱")
	registerCmd.Flags().StringP("telephone", "t", "", "新用户电话号码")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func register(username string, password string, email string, phonenumber string) {
	if username == "" || password == "" || email == "" || phonenumber == "" {
		fmt.Println("Please tell us your username[-u], password[-p], email[-e], telephone[-t]")
		return
	}

	if isValidName(username) && isValidPassword(password) && isValidEmail(email) && isValidPhone(phonenumber) {
		pass, err := service.UserRegister(username, password, email, phonenumber)
		if pass == false {
			fmt.Println("Username existed!")
			return
		} else {
			if err != nil {
				fmt.Println("Some unexpected error happened when try to record your info,Please read error.log for detail")
				return
			} else {
				fmt.Println("Successfully register!")
			}
		}
	}
}

func isValidName(n string) bool {
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
}
func isValidEmail(e string) bool {
	b := []byte(e)
	val, _ := regexp.Match("\\w*@\\w*\\.w*", b)

	if !val {
		fmt.Println("email is invaild")
	}
	return val
}
func isValidPhone(p string) bool {
	b := []byte(p)

	val, _ := regexp.Match("[0-9]+", b)

	if !val {
		fmt.Println("phone is invaild")
	}
	return val
}
