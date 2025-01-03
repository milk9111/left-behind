package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sizeFlag := flag.String("cr", "2,2", "columns and rows of level in form \"[c],[r]\"")
	nameFlag := flag.String("n", "", "name of the level (lowercase will become filename)")
	dryRunFlag := flag.Bool("dry", false, "doesn't write the file")

	flag.Parse()

	if sizeFlag == nil || *sizeFlag == "" {
		panic("missing -cr flag")
	}

	if nameFlag == nil || *nameFlag == "" {
		panic("missing -n flag")
	}

	filename := fmt.Sprintf("assets/levels/%s.json", strings.ToLower(*nameFlag))
	f, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	if len(f) > 0 {
		panic("level already exists")
	}

	size := strings.Split(*sizeFlag, ",")
	if len(size) != 2 {
		panic("invalid -cr flag")
	}

	cols, err := strconv.Atoi(size[0])
	if err != nil {
		panic(err)
	}

	rows, err := strconv.Atoi(size[1])
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(schema{
		Name: *nameFlag,
		Cols: cols,
		Rows: rows,
		Data: make([]string, cols*rows),
	}, "", "\t")
	if err != nil {
		panic(err)
	}

	reg := regexp.MustCompile(`(\[[^[\]]+\])`)
	b = reg.ReplaceAllFunc(b, func(v []byte) []byte {
		t := &traverser{
			b: bytes.Join(bytes.Fields(v), nil),
		}

		var a []byte
		var startedNextVal bool
		var endedNextVal bool
		var colI int
		for t.hasNext() {
			c := t.consume()
			a = append(a, c)
			if c == '"' {
				if !startedNextVal {
					startedNextVal = true
				} else {
					endedNextVal = true
				}

				if endedNextVal {
					startedNextVal = false
					endedNextVal = false
					colI++
					if colI == cols {
						colI = 0

						if t.hasNext() && t.peek() == ',' {
							a = append(a, t.consume())
						}

						indent := "\t\t"
						if t.peek() == ']' {
							indent = "\t"
						}

						a = append(a, []byte("\n"+indent)...)
					}
				}
			} else if c == '[' {
				a = append(a, []byte("\n\t\t")...)
			}
		}

		return a
	})

	if !*dryRunFlag {
		if err = os.WriteFile(filename, b, os.ModeAppend); err != nil {
			panic(err)
		}
	} else {
		fmt.Println(string(b))
	}
}

type schema struct {
	Name string   `json:"name"`
	Cols int      `json:"cols"`
	Rows int      `json:"rows"`
	Data []string `json:"data"`
}

type traverser struct {
	i int
	b []byte
}

func (t *traverser) hasNext() bool {
	return t.i < len(t.b)
}

func (t *traverser) consume() byte {
	r := t.b[t.i]
	t.i++

	return r
}

func (t *traverser) peek() byte {
	return t.b[t.i]
}
