package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jzelinskie/cobrautil/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/tupyy/assisted-migration-agent/internal/config"
)

func NewRunCommand(cfg *config.Configuration) *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "A brief description of your command",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("run called")
			return nil
		},
	}

	registerFlags(runCmd, cfg)

	return runCmd
}

func registerFlags(cmd *cobra.Command, config *config.Configuration) {
	nfs := cobrautil.NewNamedFlagSets(cmd)

	serverFlagSet := nfs.FlagSet(color.New(color.FgCyan, color.Bold).Sprint("Server"))
	registerServerFlags(serverFlagSet, config)

	authenticationFlagSet := nfs.FlagSet(color.New(color.FgBlue, color.Bold).Sprint("Authentication"))
	registerAuthenticationFlags(authenticationFlagSet, config)

	nfs.AddFlagSets(cmd)
}

func registerServerFlags(flagSet *pflag.FlagSet, config *config.Configuration) {
	flagSet.StringVar(&config.Mode, "mode", config.Mode, "agent starting mode. Cound be connected to disconnected. Defaults to disconnected")
	flagSet.IntVar(&config.HTTPPort, "http-port", config.HTTPPort, "port on which the HTTP server is listening")
	flagSet.StringVar(&config.StaticsFolder, "statics-folder", config.StaticsFolder, "path to statics")
	flagSet.StringVar(&config.DataFolder, "data-folder", config.DataFolder, "path to the root folder container media")
	flagSet.StringVar(&config.ServerMode, "server-mode", config.ServerMode, "server mode: either prod or dev. If prod it statics folder must be set")
}

func registerAuthenticationFlags(flagSet *pflag.FlagSet, config *config.Configuration) {
	flagSet.BoolVar(&config.Auth.Enabled, "authentication-enabled", config.Auth.Enabled, "enable authentication when connecting to console")
	flagSet.StringVar(&config.Auth.JWTFilePath, "authentication-jwt-filepath", config.Auth.JWTFilePath, "path of the jwt file")
}
