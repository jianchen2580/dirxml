package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/mholt/archiver"
)

func main() {
	dir := "/tmp/ttt"
	WatchDir(dir)
}

// watch dir changes
func WatchDir(dir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if strings.HasSuffix(event.Name, "rar") {
					log.Println("event:", event.Name)
					tmpDir := os.TempDir()
					tmpPath := filepath.Join(tmpDir, "xmldir")

					// err := archiver.Unarchive(event.Name, "/tmp/bbb")
					fmt.Println(tmpPath)
					err := archiver.Unarchive(event.Name, tmpPath)
					if err != nil {
						fmt.Println(err)
					}
					files, err := ioutil.ReadDir(tmpPath)
					if err != nil {
						log.Fatal(err)
					}
					for _, f := range files {
						path := filepath.Join(tmpPath, f.Name())
						defer os.Remove(path)
						users := xml2obj(path)
						for i := 0; i < len(users.Users); i++ {
							fmt.Println("User Type: " + users.Users[i].Type)
							fmt.Println("User Name: " + users.Users[i].Name)
							fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
						}

						fmt.Println(f.Name())
					}
				}
				// if event.Op&fsnotify.Write == fsnotify.Write {
				//     log.Println("modified file:", event.Name)
				// }
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
