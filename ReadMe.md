Project1 Overview

This project entailed developing a simple text editing and auto-correction tool. This tool will take as an input a text file name with the text that requires modifications and outputs the modified text into another text file. It performs various text transformations based various set conditions.

The program for this project will only take 3 arguments is name of program as the first index, name of the input file as the second index and the name of the outputfile as the third index.

To run this program; first create your inputfile say "sample.txt" then go run . sample.txt result.txt (where sample.txt and result.txt are your inputfile and outputfile respectively)

Features:

    Conversion of Hexadecimal and Binary Numbers:
        Replaces every instance of (hex) with the decimal version of the preceding hexadecimal number.
        Replaces every instance of (bin) with the decimal version of the preceding binary number.

    Case Modification:
        (up): Converts the preceding word to uppercase.
        (low): Converts the preceding word to lowercase.
        (cap): Capitalizes the preceding word.
        (up, <number>), (low, <number>), (cap, <number>): Modifies the specified number of preceding words accordingly.

    Punctuation Formatting:
        Ensures proper spacing around punctuation marks (., ,, !, ?, :, and ;).
        Handles special cases like groups of punctuation (...) or (!?).

    Handling Single Quotes (''):
        Places single quotes to the right and left of the word(s) enclosed between them without any spaces.

    Handling Indefinite Articles:
        Changes 'a' to 'an' if the next word begins with a vowel or 'h'.


The program will apply the specified modifications to the input text and save the result in the output file.
