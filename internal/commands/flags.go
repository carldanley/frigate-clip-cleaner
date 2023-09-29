package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/thediveo/enumflag"
)

func (r *Root) initializeCommonFlags() {
	logMap := map[logrus.Level][]string{
		logrus.ErrorLevel: {"error", "err", "e"},
		logrus.InfoLevel:  {"info", "i"},
		logrus.WarnLevel:  {"warning", "warn", "w"},
		logrus.DebugLevel: {"debug", "d", "verbose", "v"},
	}

	r.Command.Flags().VarP(
		enumflag.New(&r.Options.LogLevel, "log-level", logMap, enumflag.EnumCaseInsensitive),
		"log-level",
		"v",
		FccHelpLogLevel,
	)
}

func (r *Root) initializeCleanupFlags() {
	r.Command.Flags().StringVarP(
		&r.Options.FrigateConfigPath,
		"frigate-config", "c",
		r.Options.FrigateConfigPath,
		"(env: FCC_CONFIG_FILE) frigate config file (default is /config/config.yml)",
	)

	r.Command.Flags().StringVarP(
		&r.Options.FrigateClipDirectory,
		"frigate-clip-dir", "x",
		r.Options.FrigateClipDirectory,
		"(env: FCC_CLIP_DIR) frigate clip directory (default is /media/frigate/clips)",
	)

	r.Command.Flags().BoolVarP(
		&r.Options.DryRun,
		"dry-run", "d", false,
		"perform a dry run without deletions",
	)
}
