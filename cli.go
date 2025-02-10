package main

import (
	"context"
	"errors"
	"github.com/urfave/cli/v3"
	"golang.org/x/exp/slices"
	"os"
	"os/exec"
	"strings"
)

func GetCLI(ctx context.Context, action cli.ActionFunc) *cli.Command {
	scopeFlag := &cli.StringFlag{
		Name:    "scope",
		Aliases: []string{"s"},
		Usage:   "A scopes: file (.gitignore), local (.git/info/exclude) and global ($HOME/.config/git/ignore)",
		Value:   "file",
		Validator: func(s string) error {
			if !slices.Contains([]string{"file", "local", "global"}, s) {
				return errors.New("must select file, local or global as scope.")
			}
			return nil
		},
	}

	modeFlag := &cli.StringFlag{
		Name:    "mode",
		Aliases: []string{"m"},
		Usage:   "A mode: edit or list. In edit mode, you can edit 'gitignore file'. In list mode, All ignored files in the current repository are listed.",
		Value:   "edit",
		Validator: func(m string) error {
			if !slices.Contains([]string{"edit", "list"}, m) {
				return errors.New("must specify edit or list as mode.")
			}
			return nil
		},
	}

	editorFlag := &cli.StringFlag{
		Name:    "editor",
		Aliases: []string{"e"},
		Usage:   "A text editor used in edit mode. You can include options of the editor. By default, 'core.editor of git' > '$EDITOR' > 'vi' will be selected in order.",
		Value:   findDefaultEditorMust(),
	}

	templatePathFlag := &cli.StringFlag{
		Name:    "template",
		Aliases: []string{"t"},
		Usage:   "A template path from https://github.com/github/gitignore as the origin of a new .gitignore. Please also see template section in README (https://github.com/tomatod/gig)",
		Value:   "",
		Validator: func(t string) error {
			if strings.Trim(t, " \r\n") == "" {
				return errors.New("cannot specify empty string.")
			}
			return nil
		},
	}

	return &cli.Command{
		Name:                  "gig",
		Usage:                 "Edit the files that is used by Git to ignore files from its tracking (.gitignore and so on)",
		EnableShellCompletion: true,
		Flags: []cli.Flag{
			scopeFlag,
			modeFlag,
			editorFlag,
			templatePathFlag,
		},
		Action: action,
		After: func(_ context.Context, cli *cli.Command) error {
			if cli.String("mode") == "edit" && strings.Trim(cli.String("editor"), " \r\n") == "" {
				errors.New("must specify an editor to edit 'gitignore file'")
			}
			if cli.String("template") != "" && cli.String("scope") != "file" {
				return errors.New("must select a file scope when you make a new .gitignore file from template.")
			}
			return nil
		},
	}
}

func findDefaultEditor() (string, error) {
	// refer "git config core.editor"
	cmd := exec.Command("git", "config", "core.editor")
	out := strings.Builder{}
	cmd.Stdout = &out
	cmd.Run()
	result := strings.Trim(out.String(), " \n")
	if result != "" {
		return result, nil
	}

	// refer $EDITOR environment valuable
	result = os.Getenv("EDITOR")
	result = strings.Trim(out.String(), " ")
	if result != "" {
		return result, nil
	}

	// check whether there is vi a text editor.
	if _, err := exec.LookPath("vi"); err == nil {
		return "vi", nil
	}

	return "", errors.New("cannot find any appriciate text editor.")
}

func findDefaultEditorMust() string {
	editor, err := findDefaultEditor()
	if err != nil {
		panic(err)
	}
	return editor
}
