package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"log"
	"os"
)

func downloadGitIntel() {

	fmt.Println("\nDownloading Yara Rules Git Repo...")
	_, err := git.PlainClone("raw/git/yara", false, &git.CloneOptions{
		URL:      "https://github.com/Yara-Rules/rules.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("Error pulling Yara-Rules repo. Skipping")
	}

	fmt.Println("\nDownloading Neo23x0 Rules Git Repo...")
	_, err = git.PlainClone("raw/git/neo23x0", false, &git.CloneOptions{
		URL:      "https://github.com/Neo23x0/signature-base.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("Error pulling Neo23x0 repo. Skipping")
	}

	fmt.Println("\nDownloading bartblaze Rules Git Repo...")
	_, err = git.PlainClone("raw/git/bartblaze", false, &git.CloneOptions{
		URL:      "https://github.com/bartblaze/Yara-rules.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("Error pulling bartblaze repo. Skipping")
	}

	fmt.Println("\nDownloading reversinglabs Rules Git Repo...")
	_, err = git.PlainClone("raw/git/reversinglabs", false, &git.CloneOptions{
		URL:      "https://github.com/reversinglabs/reversinglabs-yara-rules.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("Error pulling reversinglabs repo. Skipping")
	}

	fmt.Println("\nDownloading JPCERTCC Rules Git Repo...")
	_, err = git.PlainClone("raw/git/JPCERTCC", false, &git.CloneOptions{
		URL:      "https://github.com/JPCERTCC/jpcert-yara.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("Error pulling JPCERTCC repo. Skipping")
	}

	log.Println("Completed Intel Pull From Git Repos\n")
}
