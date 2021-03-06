package root

import (
	"fmt"

	"github.com/jrapoport/gothic/config"
	"github.com/jrapoport/gothic/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	configFile   string
	rootPassword string
	cfg          *config.Config
)

var rootCmd = &cobra.Command{
	Use:               utils.ExecutableName(),
	Version:           config.BuildVersion(),
	RunE:              rootRunE,
	PersistentPreRunE: initConfig,
}

func init() {
	rootCmd.Short = fmt.Sprintf("control plane for %s", config.BuildName())
	pf := rootCmd.PersistentFlags()
	pf.StringVarP(&configFile, "config", "c", "", "the config file to use")
	pf.StringVar(&rootPassword, "root", "", "the root password to use for super admin access")
}

func initConfig(cmd *cobra.Command, _ []string) (err error) {
	cfg, err = config.LoadConfig(configFile, config.SkipRequired())
	if err != nil {
		return err
	}
	cfg.DB.AutoMigrate = false
	cfg.Signup.Default.Username = true
	cfg.Validation.PasswordRegex = ""
	if rootPassword != "" {
		cfg.RootPassword = rootPassword
	}
	return nil
}

func rootRunE(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

// Execute executes the main cmd
func Execute() error {
	return rootCmd.Execute()
}

func Config() *config.Config {
	return cfg
}

func ConfirmAction(format string, a ...interface{}) bool {
	p := fmt.Sprintf(format, a...)
	p = fmt.Sprintf("%s? [Yes/No]", p)
	prompt := promptui.Select{
		Label: p,
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil || result != "Yes" {
		return false
	}
	return true
}
