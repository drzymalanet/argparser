# argparser
This package allows for collecting more than one value for flag arguments. The standard `flags` package does not provide such capability.

## Installation

    go get github.com/drzymalanet/argparser

And then

    import "github.com/drzymalanet/argparser"

## Quick manual

1. Create parser with `NewParser(...)`,
2. Parse the input arguments with `Parse(...)`,
3. Check and get the flags with `Got(...)` and `Get(...)`.

## Features

### Easy to remember and simple to use

	func main() {
		parser := argparser.NewParser("--help", "--version").Parse(os.Args[1:])
		if parser.Got("--help") {
			log.Fatal("Add help message here")
		}
		if parser.Got("--version") {
			log.Fatal("Add version number here")
		}
		for _, arg := range parser.Get() {
			log.Printf("Got argument: %s", arg)
		}
	}

### Allows for multi-value switches

To retrieve arguments such as in: `app --input a b c --output d e f`

	func main() {
		parser := argparser.NewParser("--input", "--output").Parse(os.Args[1:])
		for _, f := range parser.Get("--input") {
			log.Printf("Got input file: %s", f)
		}
		for _, f := range parser.Get("--output") {
			log.Printf("Got output file: %s", f)
		}
	}

### Works well in complex circumstances

Given a task to build a parser for something like this:

    app a b c --input y1 y2 y3 --output x1 x2 -- u v w --input y4 y5

With the goal of parsing it to:
 - input flags: y1 y2 y3 y4 y5
 - output flags: x1 x2
 - and other args: a b c u v w
 
The code will look like:

	func main() {
		options := []string{"-i", "-o", "--input", "--output", "--help", "--"}
		parser := argparser.NewParser(options...).Parse(os.Args[1:])
		for _, f := range parser.Get("-i", "--input") {
			log.Printf("Got input file: %s", f)
		}
		for _, f := range parser.Get("-o", "--output") {
			log.Printf("Got output file: %s", f)
		}
		for _, tag := range parser.Get("", "--") {
			log.Printf("Found tag: %s", tag)
		}
	}

Result:

	$ app a b c --input y1 y2 y3 --output x1 x2 -- u v w --input y4 y5
	2019/07/29 19:16:56 Got input file: y1
	2019/07/29 19:16:56 Got input file: y2
	2019/07/29 19:16:56 Got input file: y3
	2019/07/29 19:16:56 Got input file: y4
	2019/07/29 19:16:56 Got input file: y5
	2019/07/29 19:16:56 Got output file: x1
	2019/07/29 19:16:56 Got output file: x2
	2019/07/29 19:16:56 Found tag: a
	2019/07/29 19:16:56 Found tag: b
	2019/07/29 19:16:56 Found tag: c
	2019/07/29 19:16:56 Found tag: u
	2019/07/29 19:16:56 Found tag: v
	2019/07/29 19:16:56 Found tag: w

## Full documentation
https://godoc.org/github.com/drzymalanet/argparser
