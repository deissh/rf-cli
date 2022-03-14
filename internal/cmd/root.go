package cmd

import (
	"github.com/deissh/rf-cli/internal/build"
	configCmd "github.com/deissh/rf-cli/internal/cmd/config"
	extensionCmd "github.com/deissh/rf-cli/internal/cmd/extension"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/log"
	"github.com/spf13/cobra"
)

var (
	configPath string
	debug      bool
)

func init() {
	cobra.OnInitialize(func() {
		if debug {
			log.Level = log.DebugLevel
		}

		log.Debug("RF CLI version %s (%s)", build.Version, build.Date)

		path := config.GetConfigFile()

		if configPath != "" {
			path = configPath
		}

		if !utils.FileExists(path) {
			log.Warn("Missing configuration file.")
			log.Warn("Run 'rf config' to configure the tool.")
			return
		}

		if err := config.Load(path); err != nil {
			log.Warn("Config not loaded, %e\n", err)
			return
		}

		log.Debug("Using config file: %s", config.GetConfigFile())
	})
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rf <command> <subcommand> [flags]",
		Short:   "CLI include some shortcuts for RedForester",
		Version: build.Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVarP(
		&configPath, "config", "c", config.GetConfigFile(),
		"Config file",
	)
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "Turn on debug output")

	addChildCommands(cmd)

	return cmd
}

func addChildCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		configCmd.NewCmdConfig(),
		extensionCmd.NewCmdExtension(),
	)
}
