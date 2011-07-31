package dump_test

import (
	. "dump"
	"testing"
//	"go/parser"
	"go/token"
//	"fmt"
)


var emptyString = ""

type S struct {
	A int
	B int
}

type T struct {
	S
	C int
}

type Circular struct {
	c *Circular
}

func TestDump(t *testing.T) {
	// empty

	// func ParseFile(filename string, src interface{}, scope *ast.Scope, mode uint) (*ast.File, os.Error)
/*
	file, e := parser.ParseFile(nil, "dump_test.go", nil, parser.ParseComments)
	if e != nil {
		fmt.Println("error", e)
	} else {
*/
//		fmt.Printf("%#v\n", file);
//		Dump(file)
		Dump(map[string]int{"satu": 1, "dua": 2})
		Dump([]int{1, 2, 3})
		Dump([3]int{1, 2, 3})
		Dump(&[][]int{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}})
		Dump(&emptyString)
		Dump(T{S{1, 2}, 3})
		Dump(token.STRING)

		bulet := make([]Circular, 3)
		bulet[0].c = &bulet[1]
		bulet[1].c = &bulet[2]
		bulet[2].c = &bulet[0]

		Dump(struct{ a []Circular }{bulet})
//	}
}
