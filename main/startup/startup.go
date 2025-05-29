package startup

import (
	appService "moss/application/service"
	"moss/infrastructure/general/command"
	"moss/plugins"
	"os"

	"github.com/gookit/color"
)

func init() {
	executeCommand()
	initPlugins()
}

func executeCommand() {
	if command.AdminPath != "" {
		if err := appService.AdminPathUpdate(command.AdminPath); err != nil {
			panic(err)
		}
		color.Green.Println("admin path updated successfully\n")
	}
	if command.AdminUsername != "" {
		if err := appService.AdminUsernameUpdate(command.AdminUsername); err != nil {
			panic(err)
		}
		color.Green.Println("admin username updated successfully\n")
	}
	if command.AdminPassword != "" {
		if err := appService.AdminPasswordUpdate(command.AdminPassword); err != nil {
			panic(err)
		}
		color.Green.Println("admin password updated successfully\n")
	}
	if command.AdminPath != "" || command.AdminUsername != "" || command.AdminPassword != "" {
		os.Exit(0)
	}
}

func initPlugins() {
	appService.PluginInit(
		plugins.NewGenerateSlug(),
		plugins.NewArticleSanitizer(),
		plugins.NewSaveArticleImages(),
		plugins.NewDetectLinks(),
		plugins.NewGenerateDescription(),
		plugins.NewPreBuildArticleCache(),
		plugins.NewPushToBaidu(),
		plugins.NewPushToBing(),
		plugins.NewMakeCarousel(),
		plugins.NewPostStore(),
		plugins.NewDidiAuto(),
	)

}
