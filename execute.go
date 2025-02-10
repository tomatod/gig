package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func Execute(cfg *Config) error {
	if cfg.GetMode() == "list" {
		return listGitignoreFiles(cfg)
	}

	if err := os.MkdirAll(cfg.GetCurrentDirPathOfTarget(), 0775); err != nil {
		return err
	}

	if err := makeGitignoreFile(cfg); err != nil {
		return err
	}

	cmdLine := cfg.GetEditCommandLine()
	cmd := exec.Command(cmdLine[0], cmdLine[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func listGitignoreFiles(cfg *Config) error {
	paths := cfg.GetAllFilesForGitignore()
	for _, scope := range []string{"global", "local", "file"} {
		if path, ok := paths[scope]; ok {
			if file, err := os.Open(path); err != nil {
				continue
			} else {
				if bytes, err := io.ReadAll(file); err != nil {
					return fmt.Errorf("failed to read %s. %v", path, err)
				} else {
					fmt.Printf("--- %s: %s \n", scope, path)
					fmt.Printf("%s\n", bytes)
				}
			}
		}
	}
	return nil
}

func makeGitignoreFile(cfg *Config) error {
	if !cfg.WillCreateGitignoreFromTemplate() {
		if file, err := os.OpenFile(cfg.GetTargetPath(), os.O_RDONLY|os.O_CREATE, 0644); err != nil {
			return err
		} else {
			file.Close()
			return nil
		}
	}

	// get .gitignore file path
	gitroot, err := GetCurrentGitRoot()
	if err != nil {
		return fmt.Errorf("failed to create a new .gitignore file. %v", err)
	}
	gitignoreFilePath := filepath.Join(gitroot, ".gitignore")
	if _, err := os.Stat(gitignoreFilePath); err == nil {
		return fmt.Errorf("there is already  %s.", gitignoreFilePath)
	}

	// get the template from github/gitignore
	configroot, err := GetConfigDir()
	if err != nil {
		return fmt.Errorf("failed to create a new .gitignore file. %v", err)
	}
	githubGitignoreBasePath, err := cloneGitHubGitignore(configroot)
	if err != nil {
		return err
	}
	templateFilePath := filepath.Join(githubGitignoreBasePath, cfg.templatePath)

	// copy the template to .gitignore
	templateFile, err := os.Open(templateFilePath)
	if err != nil {
		return err
	}
	defer templateFile.Close()
	gitignoreFile, err := os.Create(gitignoreFilePath)
	if err != nil {
		return err
	}
	defer gitignoreFile.Close()
	_, err = io.Copy(gitignoreFile, templateFile)
	return err
}

func cloneGitHubGitignore(base string) (string, error) {
	if err := os.MkdirAll(base, 0775); err != nil {
		return "", err
	}

	dst := filepath.Join(base, "gitignore")
	if _, err := os.Stat(dst); err == nil {
		return dst, nil
	}

	fmt.Println(dst)

	cmd := exec.Command("git", "clone", "https://github.com/github/gitignore", dst)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return dst, cmd.Run()
}
