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

	atoanresult := mine.AToAn(filecontent)
	HexResult := mine.Hex(atoanresult)
	binResult := mine.Bin(HexResult)
	capResult := mine.Cap(binResult)

	// writting the modufied input text in the output file
	err = os.WriteFile(outputFile, []byte(capResult), 0o644)
	if err != nil {
		fmt.Print("Error writting output file", err)
		return
	}
}
