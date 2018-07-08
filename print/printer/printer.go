package printer

import (
	"fmt"

	"github.com/fatih/color"
)

// Printer struct
type Printer struct {
	c      []color.Attribute
	prefix string
}

// New Printer object
func New(prefix string, clr ...color.Attribute) *Printer {
	// set color to prefix if color is available
	if len(clr) > 0 {
		prefix = color.New(clr...).Sprintf(prefix)
	}

	p := Printer{
		c:      clr,
		prefix: prefix,
	}
	return &p
}

// WithPrefix to append prefix with more string
func (p Printer) WithPrefix(prefix string) *Printer {
	newPrefix := prefix
	if len(p.c) > 0 {
		newPrefix = color.New(p.c...).Sprintf(newPrefix)
	}
	p.prefix = p.prefix + newPrefix
	return &p
}

// Print text
func (p *Printer) Print(v ...interface{}) {
	print(p.prefix, v...)
}

func print(prefix string, v ...interface{}) {
	// naively reject if only tag
	if len(v) == 0 {
		return
	}
	// return if parased argument is not valid
	parsedIntf := parseArgs(v...)
	if len(parsedIntf) == 0 {
		return
	}
	newIntf := []interface{}{prefix}
	newIntf = append(newIntf, parsedIntf...)
	fmt.Println(newIntf...)
}

// TODO: count on interface{} length and discard append to reduce memory use
func parseArgs(v ...interface{}) []interface{} {
	var newIntf []interface{}
	for key, val := range v {
		switch val.(type) {
		// dispatch if array of string
		case []string:
			arrOfString := val.([]string)
			for _, stringval := range arrOfString {
				newIntf = append(newIntf, stringval)
			}
			continue
		case nil:
			continue
		}
		newIntf = append(newIntf, v[key])
	}
	return newIntf
}
