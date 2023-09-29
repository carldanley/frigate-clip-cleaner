package cleanup

import "github.com/sirupsen/logrus"

type Config struct {
	FrigateConfigPath    string
	FrigateClipDirectory string
	LogLevel             logrus.Level
	DryRun               bool
}
