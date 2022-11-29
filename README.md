# Dotflow

Dotflow is a tool to manage your dotfiles.

1. Link a file/folder to dotflow, this will move your file/folder to ~/.config/dotflow/ and create a symlink to the old path

```shell
dotflow link -p PATH
```

2. Install will recreate the symlink to the original path again, this should be used when you want to use your dotfiles in a new system

```shell
dotflow install
```

3. Remove will move the file back to its original path, remove the rymlink and remove from dotflow control

```shell
dotflow remove -p PATH
```

## How to install

### Build by yourself

1. Install go: https://go.dev/doc/install
2. Install the binary

```shell
go install
```

3. Add go path to your env (.bash, .zshrc)

```shell
export PATH=$PATH:~/go/bin/
```
