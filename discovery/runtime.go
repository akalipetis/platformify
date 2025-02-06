package discovery

import (
	"github.com/platformsh/platformify/internal/utils"
	"github.com/platformsh/platformify/platformifier"
)

var (
	languageMap = map[string]string{
		".py":   "python",
		".js":   "nodejs",
		".go":   "golang",
		".php":  "php",
		".rb":   "ruby",
		".exs":  "elixir",
		".ex":   "elixir",
		".cs":   "dotnet",
		".rs":   "rust",
		".lisp": "lisp",
		".lsp":  "lisp",
		".l":    "lisp",
		".cl":   "lisp",
		".fasl": "lisp",
		".java": "java",
	}
)

// Returns the Runtime, either from memory or by discovering it on the spot
func (d *Discoverer) Type() (string, error) {
	if stack, ok := d.memory["type"]; ok {
		return stack.(string), nil
	}

	typ, err := d.discoverType()
	if err != nil {
		return "", err
	}

	d.memory["type"] = typ
	return typ, nil
}

func (d *Discoverer) discoverType() (string, error) {
	stack, err := d.Stack()
	if err != nil {
		return "", err
	}

	switch stack {
	case platformifier.Laravel, platformifier.Symfony:
		return "php", nil
	case platformifier.Django, platformifier.Flask:
		return "python", nil
	case platformifier.Express, platformifier.NextJS, platformifier.Strapi:
		return "nodejs", nil
	}

	extCount, err := utils.CountFiles(d.fileSystem)
	if err != nil {
		return "", err
	}

	langCount := make(map[string]int)
	for ext, count := range extCount {
		if lang, ok := languageMap[ext]; ok {
			if _, _ok := langCount[lang]; !_ok {
				langCount[lang] = 0
			}
			langCount[lang] += count
		}
	}

	max := 0
	selectedLang := ""
	for lang, count := range langCount {
		if count > max {
			max = count
			selectedLang = lang
		}
	}

	return selectedLang, nil
}
