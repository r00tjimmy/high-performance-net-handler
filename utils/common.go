package utils

import (
  "log"
  "os"
)

func CheckErr(err error) {
  if err != nil {
    panic(err.Error())
  }
}

/**
output debug info & save log file
 */
func DebugLog(content string) {
  if debug_flag == true {
    log.Println(content)
  }

  if log_flag == true {
    log_file_handle, err := os.OpenFile(log_file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
    CheckErr(err)
    log_file_handle.WriteString(content + "\n\n")
  }
}


/**
check os path exist or not
 */
func PathExists(path string) (bool, error) {
  _, err := os.Stat(path)
  if err == nil {
    return true, nil
  }
  return false, err
}