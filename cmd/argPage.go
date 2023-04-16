/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// argPageCmd represents the argPage command
var argPageCmd = &cobra.Command{
	Use:   "argPage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//检测是否输入args没有输入args就打印帮助
		if len(args) == 0 && len(os.Args) == 1 {
			// 如果没有提供命令或标志，则打印所有标志的帮助信息
			cmd.HelpFunc()(cmd, args)
			os.Exit(0)
		}
	},
	Run: runArgPage,
}

func init() {
	rootCmd.AddCommand(argPageCmd)
	//这里添加argPageCmd的标志位置
	argPageCmd.PersistentFlags().StringP("All", "a", "", "All to use")
}

func runArgPage(cmd *cobra.Command, args []string) {

	//处理All
	All, err := cmd.Flags().GetString("All")
	if err != nil {
		// Handle error
	}
	if All != "" {
		fmt.Println("All is:", All)
	}
}
