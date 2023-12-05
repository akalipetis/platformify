package discovery

import (
	"io/fs"
)

const (
	settingsPyFile   = "settings.py"
	managePyFile     = "manage.py"
	composerJSONFile = "composer.json"
	packageJSONFile  = "package.json"
	symfonyLockFile  = "symfony.lock"
)

type Discoverer struct {
	fileSystem fs.FS
	memory     map[string]any
}

func New(fileSystem fs.FS) *Discoverer {
	return &Discoverer{fileSystem: fileSystem, memory: make(map[string]any)}
}
