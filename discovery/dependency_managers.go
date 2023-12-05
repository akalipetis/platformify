package discovery

import (
	"slices"

	"github.com/platformsh/platformify/internal/utils"
)

var (
	dependencyManagersMap = []struct {
		lang     string
		lockFile string
		name     string
	}{
		{lockFile: "yarn.lock", name: "yarn", lang: "nodejs"},
		{lockFile: "package-lock.json", name: "npm", lang: "nodejs"},
		{lockFile: "poetry.lock", name: "poetry", lang: "python"},
		{lockFile: "Pipfile.lock", name: "pipenv", lang: "python"},
		{lockFile: "requirements.txt", name: "pip", lang: "python"},
		{lockFile: "composer.lock", name: "composer", lang: "php"},
	}
)

// Returns the dependency managers, either from memory or by discovering it on the spot
func (d *Discoverer) DependencyManagers() ([]string, error) {
	if dependencyManagers, ok := d.memory["dependency_managers"]; ok {
		return dependencyManagers.([]string), nil
	}

	dependencyManagers, err := d.discoverDependencyManagers()
	if err != nil {
		return nil, err
	}

	d.memory["dependency_managers"] = dependencyManagers
	return dependencyManagers, nil
}

func (d *Discoverer) discoverDependencyManagers() ([]string, error) {
	dependencyManagers := make([]string, 0)
	matchedLanguages := make([]string, 0)
	for _, dependencyManager := range dependencyManagersMap {
		if slices.Contains(matchedLanguages, dependencyManager.lang) {
			continue
		}

		if utils.FileExists(d.fileSystem, "", dependencyManager.lockFile) {
			dependencyManagers = append(dependencyManagers, dependencyManager.name)
			matchedLanguages = append(matchedLanguages, dependencyManager.lang)
		}
	}

	return dependencyManagers, nil
}
