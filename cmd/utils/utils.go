package utils

import (
	"log"
	"os"
	"path/filepath"
  "fmt"
  "strings"
  "io/ioutil"
)

func AppendToConfig(path string) (string, bool) {
  dirname, err := os.UserHomeDir()
  if err != nil {
    log.Fatal( err )
    return "", true
  }

  config_folder := filepath.Join(dirname,".config/dotflow/")
  config := filepath.Join(config_folder, ".dotflow")
  f, err := os.OpenFile(config, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
      log.Fatal(err)
      return "", true
  }
  from := strings.Replace(path, dirname, "$HOME", 1)
  to := strings.Replace(filepath.Join(config_folder, filepath.Base(path)), dirname, "$HOME", 1)
  record := fmt.Sprintf("%s:%s\n", from , to)
  if _, err := f.Write([]byte(record)); err != nil {
      log.Fatal(err)
      return "", true
  }
  if err := f.Close(); err != nil {
      log.Fatal(err)
      return "", true
  } 
  return record, false
}

func RemoveFromConfig(path string) (string, bool) {
  dirname, err := os.UserHomeDir()
  if err != nil {
    log.Fatal( err )
    return "", true
  }
  config_folder := filepath.Join(dirname,".config/dotflow/")
  config := filepath.Join(config_folder, ".dotflow")
  var path_to_remove = strings.Replace(path, dirname, "$HOME", 1)

 	fileBytes, err := ioutil.ReadFile(config)

 	if err != nil {
    fmt.Println(err)
    os.Exit(1)
 	}

 	lines := strings.Split(string(fileBytes), "\n")
  config_tmp := filepath.Join(config_folder, ".dotflow_tmp")
  f, err := os.OpenFile(config_tmp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
      log.Fatal(err)
      return "", true
  }
  
  var og_path string

 	for _, line := range lines {
 		if !strings.Contains(line, path_to_remove){
      if _, err := f.Write([]byte(line)); err != nil {
        log.Fatal(err)
        return "", true
      }
 		} else {
      og_path = line 
    }
 	}
  if err := f.Close(); err != nil {
      log.Fatal(err)
      return "", true
  } 

  os.Rename(config_tmp, config)
  return og_path, false
}
