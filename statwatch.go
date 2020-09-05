package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("usage: statwatch dir fname '*.x'")
	}

	dir := os.Args[1]
	patterns := os.Args[2:]

	infos := findFiles(dir, patterns)

	if len(infos) == 0 {
		log.Fatalf("failed to find any files to watch")
	}

	log.Printf("watching %d files", len(infos))

	for {
		time.Sleep(500 * time.Millisecond)
		checkFiles(infos)
	}
}

func checkFiles(infoMap map[string]os.FileInfo) {

	for path, info := range infoMap {

		newInfo, err := os.Stat(path)
		if err != nil {
			log.Fatalf("failed to stat %s: %v", path, err)
		}

		if newInfo.ModTime().After(info.ModTime()) {
			log.Printf("%s was modified", path)
			os.Exit(0)
		}
	}
}

func findFiles(dir string, patterns []string) map[string]os.FileInfo {

	found := map[string]os.FileInfo{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			if info.Name() != "." && strings.HasPrefix(info.Name(), ".") {
				// log.Printf("skipping dir: %s, %s", path, info.Name())
				return filepath.SkipDir
			}

			// log.Printf("found dir: %s, %s", path, info.Name())
			return nil
		}

		for _, p := range patterns {
			match, err := filepath.Match(p, info.Name())
			if err != nil {
				return err
			}
			if match {
				found[path] = info
				// log.Printf("found file: %s, %s", path, info.Name())
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("failed to find files in dir %s: %v", dir, err)
	}

	return found
}
