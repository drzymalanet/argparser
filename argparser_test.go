package argparser

import (
	"log"
	"runtime"
	"strings"
	"testing"
)

func ASSERT(t *testing.T, condition bool) {
	if !condition {
		t.Fail()
		_, file, line, _ := runtime.Caller(1)
		t.Logf("ASSERTION FAILED in \"%s\" Line: %d", file, line)
	}
}

func COMPARE(t *testing.T, slice1, slice2 []string) {
	if len(slice1) != len(slice2) {
		t.Fail()
		_, file, line, _ := runtime.Caller(1)
		t.Logf("COMPARE FAILED in \"%s\" Line: %d; Length %d != %d", file, line, len(slice1), len(slice2))
	}
	for i, e := range slice1 {
		if e != slice2[i] {
			t.Fail()
			_, file, line, _ := runtime.Caller(1)
			t.Logf("COMPARE FAILED in \"%s\" Line: %d; Element %d \"%s\" != \"%s\"", file, line, i, e, slice2[i])
		}
	}
}

func TestArgparser(t *testing.T) {
	checkEmptyArgs(t)
	checkEmptyFlags(t)
	checkNonFlags(t)
	checkSingleFlag(t)
	checkSingleFlagValue(t)
	checkTwoFlags(t)
	checkArgsFlagsArgs(t)
	checkFlagsArgsFlags(t)
}

func checkEmptyArgs(t *testing.T) {
	p := NewParser("a", "b", "c")
	args := []string{}
	p.Parse(args)
	ASSERT(t, p.Got("a") == false)
	ASSERT(t, p.Got("b") == false)
	ASSERT(t, p.Got("c") == false)
	ASSERT(t, p.Got("d") == false)
	ASSERT(t, len(p.Get("a")) == 0)
	ASSERT(t, len(p.Get("b")) == 0)
	ASSERT(t, len(p.Get("c")) == 0)
	ASSERT(t, len(p.Get("d")) == 0)
}

func checkEmptyFlags(t *testing.T) {
	p := NewParser()
	args := strings.Fields("a b c d")
	p.Parse(args)
	ASSERT(t, p.Got("a") == false)
	ASSERT(t, p.Got("b") == false)
	ASSERT(t, p.Got("c") == false)
	ASSERT(t, p.Got("d") == false)
	ASSERT(t, len(p.Get("a")) == 0)
	ASSERT(t, len(p.Get("b")) == 0)
	ASSERT(t, len(p.Get("c")) == 0)
	ASSERT(t, len(p.Get("d")) == 0)
	COMPARE(t, args, p.Get())
}

func checkNonFlags(t *testing.T) {
	p := NewParser("--a", "--b")
	args := strings.Fields("1 2 3 4")
	p.Parse(args)
	ASSERT(t, p.Got("--a") == false)
	ASSERT(t, p.Got("--b") == false)
	COMPARE(t, args, p.Get())
}

func checkSingleFlag(t *testing.T) {
	p := NewParser("--flag")
	args := strings.Fields("--flag")
	p.Parse(args)
	log.Printf(p.String())
	ASSERT(t, p.Got("--flag") == true)
	COMPARE(t, p.Get("--flag"), []string{})
	ASSERT(t, p.Got() == false)
	COMPARE(t, p.Get(), []string{})
}

func checkSingleFlagValue(t *testing.T) {
	p := NewParser("--flag")
	args := strings.Fields("--flag value")
	p.Parse(args)
	ASSERT(t, p.Got("--flag") == true)
	COMPARE(t, p.Get("--flag"), []string{"value"})
	ASSERT(t, p.Got() == false)
	COMPARE(t, p.Get(), []string{})
}

