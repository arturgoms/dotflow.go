/*
Copyright © 2022 Artur Gomes <contact@arturgomes.com>
*/
package link

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

var LinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link your dot file or folder to Dotflow",
	Long: `Link needs a path for a folder or file, and it will move that 
  file/folder to the dotflow folder and creates a symbolic link to its old path.

  Example:  
    dotflow link -p ~/.config/nvim
  `,
	Run: func(cmd *cobra.Command, args []string) {
    path, _ := filepath.Abs(path)
    if paths, err := utils.AppendToConfig(path); err == true {
      log.Fatal("Could not write in dotflow file.")
    } else {
      records := strings.Split(paths, ":")
      dirname, err := os.UserHomeDir()
      if err != nil {
        log.Fatal( err )
      }
      from := strings.Replace(records[0], "$HOME", dirname, 1)
      to := strings.Trim(strings.Replace(records[1], "$HOME", dirname, 1), "\n")
      os.Rename(from, to)
      os.Symlink(to, from)
      println("Linked to ", from)
    }
	},
}
func init() {
  LinkCmd.Flags().StringVarP(&path, "path", "p", "", "Path for file or folder.")
  if err := LinkCmd.MarkFlagRequired("path"); err != nil {
    fmt.Print(err)
  }
}
