package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"golang-learning/practise/cobra-practise/imp"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "demo",
	Short: "A test demo",
	Long: `Demo is a test appcation for print things.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		imp.Show("caoyingjun", 18)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(".demo")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	RootCmd.Flags().BoolP("toggle", "t", false, "help message for toggle")

}
