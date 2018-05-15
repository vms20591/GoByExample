package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "s"))
	p("Join:", s.Join([]string{"go", "lang"}, "-"))
	p("Repeat:", s.Repeat("go", 5))
	p("Replace:", s.Replace("test", "t", "a", -1))
	p("Replace:", s.Replace("test", "t", "a", 1))
	p("Split:", s.Split("91-1243567890", "-"))
	p("ToLower:", s.ToLower("VMS"))
	p("ToUpper:", s.ToUpper("vms"))

	p()

	p("Length:", len("test"))
	p("Char:", "test"[1])
}
