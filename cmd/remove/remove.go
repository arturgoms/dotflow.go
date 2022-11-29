/*
Copyright © 2022 Artur Gomes <contact@arturgomes.com>
*/
package remove

import (
	"dotflow/cmd/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
  "strings"
	"github.com/spf13/cobra"
)

var (
  path string
)

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the file or folder from dotflow",
	Long: `Remove command will need a path from a folder or file from inside 
  the dotflow folder and it will remove it from dotfiles

  Example:
    dotflow remove -p ~/.config/dotflow/nvim`,
	Run: func(cmd *cobra.Command, args []string) {
    path, _ := filepath.Abs(path)
    if og_path, err := utils.RemoveFromConfig(path); err == true {
      log.Fatal("Could not write in dotflow file.")
    } else {
      dirname, err := os.UserHomeDir()
      if err != nil {
        log.Fatal( err )
      }
      records := strings.Split(og_path, ":")
      config_folder := filepath.Join(dirname,".config/dotflow/")
      from := strings.Replace(records[0], "$HOME", dirname, 1)
      to := strings.Replace(filepath.Join(config_folder, filepath.Base(path)), "$HOME", dirname, 1)
      os.Rename(to, from) 
      println("Removed: ", to) 

    }
	},
}

func init() {
  RemoveCmd.Flags().StringVarP(&path, "path", "p", "", "Path for file or folder.")
  if err := RemoveCmd.MarkFlagRequired("path"); err != nil {
    fmt.Print(err)
  }
}
