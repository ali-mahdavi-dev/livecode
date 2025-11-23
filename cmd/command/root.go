package command

import (
	"fmt"
	"livecode/config"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	cfg     config.Config
	envFile string
	rootCmd = &cobra.Command{
		Use: "",
		Run: func(cmd *cobra.Command, args []string) {
			initializeConfigs()
		},
	}
)

func initializeConfigs() {
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

	c, err := config.Load()
	if err != nil {
		log.Fatalf("could not load configuration %s\n", err.Error())
	}

	cfg = *c
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVarP(&envFile, "env-file", "e", ".env", ".env file")

	rootCmd.AddCommand(httpCmd())
	rootCmd.AddCommand(migrateCmd())
	rootCmd.AddCommand(readUserFromJsonCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
