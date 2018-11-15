package main

import (
	"fmt"
	"os"
	"path/filepath"
	"Play-with-Data-Structures/10-Trie/03-Searching-in-Trie/src/BSTSet"
	"Play-with-Data-Structures/10-Trie/03-Searching-in-Trie/src/FileOperation"
	"time"
	"Play-with-Data-Structures/10-Trie/03-Searching-in-Trie/src/Trie"
)

func main() {
	fmt.Println("Pride and Prejudice")

	projectPath, _ := os.Getwd()
	currentPath := filepath.Join(projectPath, "10-Trie", "03-Searching-in-Trie")

	filename := filepath.Join(currentPath, "pride-and-prejudice.txt")
	words := FileOperation.ReadFile(filename)

	startTime := time.Now()

	bstSet := BSTSet.Constructor()
	for _, word := range words {
		bstSet.Add(word)
	}
	for _, word := range words {
		bstSet.Contains(word)
	}

	diffTime := time.Now().Sub(startTime)

	fmt.Println("Total different words:", bstSet.GetSize())
	fmt.Println("BSTSet:", diffTime)

	// ---

	startTime = time.Now()

	trie := Trie.Constructor()
	for _, word := range words {
		trie.Add(word)
	}
	for _, word := range words {
		trie.Contains(word)
	}

	diffTime = time.Now().Sub(startTime)

	fmt.Println("Total different words:", trie.GetSize())
	fmt.Println("BSTSet:", diffTime)
}