package helper

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
)

func startWatch(filepath string, doneChan chan bool) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("error on watcher: ", err)
	}
	defer watcher.Close()

	// Add a path.
	err = watcher.Add(filepath)
	if err != nil {
		fmt.Println("error adding file: ", err)
	}

	lastModified := time.Now()

	fmt.Println("\nWatching on: ", filepath)

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					fmt.Println("⚠️ Watcher event channel closed")
					return
				}

				if event.Has(fsnotify.Write) && time.Since(lastModified) > time.Second {
					fmt.Println("\n💾 File saved! Checking your solution...")

					if CorrectExercice(filepath, true, true) {
						fmt.Println("▶️ Enter 'n' or 'next' to contine...")
						doneChan <- true
						doneChan <- true
						return
					}
				}

				lastModified = time.Now()

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("⚠️ Watcher error:", err)
			}
		}
	}()

	<-doneChan // Block until true
}
