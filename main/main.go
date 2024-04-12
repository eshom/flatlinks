package main

import (
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

const USAGE string = "Usage: flatlinks WALK_DIR LINK_DIR"
const MAX_DIR_LEN = 100

// Global variable required until I figure out how to pass it to WalkFunc
var globalLinkPaths []string = make([]string, 0, 50)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("ERROR: Cannot get working directory. %s", err.Error())
	}

	if len(os.Args) != 3 {
		log.Fatalf("ERROR: Two arguments are expected.\n\n%s\n\n", USAGE)
	}

	log.Printf("Current directory: %s", root)

	walkDir := os.Args[1]
	linkDir := os.Args[2]

	if len(walkDir) > MAX_DIR_LEN || len(linkDir) > MAX_DIR_LEN {
		log.Fatalf("ERROR: Directory names over 100 characters are not allowed")
	}

	log.Printf("Walk Path: %s", walkDir)
	log.Printf("Generate links in: %s", linkDir)

	err = filepath.Walk(walkDir, fileLinkGen)
	if err != nil {
		log.Fatalf("ERROR: Problem walking directoy path. %s", err.Error())
	}

	paths := globalLinkPaths
	globalLinkPaths = nil

	log.Printf("Generating links for: %v ", paths)

	for idx := range paths {
		paths[idx], err = filepath.Rel(path.Join(root, path.Base(linkDir)), paths[idx])
		if err != nil {
			log.Fatalf("ERROR: Unable to detrmine relative link path. %s", err.Error())
		}
	}

	log.Printf("Link strings: %v", paths)

	err = os.Chdir(linkDir)
	if err != nil {
		log.Fatalf("ERROR: Unable to access %s : %s", linkDir, err.Error())
	}

	for _, pth := range paths {
		err = os.Symlink(pth, path.Base(pth))
		if err != nil {
			log.Printf("WARNING: Unable to create symlink to %s : %s", pth, err.Error())
		}
	}

	err = os.Chdir(root)
	if err != nil {
		log.Fatalf("ERROR: Unable to access %s : %s", root, err.Error())
	}
}

func fileLinkGen(path string, info fs.FileInfo, err error) error {
	isDir := info.IsDir()
	if isDir {
		log.Printf("Skipping directory: %s", path)
		return nil
	}

	isLink := (info.Mode()&fs.ModeSymlink)>>27 == 1
	if isLink {
		log.Printf("Skipping symlink: %s", path)
		return nil
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	globalLinkPaths = append(globalLinkPaths, absPath)

	return nil
}
