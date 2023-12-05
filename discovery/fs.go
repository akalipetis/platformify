package discovery

import (
	"io/fs"
	"path/filepath"
	"slices"
)

var skipDirs = []string{
	"vendor",
	"node_modules",
	".next",
	".git",
}

func (d *Discoverer) CountFiles() (map[string]int, error) {
	fileCounter := make(map[string]int)
	err := fs.WalkDir(d.fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			if slices.Contains(skipDirs, d.Name()) {
				return filepath.SkipDir
			}

			return nil
		}

		ext := filepath.Ext(path)
		_, ok := fileCounter[ext]
		if !ok {
			fileCounter[ext] = 0
		}

		fileCounter[ext]++
		return nil
	})
	if err != nil {
		return nil, err
	}

	return fileCounter, nil
}
