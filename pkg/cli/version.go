package cli

import (
	"github.com/fabiorphp/kongo-cli/pkg/template"
	"github.com/urfave/cli"
	"runtime"
)

var (
	versionTmpl = `Version: {{ .Version }}
GO Version: {{ .GoVersion }}
OS/Arch: {{ .Os }}/{{ .Arch }}
`
)

type (
	// VersionInfo structure
	VersionInfo struct {
		// Cli version
		Version string

		// Golang version
		GoVersion string

		// OS name
		Os string

		// Architecture name
		Arch string
	}
)

// Version retrives a command responsible for showing version information
func Version() cli.Command {
	return cli.Command{
		Name:  "version",
		Usage: "Shows the Kongo version information",
		Action: func(c *cli.Context) error {
			tmpl, err := template.NewPlain(versionTmpl)

			if err != nil {
				return err
			}

			info := VersionInfo{
				Version:   c.App.Version,
				GoVersion: runtime.Version(),
				Os:        runtime.GOOS,
				Arch:      runtime.GOARCH,
			}

			if err := tmpl.Write(info); err != nil {
				return err
			}

			return tmpl.Flush(c.App.Writer)
		},
	}
}
