package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
)

func whenFileChangesThenValidate(eslFile string, folder string) {
	w := watcher.New()
	w.FilterOps(watcher.Write)
	w.IgnoreHiddenFiles(true)

	go func() {
		for {
			select {
			case event := <-w.Event:
				processEvent(event)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add(folder); err != nil {
		log.Fatalln(err)
	}

	// Initial validation on startup:
	for _, f := range w.WatchedFiles() {
		if strings.HasSuffix(f.Name(), eslFile) || strings.HasSuffix(f.Name(), defaultEmlFile) {
			fmt.Println("Initial validation: " + f.Name())
			isFileValidEslOrEml(f.Name())
		}
	}

	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Write, nil)
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func processEvent(event watcher.Event) {
	if !event.IsDir() && len(event.Path) >= 4 {
		isFileValidEslOrEml(event.Path)
	}
}

func isFileValidEslOrEml(fileName string) bool {
	if len(fileName) >= 4 {
		isValidatingEslFile := strings.HasSuffix(fileName, ".esl")
		isValidatingEmlFile := strings.HasSuffix(fileName, ".eml.yaml")
		if isValidatingEslFile {
			isValidEsl, err := convertFileToEml(fileName, generatedEmlFile)
			if err != nil {
				log.Panicln(err)
			}
			isValidEml, err := checkIfFileContainsValidEml(generatedEmlFile)
			if err != nil {
				log.Panicln(err)
			}
			if isValidEml && isValidEsl {
				fmt.Println("OK")
				return true
			}
			return false
		}
		if isValidatingEmlFile {
			isValidEsl, err := checkIfFileContainsValidEml(fileName)
			if err != nil {
				log.Panicln(err)
			}
			if isValidEsl {
				fmt.Println("OK")
				return true
			}
			return false
		}
	}
	return false
}
