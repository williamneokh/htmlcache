package htmlcache

import (
	"errors"
	"html/template"
	"path/filepath"
)

type cache struct {
	PagePath   string
	LayoutPath string
}

func New(pagePath, layoutPath string) *cache {
	return &cache{
		PagePath:   pagePath,
		LayoutPath: layoutPath,
	}
}

func (c cache) CreateTemplateCache() (map[string]*template.Template, error) {
	var (
		cache = map[string]*template.Template{}
	)
	if c.LayoutPath == "" || c.PagePath == "" {
		return cache, errors.New("please provide dir path")
	}
	//find all pages from given dir
	pages, err := filepath.Glob(c.PagePath)
	if err != nil {
		return cache, err
	}

	//find all layout from given dir
	layouts, err := filepath.Glob(c.LayoutPath)
	if err != nil {
		return cache, err
	}

	//range pages and create new template type for each page
	//once each page is created, range layouts and parse into page
	for _, page := range pages {
		name := filepath.Base(page)
		tc, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		if len(layouts) > 0 {
			for _, layout := range layouts {
				tc.ParseFiles(layout)
			}
		}
		cache[name] = tc

	}
	return cache, nil
}
