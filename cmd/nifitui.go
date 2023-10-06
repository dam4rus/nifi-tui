package main

import (
	"dam4rus/nifi-tui/internal/pkg/nifitui"
	"net/url"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "nifi-tui",
		Short: "TUI for managing NiFi instances",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		RunE: func(cmd *cobra.Command, args []string) error {
			nifiUrl, err := cmd.Flags().GetString("url")
			if err != nil {
				return err
			}
			userName, err := cmd.Flags().GetString("user-name")
			if err != nil {
				return err
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				return err
			}
			nifiApiUrl, err := url.JoinPath(nifiUrl, "nifi-api")
			if err != nil {
				return err
			}
			app, err := nifitui.NewApp(nifiApiUrl, userName, password)
			if err != nil {
				return err
			}
			app.EnterRootProcessGroupScreen()
			if err := app.Run(); err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.Flags().String("url", "https://localhost:8443", "URL of the Nifi server")
	rootCmd.Flags().String("user-name", "", "User name to login with")
	rootCmd.Flags().String("password", "", "Password to login with")
	rootCmd.MarkFlagRequired("username")
	rootCmd.MarkFlagRequired("password")
	rootCmd.Execute()
}
