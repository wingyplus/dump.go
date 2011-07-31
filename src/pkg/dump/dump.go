package dump

import (
	r "reflect"
	"fmt"
	"strconv"
	"io"
	"os"
)

var emptyString = ""

// Prints to the writer the value with indentation.
func Fdump(out io.Writer, v_ interface{}) {
	// forward decl
	var dump0 func(r.Value, int)
	var dump func(r.Value, int, *string, *string)

	done := make(map[string]bool)

	dump = func(v r.Value, d int, prefix *string, suffix *string) {
		pad := func() {
			res := ""
			for i := 0; i < d; i++ {
				res += "  "
			}
			fmt.Fprintf(out, res)
		}

		padprefix := func() {
			if prefix != nil {
				fmt.Fprintf(out, *prefix)
			} else {
				res := ""
				for i := 0; i < d; i++ {
					res += "  "
				}
				fmt.Fprintf(out, res)
			}
		}

		printv := func(o interface{}) { fmt.Fprintf(out, "%v", o) }

		//printf := func(s string, args ...interface{}) { fmt.Fprintf(out, s, args) }

		// prevent circular for composite types
		switch v.Kind() {
		case r.Array, r.Slice, r.Map, r.Ptr, r.Struct, r.Interface:
			//addr := v.Addr()
			key := fmt.Sprintf("%p %v", v.Interface(), v.Type())
			// if have value in done[key]
			if _, exists := done[key]; exists {
				padprefix()
				fmt.Printf("<%s>", key)
				return
			} else {
				done[key] = true
			}
		default:
			// do nothing
		}
		
		switch v.Kind() {
		case r.Array:
			padprefix()
			fmt.Fprintf(out, "%s {\n", v.Type())
			for i := 0; i < v.Len(); i++ {
				dump0(v.Index(i), d+1)
				if i != v.Len()-1 {
					fmt.Fprintf(out, ",\n")
				}
			}
			print("\n")
			pad()
			print("}")

		case r.Slice:
			padprefix()
			fmt.Fprintf(out, "%s (len=%d) {\n", v.Type(), v.Len())
			for i := 0; i < v.Len(); i++ {
				dump0(v.Index(i), d+1)
				if i != v.Len()-1 {
					fmt.Fprintf(out, ",\n")
				}
			}
			print("\n")
			pad()
			print("}")

		case r.Map:
			padprefix()
			t := v.Type()
			fmt.Fprintf(out, "map[%s]%s {\n", t.Key(), t.Elem())
			for i, k := range v.MapKeys() {
				dump0(k, d+1)
				fmt.Fprintf(out, ": ")
				dump(v.MapIndex(k), d+1, &emptyString, nil)
				if i != v.Len()-1 {
					fmt.Fprintf(out, ",\n")
				}
			}
			print("\n")
			pad()
			print("}")

		case r.Ptr:
			padprefix()
			
			if !v.Elem().IsValid() {
				fmt.Fprintf(out, "(%v) nil", v.Type())
			} else {
				print("&")
				dump(v.Elem(), d, &emptyString, nil)
			}

		case r.Struct:
			padprefix()
			t := v.Type()
			fmt.Fprintf(out, "%v {\n", t)
			d += 1
			for i := 0; i < v.NumField(); i++ {
				pad()
				printv(t.Field(i).Name)
				printv(": ")
				dump(v.Field(i), d, &emptyString, nil)
				if i != v.NumField()-1 {
					fmt.Fprintf(out, ",\n")
				}
			}
			d -= 1
			print("\n")
			pad()
			print("}")

		case r.Interface:
			padprefix()
			fmt.Fprintf(out, "(%v) ", v.Type())
			dump(v.Elem(), d, &emptyString, nil)

		case r.String:
			padprefix()
			fmt.Fprintf(out, "%v", strconv.Quote(v.Interface().(string)))
		
		case r.Bool,
			r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
			r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64,
			r.Float32, r.Float64:
			padprefix()
			//printv(v.Interface());
			/*
			i := v.Interface()
			if stringer, ok := i.(interface {
				String() string
			}); ok {
				printf("(%v) %s", v.Type(), stringer.String())
			} else {
				printv(i)
			}
			*/
			//printf("(%v) %v", v.Type(), v.Interface())
			fmt.Fprintf(out, "(%v) %v", v.Type(), v.Interface())
		
		case r.Invalid:
			padprefix()
			printv("nil")
		
		default:
			
			padprefix()
			fmt.Fprintf(out, "(%v) %v", v.Type(), v.Interface())
//			printv("nil")
		}
	}

	dump0 = func(v r.Value, d int) { dump(v, d, nil, nil) }

	v := r.ValueOf(v_)
	dump0(v, 0)
	fmt.Fprintf(out, "\n")
}

// Print to standard out the value that is passed as the argument with indentation.
// Pointers are dereferenced.
func Dump(v_ interface{}) { Fdump(os.Stdout, v_) }
