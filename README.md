# Persian Text Shaper

The `persian` package is a Go library designed to process and reshape Persian (Farsi) text to ensure its proper display in environments that do not natively support the bidirectional and context-sensitive shaping required by the Persian script.

## Features

- Identification of Persian letters and digits
- Adjustment of character forms based on position in a word (isolated, initial, medial, final)
- Reversal of character order for correct display in left-to-right environments
- Segmentation of mixed Persian and Latin script texts for individual processing
- Reshaping Persian words with proper character forms

## Installation

To use the `persian` package in your project, install it using `go get`:

```shell
go get -u github.com/rodrikv/persian
```

Make sure to replace `github.com/rodrikv/persian` with the actual path to the package.

## Usage

Here is a simple example of how to use the `persian` package to reshape a Persian text string:

```go
package main

import (
	"fmt"
	"your_module/persian"
)

func main() {
	inputText := "سلام"
	shapedText := persian.ReShape(inputText)
	fmt.Println(shapedText)
}
```

Output:

```
سلام
```

## Functions

- `ReShape(input string) string`: Takes an input string and returns its reshaped equivalent for proper display.
- `IsPersianLetter(ch rune) bool`: Checks if a rune is a Persian letter.
- `IsPersian(input string) bool`: Checks if a string contains only Persian script and spaces.
- `IsDigit(letter rune) bool`: Checks if a rune is a Persian digit.
- `IsWordDigit(word string) bool`: Checks if a word contains any Persian digit.

## How it Works

The package analyzes input text and determines if each character is a Persian letter. If so, it adjusts the shape of the letter depending on its position within the word. The algorithm accounts for the unique characteristics of the Persian script, where letters can take different forms depending on their context within a word.

## Limitations

- This package does not handle all edge cases and complex script rules, such as those involving ligatures.
- It may need to be combined with other text processing tools for full-fledged text rendering support.

## Contribution

Contributions to the `persian` package are welcome! If you would like to contribute, please fork the repository, make your changes, and submit a pull request.

## License

license?