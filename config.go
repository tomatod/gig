package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	scope        string
	mode         string
	editor       []string
	basePath     string
	templatePath string
}

var (
	scopeTargets = map[string]string{
		"file":   ".gitignore",
		"local":  ".git/info/exclude",
		"global": ".config/git/ignore",
	}
	modes = map[string]bool{
		"edit": true,
	}
)

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetConfigFromCLI(_ context.Context, cmd *cli.Command) error {
	c.scope = cmd.String("scope")
	c.mode = cmd.String("mode")
	c.editor = strings.Split(cmd.String("editor"), " ")
	c.templatePath = cmd.String("template")
	return c.setbasePath()
}

func (c *Config) setbasePath() error {
	var err error
	if c.scope == "file" || c.scope == "local" {
		c.basePath, err = GetCurrentGitRoot()
		if err != nil {
			return err
		}
	} else {
		c.basePath, err = os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home dir path. %v", err)
		}
	}

	return nil
}

func (c *Config) GetEditCommandLine() []string {
	return append(c.editor, c.GetTargetPath())
}

func (c *Config) GetCurrentDirPathOfTarget() string {
	dir, _ := filepath.Split(c.GetTargetPath())
	return dir
}

func (c *Config) GetTargetPath() string {
	return filepath.Join(c.basePath, scopeTargets[c.scope])
}

func (c *Config) GetAllFilesForGitignore() map[string]string {
	result := map[string]string{}
	if base, err := GetCurrentGitRoot(); err == nil {
		result["file"] = filepath.Join(base, scopeTargets["file"])
		result["local"] = filepath.Join(base, scopeTargets["local"])
	}
	if base, err := os.UserHomeDir(); err == nil {
		result["global"] = filepath.Join(base, scopeTargets["global"])
	}
	return result
}

func (c *Config) GetMode() string {
	return c.mode
}

func (c *Config) GetScope() string {
	return c.scope
}

func (c *Config) WillCreateGitignoreFromTemplate() bool {
	if c.GetScope() != "file" {
		return false
	}
	if c.templatePath == "" {
		return false
	}
	return true
}

func (c *Config) GetTemplatePath() (string, error) {
	base, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "gitignore", c.templatePath), nil
}
