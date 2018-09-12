// Package main is a program that sorts semantic versions read from stdin,
// and writes the same semvers to stdout, sorted in nondecreasing order.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/blang/semver"
)

// rawSemver contains a parsed semver and a version string.
type rawSemver struct {
	raw     string
	version semver.Version
}

type semvers []rawSemver

var _ sort.Interface = semvers{}

// Len implements sort.Interface.
func (s semvers) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s semvers) Less(i, j int) bool {
	return s[i].version.Compare(s[j].version) < 0
}

// Swap implements sort.Interface.
func (s semvers) Swap(i, j int) {
	t := s[i]
	s[i] = s[j]
	s[j] = t
}

// SortSemver reads semver words from r and writes sorted semvers to w.
func SortSemver(r io.Reader, w io.Writer) error {
	var sv semvers

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		v := rawSemver{}
		v.raw = s.Text()
		s := v.raw
		if s[0] == 'v' {
			s = v.raw[1:len(v.raw)]
		}
		var err error
		v.version, err = semver.Make(s)
		if err != nil {
			return fmt.Errorf("can not parse as semver: %q: %v", v.raw, err)
		}
		sv = append(sv, v)
	}
	sort.Sort(sv)
	for _, s := range sv {
		fmt.Fprintf(w, "%v\n", s.raw)
	}
	return nil
}

func main() {
	if err := SortSemver(os.Stdin, os.Stdout); err != nil {
		log.Fatalf("%v\n", err)
	}
}
