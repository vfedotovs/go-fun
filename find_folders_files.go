package main

import (
        "fmt"
        "os"
        "path/filepath"
        "strings"
        "log"
)

func scan_recursive(dir_path string, ignore []string) ([]string, []string) {

        folders := []string{}
        files   := []string{}

        // Scan
        filepath.Walk(dir_path, func(path string, f os.FileInfo, err error) error {

                _continue := false

                // Loop : Ignore Files & Folders
                for _, i := range ignore {

                        // If ignored path
                        if strings.Index(path, i) != -1 {

                                // Continue
                                _continue = true
                        }
                }

                if _continue == false {

                        f, err = os.Stat(path)

                        // If no error
                        if err != nil {
                                log.Fatal(err)
                        }

                        // File & Folder Mode
                        f_mode := f.Mode()

                        // Is folder
                        if f_mode.IsDir() {

                                // Append to Folders Array
                                folders = append(folders, path)

                        // Is file
                        } else if f_mode.IsRegular(){

                                // Append to Files Array
                                files = append(files, path)
                        }
                }

                return nil
        })

        return folders, files
}

func main() {

        folders, files := scan_recursive("./", []string{".git" , ".gitignore" })

        // Files
        for _, i := range files {
                fmt.Println(i)
        }

        // Folders
        for _, i := range folders {
                fmt.Println(i)
        }
}
