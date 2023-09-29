package commands

import (
	"fmt"
	"os"

	"github.com/carldanley/frigate-clip-cleaner/internal/cleanup"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Root struct {
	Log     *logrus.Logger
	Options *cleanup.Config
	Command *cobra.Command
}

func New() *Root {
	rc := &Root{
		Log: logrus.New(),
		Options: &cleanup.Config{
			LogLevel:          logrus.DebugLevel,
			DryRun:            (os.Getenv("FCC_DRY_RUN") == "true") || (os.Getenv("FCC_DRY_RUN") == "1"),
			FrigateConfigPath: os.Getenv("FCC_CONFIG_FILE"),
		},
	}

	rc.Command = rc.NewCommand()
	rc.AddFlags()
	rc.AddCommands()

	return rc
}

func (r Root) NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:                        FccUse,
		Aliases:                    []string{},
		SuggestFor:                 []string{},
		Short:                      FccShort,
		Long:                       FccLong,
		Example:                    FccHelpUsageExample,
		ValidArgs:                  []string{},
		Args:                       nil,
		ArgAliases:                 []string{},
		BashCompletionFunction:     "",
		Deprecated:                 "",
		Hidden:                     false,
		Annotations:                map[string]string{},
		Version:                    FccVersion,
		PersistentPreRun:           r.SetupLogging,
		PersistentPreRunE:          nil,
		PreRun:                     nil,
		PreRunE:                    nil,
		RunE:                       r.Execute,
		PostRun:                    nil,
		PostRunE:                   nil,
		PersistentPostRun:          nil,
		PersistentPostRunE:         nil,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          true,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
		TraverseChildren:           false,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	}

	return rootCmd
}

func (r *Root) SetupLogging(cmd *cobra.Command, args []string) {
	r.Log.SetLevel(r.Options.LogLevel)
}

func (r *Root) AddFlags() {
	r.initializeCommonFlags()
	r.initializeCleanupFlags()
}

func (r *Root) AddCommands() {
	// this is a placeholder for any new commands that may be added later
	// r.Command.AddCommand(r.NewCommandHere)
}

func (r *Root) Execute(cmd *cobra.Command, args []string) error {
	if err := cleanup.Execute(r.Options, r.Log); err != nil {
		cmd.SilenceUsage = true

		return fmt.Errorf("%w", err)
	}

	return nil
}

func (r *Root) Run() {
	if err := r.Command.Execute(); err != nil {
		r.Log.Error(err)
		os.Exit(1)
	}
}
