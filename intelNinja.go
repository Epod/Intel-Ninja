package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.Println("Intel Ninja Starting...")

	log.Println("Cleaning Up Old Files (if exist)")
	os.RemoveAll("raw")

	log.Println("Downloading Intel From Git Sources...")
	downloadGitIntel()

	//Make new folder to move only yara rules into
	log.Println("Cleaning Up Git Intel...")
	os.Mkdir("raw/yararules", 0755)
	filepath.Walk("raw/git/", cleanGitFolder)
	os.RemoveAll("raw/git/")
	log.Println("Extracted Relevant Rules...\n")
	filepath.Walk("raw/yararules/", verifyYaraRule)
	log.Println("Removed Index Files...\n")

	//Completed - list how many rules were IDed
	pattern := filepath.Join("raw/yararules/", "*")
	filesCount, _ := filepath.Glob(pattern)
	log.Printf("Completed!. Created %d rules.\n", len(filesCount))

}

func cleanGitFolder(path string, f os.FileInfo, err error) (e error) {

	if strings.HasSuffix(f.Name(), ".yar") {
		os.Rename(path, "raw/yararules/"+f.Name())
	} else if strings.HasSuffix(f.Name(), ".yara") {
		os.Rename(path, "raw/yararules/"+strings.TrimSuffix(f.Name(), "a"))
	}
	return
}

func verifyYaraRule(path string, f os.FileInfo, err error) (e error) {

	//Read File
	content, _ := ioutil.ReadFile(path)

	//Delete Index Files - We Only Want Raw Yara Rules
	if strings.Contains(string(content), "include \".") {
		os.Remove(path)
	}
	return
}

func createYaraIndex() {
	
}