func checkTwoFlags(t *testing.T) {
	p := NewParser("--one", "--two")
	args := strings.Fields("--one 1 --two 2")
	p.Parse(args)
	ASSERT(t, p.Got("--one") == true)
	COMPARE(t, p.Get("--one"), []string{"1"})
	ASSERT(t, p.Got("--two") == true)
	COMPARE(t, p.Get("--two"), []string{"2"})
	ASSERT(t, p.Got("1") == false)
	COMPARE(t, p.Get("1"), []string{})
	ASSERT(t, p.Got("2") == false)
	COMPARE(t, p.Get("2"), []string{})
	ASSERT(t, p.Got() == false)
	COMPARE(t, p.Get(), []string{})
}

func checkArgsFlagsArgs(t *testing.T) {
	p := NewParser("--one", "--two", "--")
	args := strings.Fields("a b c --one 1 --two 2 -- d e f")
	p.Parse(args)
	ASSERT(t, p.Got() == true)
	ASSERT(t, p.Got("a") == false)
	ASSERT(t, p.Got("b") == false)
	ASSERT(t, p.Got("c") == false)
	ASSERT(t, p.Got("--one") == true)
	ASSERT(t, p.Got("1") == false)
	ASSERT(t, p.Got("--two") == true)
	ASSERT(t, p.Got("2") == false)
	ASSERT(t, p.Got("--") == true)
	ASSERT(t, p.Got("d") == false)
	ASSERT(t, p.Got("e") == false)
	ASSERT(t, p.Got("f") == false)
	COMPARE(t, p.Get(), strings.Fields("a b c"))
	COMPARE(t, p.Get(""), strings.Fields("a b c"))
	COMPARE(t, p.Get("", "--"), strings.Fields("a b c d e f"))
	COMPARE(t, p.Get("--one"), []string{"1"})
	COMPARE(t, p.Get("--two"), []string{"2"})
	COMPARE(t, p.Get("1"), []string{})
	COMPARE(t, p.Get("2"), []string{})
}

func checkFlagsArgsFlags(t *testing.T) {
	p := NewParser("--one", "--two", "--three", "--four", "--")
	args := strings.Fields("--one 1 --two 2 -- a b c --three 3 --four 4 -- d e f")
	p.Parse(args)
	ASSERT(t, p.Got() == false)
	ASSERT(t, p.Got("--one") == true)
	ASSERT(t, p.Got("1") == false)
	ASSERT(t, p.Got("--two") == true)
	ASSERT(t, p.Got("2") == false)
	ASSERT(t, p.Got("--") == true)
	ASSERT(t, p.Got("a") == false)
	ASSERT(t, p.Got("b") == false)
	ASSERT(t, p.Got("c") == false)
	ASSERT(t, p.Got("--three") == true)
	ASSERT(t, p.Got("3") == false)
	ASSERT(t, p.Got("--four") == true)
	ASSERT(t, p.Got("4") == false)
	ASSERT(t, p.Got("--") == true)
	ASSERT(t, p.Got("d") == false)
	ASSERT(t, p.Got("e") == false)
	ASSERT(t, p.Got("f") == false)
	COMPARE(t, p.Get(), []string{})
	COMPARE(t, p.Get(""), []string{})
	COMPARE(t, p.Get("--"), strings.Fields("a b c d e f"))
	COMPARE(t, p.Get("", "--"), strings.Fields("a b c d e f"))
	COMPARE(t, p.Get("--one"), []string{"1"})
	COMPARE(t, p.Get("--two"), []string{"2"})
	COMPARE(t, p.Get("--three"), []string{"3"})
	COMPARE(t, p.Get("--four"), []string{"4"})
	COMPARE(t, p.Get("1"), []string{})
	COMPARE(t, p.Get("2"), []string{})
	COMPARE(t, p.Get("a"), []string{})
	COMPARE(t, p.Get("b"), []string{})
	COMPARE(t, p.Get("c"), []string{})
	COMPARE(t, p.Get("3"), []string{})
	COMPARE(t, p.Get("4"), []string{})
	COMPARE(t, p.Get("d"), []string{})
	COMPARE(t, p.Get("e"), []string{})
	COMPARE(t, p.Get("f"), []string{})
}
