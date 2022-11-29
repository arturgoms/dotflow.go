/*
Copyright © 2022 Artur Gomes <contact@arturgomes.com>

*/
package cmd

import (
	"os"
  "dotflow/cmd/link"
  "dotflow/cmd/remove"
  "dotflow/cmd/install"
	"github.com/spf13/cobra"
  "errors"
  "log"
  "path/filepath"
)



func initDotflow() {
  dirname, err := os.UserHomeDir()
  if err != nil {
    log.Fatal( err )
  }
  config_folder := filepath.Join(dirname,".config/dotflow/")
  
  if _, err := os.Stat(config_folder); errors.Is(err, os.ErrNotExist) {
    err := os.MkdirAll(config_folder, os.ModePerm)
    if err != nil {
      log.Println(err)
    }
    f, e := os.Create(filepath.Join(config_folder, ".dotflow"))
    if e != nil {
      panic(e)
    }
    defer f.Close()
  } else {

    if _, err := os.Stat(filepath.Join(config_folder, ".dotflow")); errors.Is(err, os.ErrNotExist) {
      f, e := os.Create(filepath.Join(config_folder, ".dotflow"))
      if e != nil {
        panic(e)
      }
      defer f.Close()
    } 
  }
}

var rootCmd = &cobra.Command{
	Use:   "dotflow",
	Short: "A simple manager for your dotfiles",
	Long: `Dotflow is a simple CLI that put all your dotfiles in one single 
  place and makes your life easier when you want to install then again. 
         
  Basically you just need to link your file/folder with dotflow and dotflow 
  will move the files to the dotflow folder and creates a symbolic link to the old path, 
  then you can use the dotflow folder as a git repository.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
  initDotflow()
  
  // Add link command
  rootCmd.AddCommand(link.LinkCmd)

  // Add install command
  rootCmd.AddCommand(install.InstallCmd)

  // Add remove command
  rootCmd.AddCommand(remove.RemoveCmd)
}


