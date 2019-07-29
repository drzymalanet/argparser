package argparser

import (
	"strings"
)

// Argparser holds the information about switches
type Argparser struct {
	flags map[string]*flag
}

type flag struct {
	found bool
	vals  []string
}

// NewParser creates an argument parser with given flags
func NewParser(flags ...string) (parser *Argparser) {
	// Add empty string acting as positional arguments flag
	flags = append([]string{""}, flags...)
	parser = &Argparser{
		flags: make(map[string]*flag, 0),
	}
	for _, flagName := range flags {
		parser.flags[flagName] = &flag{
			found: false,
			vals:  make([]string, 0),
		}
	}
	return parser
}

// Parse maps the arguments to switches
func (parser *Argparser) Parse(args []string) *Argparser {
	// start with appending to positional args list
	prev := parser.flags[""]
	for _, argument := range args {
		// Check if the argument is a flag
		if _, ok := parser.flags[argument]; ok {
			prev = parser.flags[argument]
			prev.found = true
			continue
		}
		prev.vals = append(prev.vals, argument)
		prev.found = true
	}
	return parser
}

// Get returns the combined list of arguments for the given flags.
// If no flags given, return the list of non flag arguments.
func (parser *Argparser) Get(flags ...string) []string {
	// If no arguments provided
	if len(flags) == 0 {
		// Return the positional arguments
		flags = []string{""}
	}
	combined := []string{}
	for _, flagName := range flags {
		if flag, ok := parser.flags[flagName]; ok {
			combined = append(combined, flag.vals...)
		}
	}
	return combined
}

// Got returns true if any of the given flags were present in the argument list.
func (parser *Argparser) Got(flags ...string) bool {
	// If no arguments provided
	if len(flags) == 0 {
		// Check the positional arguments
		flags = []string{""}
	}
	for _, flagName := range flags {
		if flag, ok := parser.flags[flagName]; ok {
			if flag.found {
				return true
			}
		}
	}
	return false
}

// String returns the printable description of the Argparser
func (parser *Argparser) String() string {
	b := strings.Builder{}
	b.WriteString("Argparser[")
	for name, flag := range parser.flags {
		b.WriteString("<")
		b.WriteString(name)
		b.WriteString(";")
		for i, val := range flag.vals {
			if i != 0 {
				b.WriteString(",")
			}
			b.WriteString(val)
		}
		b.WriteString(";")
		if flag.found {
			b.WriteString("found:true")
		} else {
			b.WriteString("found:false")
		}
		b.WriteString(">")
	}
	b.WriteString("]")
	return b.String()
}
