package platformifier

import (
	"io/fs"
	"strings"
)

var (
	databases = []string{
		"mariadb",
		"mysql",
		"oracle-mysql",
		"postgresql",
	}
	caches = []string{
		"redis",
		"redis-persistent",
		"memcached",
	}
)

const (
	Generic Stack = iota
	Django
	Laravel
	NextJS
	Strapi
	Flask
	Express
	Rails
	Symfony
	Ibexa
	Shopware
)

type Stack int

func (s Stack) Name() string {
	switch s {
	case Generic:
		return "generic"
	case Django:
		return "django"
	case Rails:
		return "rails"
	case Laravel:
		return "laravel"
	case NextJS:
		return "nextjs"
	case Strapi:
		return "strapi"
	case Flask:
		return "flask"
	case Express:
		return "express"
	case Symfony:
		return "symfony"
	case Ibexa:
		return "ibexa"
	case Shopware:
		return "shopware"
	default:
		return ""
	}
}

type Relationship struct {
	Service  string
	Endpoint string
}

// UserInput contains the configuration from user input.
type UserInput struct {
	Stack              Stack
	ApplicationRoot    string
	Name               string
	Type               string
	Environment        map[string]string
	BuildSteps         []string
	WebCommand         []string
	DeployCommand      []string
	DependencyManagers []string
	Locations          map[string]map[string]any
	Dependencies       map[string]map[string]string
	Mounts             map[string]map[string]string
	Services           []Service
	Relationships      map[string]Relationship
	HasGit             bool
	WorkingDirectory   fs.FS
}

// Service contains the configuration for a service needed by the application.
type Service struct {
	Name         string
	Type         string
	TypeVersions []string
	Disk         string
	DiskSizes    []string
}

// Database returns the first service that is a database.
func (ui *UserInput) Database() string {
	for _, service := range ui.Services {
		for _, db := range databases {
			if strings.Contains(service.Type, db) {
				return service.Name
			}
		}
	}

	return ""
}

// DatabaseUpper returns the uppercase slug for the first service that is a database.
func (ui *UserInput) DatabaseUpper() string {
	return strings.ToUpper(strings.ReplaceAll(ui.Database(), "-", "_"))
}

// Cache returns the first service that is a cache.
func (ui *UserInput) Cache() string {
	for _, service := range ui.Services {
		for _, cache := range caches {
			if strings.Contains(service.Type, cache) {
				return service.Name
			}
		}
	}

	return ""
}

// CacheUpper returns the uppercase slug for the first service that is a cache.
func (ui *UserInput) CacheUpper() string {
	return strings.ToUpper(strings.ReplaceAll(ui.Cache(), "-", "_"))
}
