package cmd

import (
	"containerPicker/server"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	port        string

	rootCmd = &cobra.Command{
		Use:   "containerPicker",
		Short: "Container picker service",
		Long:  `Container picker service`,
		Run: func(cmd *cobra.Command, args []string) {
			myserver := server.NewServer(viper.GetString("port"))
			myserver.Start()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default .containerpicker.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "A.Sosnoviy", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "5555", "port")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("port", rootCmd.PersistentFlags().Lookup("port").DefValue)

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {

		viper.SetConfigType("yaml")
		viper.SetConfigName(".containerpicker")
	}

	viper.SetEnvPrefix("CPICKER")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
