//注册一个账户

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username,_:=cmd.Flags().GetString("username")
		password,_:=cmd.Flags().GetString("password")
		email,_:=cmd.Flags().GetString("email")
		phonenumber,_:=cmd.Flags().GetString("phonenumber")
	//查看本账户是否已经被注册过
		fmt.Println("register called by"+username+password+email+phonenumber)

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username","u","Anonymous", "新用户名")
	registerCmd.Flags().StringP("password","p","admin", "新用户密码")
	registerCmd.Flags().StringP("email","e","", "新用户邮箱")
	registerCmd.Flags().StringP("phonenumber","pn","", "新用户电话号码")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
