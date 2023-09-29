package commands

var FccVersion = "unstable" //nolint: gochecknoglobals

const (
	FccUse = `fcc -c File`

	FccShort = `frigate-clip-cleaner (fcc) is a clip cleanup tool for the
Frigate clip directory.`

	FccLong = `frigate-clip-cleaner (fcc) is a tool written to address
orphaned clips in the Frigate clip directory (which can lead to tons of
unused storage).`

	FccHelpUsageExample = `fcc --frigate-config /config/config.yml --dry-run`

	FccHelpLogLevel = `(env: FCC_LOG_LEVEL) Log-level to display to stdout, one of:
		ERROR:   {"error", "err", "e"}
		INFO:    {"info", "i"}
		WARNING: {"warning", "warn", "w"}
		DEBUG:   {"debug", "d", "verbose", "v"}
`
)
