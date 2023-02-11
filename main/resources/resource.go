package resources

import (
	"embed"
	"io/fs"
)

var (
	//go:embed app
	App embed.FS
	//go:embed admin
	admin embed.FS

	//go:embed themes
	Themes        embed.FS
	ThemesDirName = "themes"
)

func Admin() fs.FS {
	s, _ := fs.Sub(admin, "admin")
	return s
}
