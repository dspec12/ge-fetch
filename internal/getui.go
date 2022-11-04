package internal

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

const (
	urlWineGECustom   = "https://api.github.com/repos/GloriousEggroll/wine-ge-custom/releases"
	urlProtonGECustom = "https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases"
	wineInstallDir    = ".local/share/lutris/runners/wine"
	protonInstallDir  = ".steam/root/compatibilitytools.d"
)

func GETUI() {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   fmt.Sprintf(`{{ "%s" | green }} {{ . | faint }}`, "▸"),
		Inactive: "  {{ . }}",
		Selected: fmt.Sprintf(`{{ "%s" | green }} {{ . | faint }}`, "✔"),
	}

	prompt := promptui.Select{
		Label:     "Select GE Type",
		Items:     []string{"Wine-GE", "Proton-GE"},
		Templates: templates,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fetch(result)
}

func fetch(geType string) {
	var r []release
	var installDir string

	switch geType {
	case "Wine-GE":
		r = getReleases(urlWineGECustom)
		installDir = wineInstallDir
	case "Proton-GE":
		r = getReleases(urlProtonGECustom)
		installDir = protonInstallDir
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   fmt.Sprintf(`{{ "%s" | green }} {{ .TagName | faint }}`, "▸"),
		Inactive: "  {{ .TagName }}",
		Selected: fmt.Sprintf(`{{ "%s" | green }} {{ .TagName | faint }}`, "✔"),
		Details: `
------------ Info ------------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Published:" | faint }}	{{ .Published }}
{{ "Release Page:" | faint }}	{{ .HTMLURL }}`,
	}

	prompt := promptui.Select{
		Label:     "Select Release",
		Items:     r,
		Templates: templates,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	r[i].download(installDir)

}
