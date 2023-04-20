package main

import (
	"ascii-art/utils"

	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("\nExample: go run . --output=banner.txt <string> <template>")
		os.Exit(1)
	}

	fontStyle := "template/standard.txt"
	if len(args) == 3 {
		utils.FontStyle(&fontStyle, args[2])
	}

	var output string
	pattern := regexp.MustCompile(`--output=`)
	match := pattern.MatchString(args[0])
	if match {
		output = args[0][len("--output="):]
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")

		fmt.Println("\nExample: go run . --output=banner.txt <string> <template>")
		os.Exit(1)
	}

	if !utils.IsCorrectWord(strings.Split(args[1], "\\n")) {
		fmt.Println("your string contains a non-ascii character")
		os.Exit(1)
	}

	textASCII := utils.HandleTheASCIIArt(fontStyle)
	var Affichage [][]string
	for _, current := range strings.Split(args[1], "\\n") {
		var line []string
		for _, elem := range current {
			line = append(line, textASCII[elem])
		}
		Affichage = append(Affichage, line)
	}

	utils.Write(Affichage, output)
}
