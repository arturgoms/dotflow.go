/*
Copyright © 2022 Artur Gomes <contact@arturgomes.com>

*/
package install

import (
  "os"
  "log"
  "path/filepath"
  "io/ioutil"
  "strings"

	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install will create again the symlink to the original path",
	Long: `Install will recreate the symbolic link to all files
  that was controlled by dotflow`,
	Run: func(cmd *cobra.Command, args []string) {
    dirname, err := os.UserHomeDir()
    if err != nil {
      log.Fatal( err )
    }
    config_folder := filepath.Join(dirname,".config/dotflow/")
    config := filepath.Join(config_folder, ".dotflow")
    fileBytes, err := ioutil.ReadFile(config)
 	  lines := strings.Split(string(fileBytes), "\n")
    for _, line := range lines {
      if line != "" {
        records := strings.Split(line, ":") 
        from := strings.Replace(records[0], "$HOME", dirname, 1)
        to := strings.Replace(records[1], "$HOME", dirname, 1)
        os.Symlink(to, from)
      }
    }
	},
}

func init() {}
