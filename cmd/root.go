package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	//usage标志位
	Use: "./awesomeProject",
	//简短的描述
	Short: "A brief description of your application",
	//详细的描述
	Long: `A longer description that spans multiple lines and likely contains
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//检测是否输入args没有输入args就打印帮助
		if len(args) == 0 && len(os.Args) == 1 {
			// 如果没有提供命令或标志，则打印所有标志的帮助信息
			cmd.HelpFunc()(cmd, args)
			os.Exit(0)
		}
	},
	Run: run,
}

//root的入口，在main.go中调用cmd.Execute()
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// 增加标志位
func init() {
	// 增加flags标志位
	rootCmd.Flags().StringP("location", "l", "", "Location to use")
	rootCmd.Flags().StringP("url", "u", "", "select url")
}

// 处理标志位指定的值
func run(cmd *cobra.Command, args []string) {

	//处理location
	location, err := cmd.Flags().GetString("location")
	if err != nil {
		// Handle error
	}
	if location != "" {
		fmt.Println("Location is:", location)
	}

	//处理url
	url, err := cmd.Flags().GetString("url")
	if err != nil {
	}
	if url != "" {
		fmt.Println("url is" + url)
	}

}
