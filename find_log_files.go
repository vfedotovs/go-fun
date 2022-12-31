package main

import (
   "io/fs"
   "path/filepath"
   "os"
   "log"

)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func glob(root string, fn func(string)bool) []string {
   var files []string
   filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
      if fn(s) {
         files = append(files, s)
      }
      return nil
   })
   return files
}

func main() {
   files := glob(".", func(s string) bool {
      return filepath.Ext(s) == ".log"
   })

   f, err := os.Create("data.txt")

   if err != nil {
        log.Fatal(err)
   }

   defer f.Close()

   for _, file := range files {
      // println(file)
      line := file + "\n"
      _, err2 := f.WriteString(line)

      if err2 != nil {
        log.Fatal(err2)
      }

  }
}
