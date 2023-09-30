package cleanup

import (
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/carldanley/frigate-clip-cleaner/pkg/filesystem"
	"github.com/sirupsen/logrus"
)

const (
	ByteDivisor  = 1024
	SecondsInDay = 86400
)

func Execute(cfg *Config, log *logrus.Logger) error {
	log.Debugf("checking if frigate config path exists: %s", cfg.FrigateConfigPath)

	if !filesystem.PathExists(cfg.FrigateConfigPath) {
		return ErrorFrigateConfigDoesNotExist
	}

	log.Debug("parsing frigate configuration")

	configData, err := ParseFrigateConfig(cfg.FrigateConfigPath)
	if err != nil {
		return err
	}

	log.Debug("getting max retention from config data")

	maxDays := getMaxRetentionFromConfigData(configData)

	log.Infof("Parsed a maximum clip retention period of %d day(s)", maxDays)

	log.Debugf("checking if frigate clip directory exists: %s", cfg.FrigateClipDirectory)

	if !filesystem.PathExists(cfg.FrigateClipDirectory) {
		return ErrorFrigateClipDirDoesNotExist
	}

	log.Debug("scanning for frigate clip directory for relevant file types")

	fileTypes := []string{".jpeg", ".jpg", ".png"}
	relevantFiles, err := filesystem.ScanForFiles(cfg.FrigateClipDirectory, fileTypes)

	if err != nil {
		return fmt.Errorf("could not scan for files: %w", err)
	}

	log.Infof("Found %d relevant clip(s)", len(relevantFiles))
	log.Debug("iterating over relevant clips found")

	var mbToSave float64

	maxRetentionDay := time.Now().Add(-(time.Second * time.Duration(SecondsInDay*(maxDays+1))))
	deletionQueue := []fs.DirEntry{}

	log.Debugf("Maximum Retention Day: %v", maxRetentionDay)

	for _, file := range relevantFiles {
		fileInfo, err := file.Info()
		if err != nil {
			return fmt.Errorf("could not get file info: %w", err)
		}

		if fileInfo.ModTime().After(maxRetentionDay) {
			deletionQueue = append(deletionQueue, file)
			mbToSave += (float64(fileInfo.Size()/ByteDivisor) / ByteDivisor)
		} else {
			log.Debugf("Skipping %s: mod time: %v", file.Name(), fileInfo.ModTime())
		}
	}

	log.Infof("Found %d clip(s) older than the %d clip retention day period (+1 additional day)", len(deletionQueue), maxDays)
	log.Infof("Reclaimable space: %.2f MB", mbToSave)

	if cfg.DryRun || (len(deletionQueue) == 0) {
		return nil
	}

	return performDeletion(cfg.FrigateClipDirectory, deletionQueue, log)
}

func performDeletion(clipDirectoryPath string, queue []fs.DirEntry, log *logrus.Logger) error {
	log.Info("Hang tight! Performing the deletion...")
	log.Info("=========================================")

	var spaceReclaimedMB float64

	filesDeleted := 0

	for _, file := range queue {
		fileInfo, err := file.Info()
		if err != nil {
			log.WithError(err).Error("could not get file info")

			continue
		}

		pathToRemove := fmt.Sprintf("%s%s", clipDirectoryPath, file.Name())
		fileSizeMB := (float64(fileInfo.Size()/ByteDivisor) / ByteDivisor)

		log.Infof("Deleting: %s: %.2f MB", pathToRemove, fileSizeMB)

		if err := os.Remove(pathToRemove); err != nil {
			log.WithError(err).Error("could not delete file")

			continue
		}

		spaceReclaimedMB += fileSizeMB
		filesDeleted++
	}

	log.Info("=========================================")
	log.Infof("Deleted %d file(s)! Total Reclaimed Space: %.2f MB", filesDeleted, spaceReclaimedMB)

	return nil
}
