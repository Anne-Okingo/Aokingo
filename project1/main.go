package main

import (
	"fmt"
	"os"

	"go-reloaded/mine"
)

func main() {
	inputfile := "sample.txt"
	outputFile := "result.txt"
	// reading our inputfile
	content, err := os.ReadFile(inputfile)
	// handling errors
	if err != nil {
		fmt.Println("Error File", err)
		return
	}

	// convert the content to string
	filecontent := string(content)
	// to modify out input text

	capResult := mine.Cap(filecontent)
	atoanresult := mine.AToAn(capResult)
	HexResult := mine.Hex(atoanresult)
	binResult := mine.Bin(HexResult)
	upResult := mine.Up(binResult)
	lowResult := mine.Low(upResult)
	puncResult := mine.Punc(lowResult)
	apostResult := mine.Apostrophe(puncResult)

	// writting the modufied input text in the output file
	err = os.WriteFile(outputFile, []byte(apostResult), 0o644)
	if err != nil {
		fmt.Print("Error writting output file", err)
		return
	}
}
 