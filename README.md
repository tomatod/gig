# gig 
gig is a simple CLI tool to edit the files that is used by Git to ignore files from its tracking (.gitignore and so on). You can edit the 'gitignore' files for several scope (.gitignore, local repository or global) by gig. If you find any bugs or feature requests, please make GitHub issues.

## Install
### Go install
```shell
go install github.com/tomatod/gig@latest
```

### Execution file download
Please access release page (that is in the right side bar), download binary files and rename it to gig if you need.

## Usage
You can show gig's help by ```gig -h```.
```
NAME:
   gig - Edit the files that is used by Git to ignore files from its tracking (.gitignore and so on)

USAGE:
   gig [global options]

GLOBAL OPTIONS:
   --scope value, -s value     Scope: file (.gitignore), local (.git/info/exclude) and global ($HOME/.config/git/ignore) (default: "file")
   --mode value, -m value      Mode: edit or list. In edit mode, you can edit 'gitignore file'. In list mode, All ignored files in the current repository are listed. (default: "edit")
   --editor value, -e value    A text editor used in edit mode. You can include options of the editor. By default, 'core.editor of git' > '$EDITOR' > 'vi' will be selected in order. (default: "vi")
   --template value, -t value  A template path from https://github.com/github/gitignore as the origin of a new .gitignore. Please also see template section in README (https://github.com/tomatod/gig)
   --help, -h                  show help
```

## Example
```shell
# list all files ignored by Git.
gig -m list

# edit .gitignore
gig

# edit .git/info/exclude for an local repository
gig -s local

# edit $HOME/.config/git/ignore for the global configuration by nano editor
gig -s global -e nano

# create a new .gitignore from templates of github.com/github/gitignore
gig -t community/Golang/Go.AllowList.gitignore
```

## Template
You can create a new .gitignore files from templates. gig fetchs templates from [github.com/github/gitignore repository](https://github.com/github/gitignore). Please follow the next step.

1. Find template you want to copy from [github.com/github/gitignore repository](https://github.com/github/gitignore)
2. Open the template file on GitHub and copy the path. You can copy the path by the copy icon under the GitHub webpage's header.
3. Run ```gig -t <template path>```
   - ex.1) [community/Golang/Go.AllowList.gitignore](https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore)   
     ```gig -t community/Golang/Go.AllowList.gitignore```
   - ex.2) [community/JavaScript/Vue.gitignore](https://github.com/github/gitignore/blob/main/community/JavaScript/Vue.gitignore)   
      ```gig -t community/JavaScript/Vue.gitignore```

If you want to update templates of github.com/github/gitignore on local, please execute the next step.
```
cd <your home>/.config/gig/gitignore
git pull origin main
```
