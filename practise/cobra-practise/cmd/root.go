package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string

	author string
	Verbose bool
	Source string
	Region string
)

var (
	rootCmd = &cobra.Command{
	Use:   "rootCmd [OPTIONS] [COMMANDS]",
	Short: "A short rootCmd demo",
	Long: `A Long rootCmd demo`,

	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exec rootCmd")
	    },
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 在 init() 函数中定义 flags 和 处理配置文件
	cobra.OnInitialize(initConfig)

	// Persistent Flags：全局 flag，指定命令和它的子命令均可用
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// Local Flags: 只能指定给特定的 command
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "source directory to read from")

	// Bind Flags with Config（vipers）
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "caoyingjun", "Author name for the demo")
    viper.BindPFlag("auther", rootCmd.PersistentFlags().Lookup("auther"))

	// 全局必须的flag
	//rootCmd.PersistentFlags().StringVarP(&Region, "region", "r", "", "global region (required)")
	//rootCmd.MarkPersistentFlagRequired("region")

    // local 必须的flag
	//rootCmd.Flags().StringVarP(&Region, "region", "r", "", "local region (required)")
	//rootCmd.MarkFlagRequired("region")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}