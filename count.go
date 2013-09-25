package main

import (
	"flag"
	"fmt"
	"github.com/josharian/countselectcases/github.com/kr/fs"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strings"
)

// casen holds the number/type of case
// statements inside a select statement.
type casen struct {
	comm uint32 // number of comm statements
	def  uint32 // number of default statements
}

// casenCount holds the number of instances
// of a particular casen that have been seen.
// It is the struct needed for making a slice out
// of a map[casen]uint32.
type casenCount struct {
	casen
	count uint32 // how many of this casen have we seen?
}

type byStatements []casenCount

func (x byStatements) Len() int      { return len(x) }
func (x byStatements) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x byStatements) Less(i, j int) bool {
	if x[i].comm != x[j].comm {
		return x[i].comm < x[j].comm
	}
	return x[i].def < x[j].def
}

type visitor struct {
	counts map[casen]uint32
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return v
	}

	sel, ok := node.(*ast.SelectStmt)
	if ok {
		var c casen
		for _, clause := range sel.Body.List {
			clause := clause.(*ast.CommClause)
			if clause.Comm == nil {
				c.def++
			} else {
				c.comm++
			}
		}

		v.counts[c]++
	}

	return v
}

func main() {
	dir := flag.String("d", "", "directory to walk")
	flag.Parse()
	walker := fs.Walk(*dir)

	v := &visitor{counts: make(map[casen]uint32)}

	for walker.Step() {

		if err := walker.Err(); err != nil {
			fmt.Printf("Error during filesystem walk: %v\n", err)
			continue
		}

		if walker.Stat().IsDir() || !strings.HasSuffix(walker.Path(), ".go") {
			continue
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, walker.Path(), nil, 0)
		if err != nil {
			// don't print err here; it is too chatty, due to (un?)surprising
			// amounts of broken code in the wild
			continue
		}

		ast.Walk(v, f)
	}

	// convert from map to slice; sort; display
	all := make([]casenCount, 0, len(v.counts))
	for key, val := range v.counts {
		all = append(all, casenCount{casen: key, count: val})
	}

	sort.Sort(byStatements(all))

	fmt.Println("Cases\tComm\tDefault\tCount")
	for _, c := range all {
		fmt.Printf("%d\t%d\t%d\t%d\n", c.comm+c.def, c.comm, c.def, c.count)
	}
}
