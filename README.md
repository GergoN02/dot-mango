# Dot-Mango
## A dotfiles version manager, right from within your configs.

### But why another dots manager???
I ran into some inconveniences switching between dotfiles in different environments.
My config needs differ at work from that of personal projects, and I like to keep
backups of my older configs in the same dotfiles repo on GitHub.

I got tired of using a 'Makefile' to symlink my configs into their directories whenever
I wanted to switch them, and adding confusing "wrapper" scripts to execute multiple
symlink operations depending on which parts of my config I wanted to switch out.

I made this small project (in 2 weekends, I know it's messy), to remedy this situation.

The entire project is written in `go` with as much reliance on the go Standard Lib as possible.
The main script is intended to live in your dotfiles repo, as to avoid as much configuration as needed,
but features for "directory-awareness", similarly to [LazyGit](https://github.com/jesseduffield/lazygit)
will be added in the future.

**This project is very early-days and PRs/issues are welcome.
This project is also part of my Go Learning journey, so any feedback or improvements are appreciated :))**


### Installation

1. Clone the repo with `git clone --depth=1 https://github.com/thegenem0/dot-mango.git`
2. Run `go build -o mango ./cmd/dot-mango/main.go`, and move the resulting executable to your dotfiles repo.
3. Call `./mango --init` to initialize a git repository if not already in git, and create a `mangoConfig.yaml` file.
4. Configure your Repo structure along with your `mangoConfig` and just run `./mango`


### Usage

When running `./mango`, your dotfile paths will get loaded from the `mangoConfig`.
- Use `Tab` to switch active panels (certain actions are bound to specific panels)
- Navigate lists with `j` and `k`
- Select active config from the left-hand side list with `Enter`
- Toggle selection on an individual dotfile/folder with `Space`
    `a` toggles all config files/folders in the selected view
- `s` will show a Popup with info about the symlink process, and hitting `y` will symlink
your dotfiles to the appropriate target directories.


### Config

An example config can be found in [mangoConfig.example.yaml](mangoConfig.example.yaml)

Different top-level folders, can be configured under the `dotfile_folders` entry in the following shape:
```
name: Name_in_configs_list
path: /relative/path/within/your/dotfiles
```
Additionally, overrides can be provided for specific files, that shouldn't get symlinked to `$XDG_CONFIG_HOME`.
Usually these files would get linked to `$HOME`, therefore the override paths can be provided relative to `$XDG_CONFIG_HOME` too.

```
config: ConfigName/all (all will apply the override to every config declared in `dotfile_folders`)
dotfile_path: .vimrc (no need to provide the config name from `dotfile_folders`, will apply relatively)
override_target: ../ (can be relative to `$XDG_CONFIG_HOME`, or absolute path)
```

### Commands

When setting up a brand new repo (which is advised, even for just experimentation), the following commands might be useful

`./mango --init`
Initializes a git repository within your current directory and creates a default mangoConfig.yaml file

`./mango --generate`
Generates empty folder structure from the defined `mangoConfig.yaml` dotfile dirs

`./mango --path /example/path`
Runs mango in the provided path as the `mangoConfig.yaml` path. Not yet useful for much.

`./mango --help`
Prints help menu
