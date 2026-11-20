package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"charm.land/huh/v2"
	"github.com/avtorgenii/cursor-opener/internal/launcher"
	"github.com/avtorgenii/cursor-opener/internal/projects"
	"github.com/spf13/cobra"
)

// dir живёт на уровне пакета, чтобы list и open видели одно и то же значение флага.
var dir string

var rootCmd = &cobra.Command{
	Use:          "co",
	Short:        "Open projects from your folder in Cursor",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		names, err := projects.List(dir)
		if err != nil {
			return err
		}

		var choices []string

		err = huh.NewMultiSelect[string]().
			Title("Which projects to open?").
			Options(huh.NewOptions(names...)...).
			Filterable(true).
			Value(&choices).
			Run()

		if errors.Is(err, huh.ErrUserAborted) {
			return nil // пользователь Ctrl + C, это не ошибка
		}

		if err != nil {
			return err
		}

		for _, name := range choices {
			err := launcher.Open(filepath.Join(dir, name))

			if err != nil {
				return err
			}
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// Особая функция. Go сам вызывает её при старте пакета, до main.
// Здесь удобно регистрировать флаги и команды.
func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot find home dir: ", err)
		os.Exit(1)
	}

	rootCmd.Flags().StringVarP(
		&dir, "dir", "d", filepath.Join(home, "dev"), "root folder with projects",
	)
}
