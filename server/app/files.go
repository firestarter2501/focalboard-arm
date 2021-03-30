package app

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/mattermost/focalboard/server/utils"
)

func (a *App) SaveFile(reader io.Reader, workspaceID, rootID, filename string) (string, error) {
	// NOTE: File extension includes the dot
	fileExtension := strings.ToLower(filepath.Ext(filename))
	if fileExtension == ".jpeg" {
		fileExtension = ".jpg"
	}

	createdFilename := fmt.Sprintf(`%s%s`, utils.CreateGUID(), fileExtension)
	filePath := filepath.Join(workspaceID, rootID, createdFilename)

	_, appErr := a.filesBackend.WriteFile(reader, filePath)
	if appErr != nil {
		return "", errors.New("unable to store the file in the files storage")
	}

	return createdFilename, nil
}

func (a *App) GetFilePath(workspaceID, rootID, filename string) string {
	folderPath := a.config.FilesPath
	rootPath := filepath.Join(folderPath, workspaceID, rootID)

	return filepath.Join(rootPath, filename)
}
