package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var label []string

func main() {
	files := os.Args
	if len(files) == 4 {
		match, _ := regexp.MatchString(`^\-\-align\=(left|right|justify|center)`, files[3])
		match2, _ := regexp.MatchString(`^(thinkertoy|shadow|standard)`, files[2])
		if !match {
			errorHandler("OPTION")

		} else if !match2 {
			errorHandler("BANNER")
		}
		line, err := os.ReadFile(files[2] + ".txt")
		if err != nil {
			panic(err)
		}
		//It will split read lines
		//
		label = strings.Split(string(line), "\n")
		align := strings.Split(string(files[3]), "=")[1]
		justTXT := justfly(align, string(files[1]))
		word := strings.Split(string(justTXT), "\\n")
		//s := 0
		for i := 0; i < len(word); i++ {
			for j := 1; j < 9; j++ {
				for k := 0; k < len(word[i]); k++ {
					if int(word[i][k]) != 32 {
						fmt.Print(label[(int(word[i][k])-32)*9+j])
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}

		}

	} else {
		errorHandler("LESS")
	}
}
func justfly(align string, file string) string {
	var justTXT string
	number := 0
	screenWidth := getWidth()
	//println(screenWidth)

	resn := regexp.MustCompile(`(\s+\\n)|(\\n\s+)`)
	witOspace := resn.ReplaceAllString(file, "\\n")
	resn = regexp.MustCompile(`(^\s+)|((\s+)$)`)
	witOspace = resn.ReplaceAllString(witOspace, "")
	lengthArray := strings.Split(witOspace, "\\n")
	ren := regexp.MustCompile(`\\n`)
	res := regexp.MustCompile(`\s`)
	if align == "justify" {
		//ren := regexp.MustCompile(`\\n`)
		//res := regexp.MustCompile(`\s`)
		resn = regexp.MustCompile(`(\s+)`)
		witOspace2 := resn.ReplaceAllString(witOspace, "")
		newCount := len(ren.FindAllString(witOspace, -1))
		newLine := len(res.FindAllString(witOspace, -1))
		asciiLength, _ := readlength(ren.ReplaceAllString(witOspace2, ""))
		//println(asciiLength)
		if newCount+newLine != 0 {

			number = (screenWidth - asciiLength) / (newCount + newLine)
			//println(number)
		}

		number2 := 0
		var tempTXT string
		for i := 0; i < len(lengthArray); i++ {
			spacesN := len(strings.Split(lengthArray[i], " ")) - 1
			repeatSp := strings.Repeat(" ", number)
			lengthArray[i] = strings.ReplaceAll(lengthArray[i], " ", repeatSp)

			if i == 0 {

				justTXT = lengthArray[i]

				number2 = number * spacesN
			} else {

				tempTXT = strings.Repeat(" ", number2) + lengthArray[i]
				number2 += number * spacesN
				justTXT += "\\n" + tempTXT
			}
			prenumber2, _ := readlength(res.ReplaceAllString(lengthArray[i], ""))
			number2 = len(strings.Repeat(" ", number+number2)) + prenumber2

		}
	} else if align == "left" {
		spNumber := int(len(label[2]))
		repeatSp := strings.Repeat(" ", spNumber)
		justTXT = strings.ReplaceAll(witOspace, " ", repeatSp)

	} else if align == "right" {
		for i := 0; i < len(lengthArray); i++ {

			wordnumber, spacesN := readlength(lengthArray[i])
			repeatSp := strings.Repeat(" ", spacesN)
			lengthArray[i] = strings.ReplaceAll(lengthArray[i], " ", repeatSp)
			number2 := screenWidth - wordnumber
			if i == 0 {
				justTXT = strings.Repeat(" ", number2) + lengthArray[i]
			} else {

				justTXT += "\\n" + strings.Repeat(" ", number2) + lengthArray[i]
			}
		}

	} else if align == "center" {

		for i := 0; i < len(lengthArray); i++ {

			wordnumber, spacesN := readlength(lengthArray[i])
			repeatSp := strings.Repeat(" ", spacesN)
			lengthArray[i] = strings.ReplaceAll(lengthArray[i], " ", repeatSp)
			number2 := (screenWidth - wordnumber) / 2
			if i == 0 {
				justTXT = strings.Repeat(" ", number2) + lengthArray[i]
			} else {

				justTXT += "\\n" + strings.Repeat(" ", number2) + lengthArray[i]
			}
		}
	}
	return justTXT
}

func readlength(word string) (int, int) {
	count := 0

	for k := 0; k < len(word); k++ {

		count += len(label[(int(word[k])-32)*9+1])

	}
	spNumber := int(len(label[2]))
	return count, spNumber
}
func getWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	dimensions, _ := cmd.Output()
	width, _ := strconv.Atoi(strings.Split(strings.TrimSuffix(string(dimensions), "\n"), " ")[1])
	return width
}
func errorHandler(errorNo string) {
	switch errorNo {
	case "OPTION":
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard  --align=right\nWrong flag! Flag --align=<type>, in which type can be : \n\n- center\n\n- left\n\n- right\n\n- justify")
		os.Exit(0)
	case "BANNER":
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard  --align=right\nWrong Banner! Banner can be : \n\n- standard\n\n- thinkertoy\n\n- shadow")
		os.Exit(0)
	case "LESS":
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard  --align=right\nToo Few Parameters")
		os.Exit(0)

	}
}
