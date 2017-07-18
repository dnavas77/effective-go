package main

import (
	"fmt"
)

// FORMATTING: automatically done by gofmt tool
type T struct {
	name  string // name of the object
	value int    // its value
}

/*
COMMENTARY:
	1. Every package should have a package comment:
		Package regexp implements a simple library for regular expressions.

		The syntax of the regular expressions accepted is:

			regexp:
				concatenation { '|' concatenation }
			concatenation:
				{ closure }
			closure:
				term [ '*' | '+' | '?' ]
			term:
				'^'
				'$'
				'.'
				character
				'[' [ '^' ] character-ranges ']'
				'(' regext ')'

	2. If the package is simple, the package comment can be brief
		// Package path implements utility goroutines for
		// manipulating slash-separated filename paths.

	3. Every exported (capitalized) name in a program should have a doc comment.

	4. Doc comments work best as complete sentences.
		// Compile parses a regular expression and returns, if successful,
		// a Regexp that can be used to match againts text.
		func Compile(str string) (*Regexp, error) {...

	5. A single doc comment can introduce a group of related constants or variables
		// Error codes returned by failures to parse an expression.
		var (
			ErrInternal      = errors.New("regexp: internal error")
			ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
			ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
			...
		)
*/

/*
NAMING CONVENTIONS:
	1. Package names are lowercase, single-word names; no underscores or mixedCaps

	2. Package name is based name of its source directory.
		e.g. "src/encoding/base64" is imported as "encoding/base64" but has name base64

	3. Do not use dot notation when importing packages.
	   Except in testing when a test needs to run outside a package it's testing.
		e.g. import (
				. "fmt" // allows to call methods directly without using package name
			 )

	4. Prefer "ring.New" instead of "ring.Ring"

	5. "Long names don't automatically make things more readable. A helpful doc comment
	can often be more valuable than an extra long name."

	6. Getters: given a variable "owner" don't do "GetOwner". Prefer "Ownner" and for
	setters, prefer "SetOwner"

	7. INTERFACE names:
		- One method interface names should be name same as the method plus and -er suffix
		or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier, etc.

	8. Semicolons: lexer automatically inserts them when scanning source.
	   Opening brace always goes in the same line as the keyword, otherwise the lexer
	   will add a semicolon before.
		   Good:
		   if i < f() {
			   f()
		   }

		   Bad:
		   if i < f()
		   {
			   g()
		   }

	9. Control Structures: Go does not have "do" or "while".
	   Always have if statements return so there's no need for "else"
*/

/*
REDECLARATION and REASSIGNMENT

	1. e.g:
		f, err := os.Open(name)
		d, err := f.Stat() // <-- uses same "err" variable above and reassigns it.

	2. Use the "blank identifier" to drop discard values in multiple assingments. e.g:
		sum := 0
		for _, value := range array {
			sum += value
		}
	3. the range keyword can break out individual unicode code points by parsing the UTF-8.

	4. "rune" is Go terminology for a single Unicode code point.

	5. Go has no comma operator and ++ and -- are statements not expressions.
	   if you want to run multiple variables in a "for" you shouls use parallel assigment.
	   -- Reverse a
	   for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		   a[i], a[j] = a[j], a[i]
	   }

	6. if "switch" has no expression it switches on "true". It's possible and idiomatic
	   to write an if-else-if-else chain as as switch.
	   Cases can be presented in comma-separated lists.
		   switch c {
		   case ' ', '?', '&', '=':
			   return true
		   }
	7. "break" can be used to break out of a switch statement early but it's not
	   common in Go. To break out of a loop put a label on top of the loop:
			Loop:
				for i := 0; i < 10; i++ {
					if i == 9 {
						break Loop
					}
				}
		"continue" also accepts an optional label but works for loops only.

		// Compare returns an integer comparing the two byte slices,
		// lexicographically.
		// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
		func Compare(a, b []byte) int {
			for i := 0; i < len(a) && i < len(b); i ++ {
				switch {
				case a[i] > b[i]:
					return 1
				case  a[i] < b[i]:
					return -1
				}
			}
			switch {
			case len(a) > len(b):
				return 1
			case len(a) < len(b)
				return -1
			}
			return 0
		}

	8. Type Switch:
		var t interface {}
		t = functionOfSomeType()
		switch t := t.(type) {
		case bool:
			fmt.Printf("boolean %t\n", t)
		case int:
			fmt.Printf("integer %d\n", t)
		default:
			fmt.Printf("unexpected type %T\n", t)
		}
*/

/*
FUNCTIONS:

	1. Multiple return values: func nextInt(b []byte, i int) (int, int) {...
	2. Named result params: funct nextInt(b []byte, i int) (value, nextPos int) {...
	3. Go's "defer" statement schedules a function call (the deferred function) to be run
	   immediately before the function executing the "defer" function returns.
*/

/*
DATA:

	1. "new": a built-in function that allocates memory, but unlike its namesakes in some
	other languages it does not initialize the memory, it only zeroes it.
	e.g:
		type SyncedBuffer struct {
			lock   sync.Mutex
			buffer bytes.Buffer
		}
		p := new(SyncedBuffer) // type *SyncedBuffer
		var v SyncedBuffer	   // type SyncedBuffer
*/

/*
CONSTRUCTORS and COMPOSITE LITERALS:

    1. Sometimes the zero value isn't good enough and and initializing constructor
	   is necessary (using composite literal, it's an expression that creates a new instance
	   each time is evaluated).
	   e.g:
			func NewFile(fd int, name string) *File {
				if fd < 0 {
					return nil
				}
				return &File{fd, name, nil, 0
			}
	2. The expressions "new(File)" and "&File{}" are equivalent
*/

/*
ALLOCATION with "make":
	1. Creates slices, channels, and maps only. And returns an initialized (not zeroed) value of type T (not *T)
	2. new([]int) returns a pointer to a newly allocated zeroed slice structure, that is a pointer to a nil slice.
	3. "make" vs "new":
		var p *[]int = new([]int) // allocates slice structure; *p == nil; rarely useful
		var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

		// unnecessarily complex:
		var p *[]int = new([]int)
		*p = make([]int, 100, 100)

		// Idiomatic:
		v := make([]int, 100)

		// three dots tell the compiler to make an array instead of a slice.
		// by first counting the number of elements in the braces.
		m := [...]float64{7.0, 8.5, 9.1}
*/

/*
ARRAYS:
	1. Assignning one array to another copies all the elements
	2. Passing an array to a function will receive a copy of the array, not a ponter
	3. The size of the array is part of the type.
*/

/*
SLICES:
	1. Slices are passed by value but elements in slice can be modified.
	2. Length of slices can change.
	3. capacity is the maximum length the slice may assume.

	func Append(slice, data []byte) []byte {
		l := len(slice)
		if l + len(data) > cap(slice) { // reallocate
			// Allocate double what's needed, for future growth.
			newSlice := make([]byte, (l+len(data)) * 2)
			// the copy function is predeclared and works for any slice type
			copy(newSlice, slice)
			slice = newSlice
		}
		slice = slice[0:1+len(data)]
		for i, c := range data {
			slice[l+i] = c
		}
		return slice
	}
*/

/*
TWO-DIMENSIONAL SLICES:
	1. Sometimes it's necessary to allocate a 2D slice, for example when processing
	   scan lines of pixels.

	   // Allocate the top-level slice, the same as before
	   picture := make([][]uint8, YSize) // one row per unit of Y
	   // allocate one large slice to hold all the pixels.
	   pixels := make([]uint8, XSize*YSize) // has type []uint8 even though picture is [][]uint8
	   // loop over the rows, slicing each row from the front of the remaining pixels slice
	   for i:= range picture {
		   picture[i], pixels = pixels[:XSize], pixels[XSize:]
	   }
*/

/*
MAPS:
	1. Holds a reference to an underlying data structure.
	2. If you change the contents of a map in a function it'll be visible in the caller.
	3. maps can be built using the composite literal, so it's easy to build them during
	   initialization.
	   var timeZone = map[string]int{
		   "UTC": 0*60*60,
		   "EST": -5*60*60,
	   }
	4. Assigning and fetching values from a map is the same as for arrays and slices.
	offset := timeZone["EST"]

	5. Attempting to fetch a map value with a key that is not present in the map will return
	   the zero value for the type of the entities in the map.
	   e.g.:
			attended := map[string]bool{
				"Ann": true,
			}
			if attented[person] { ... // if person is not in the map it'll be false

	6. Sometimes you need to distinguised a missing entry from a zero value.
		e.g.:
		// if tz is present, seconds will be set appropriately and ok will be true
		// if not, seconds will set to zero and ok will be false.
		seconds, ok := timeZone[tz]

	7. To test for presence you can use the blank identifier _
		e.g.:
			_, present := timeZone[tz]
	8. To delete a map entry use the delete built-in function, whose arguments are the map
	   and the key to be deleted.
	   e.g.:
			delete(timeZone, "PDT")
*/

/*
PRINTING:
	1. fmt.Printf, fmt.Fprintf, fmt.Sprintf (returns string)
		example:
		// all output the same "Hello 23"
		fmt.Printf("Hello %d\n", 23)
		fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
		fmt.Println("Hello", 23)
		fmt.Println(fmt.Sprint("Hello ", 23))

	2. %#v : prints the value in full Go syntax.
	3. %T : prints the type of a value
	4. If you want control of the default format for a custom type, all that's
	   required is to define a method with the signature String() on the type.
	   func (t *T) String() string {...}
	5. Dont' cause infinite recursion in String() for a type
		e.g.:

		*type MyString string
		func (m MyString) String() string {
			return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
		}
		// Fix, convert argument to string first
		func (m MyString) String() string {
			return fmt.Sprintf("MyString=%s", string(m)) // Error: will recur forever.
		}
*/

/*
APPEND:
	1. func append(slice []T, elements ...T) []T {...
*/

/*
INITIALIZATION:
	1. Is more powerful than in C. complex structures an be built during intialization.

	2. CONSTANTS:
		1 << 3 // constant expression
		math.Sin(math.Pi/4) // not constant 'cause math.Sin needs to happen at run time
	3. Enumerated constants are created using the "iota" enumerator.
			type ByteSize float64

			const (
				_           = iota // ignore first value by assigning to blank identifier
				KB ByteSize = 1 << (10 * iota)
				MB
				GB
				TB
				PB
				EB
				ZB
				YB
			)
	4. Sprintf will only call the String method when it wants a string. %f is safe
	   cause it wants a float.
*/

/*
VARIABLES:
	var (
		home = os.Getenv("HOME")
		user = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
*/

/*
THE INIT FUNCTION:
	1. each source file can identify its own niladic init function to setup whatever
	   state is required (each file can have multiple init function). init is called
	   after all the variable declarations in the package have evaluated their
	   initializers and after all import packages have been initialized.
	2. a common use of init functions is to verify or repair correctness of the program
	   state before real execution begins.
	   e.g.:
	   func init() {
		   if user == "" {
			   log.Fatal("$USER not set")
		   }
		   if home == "" {
			   home = "/home/" + user
		   }
		   if gopath == "" {
			   gopath = home + "/go
		   }
	   }

*/

/*
METHODS: Pointers vs. Values
		 1. methods can be defined for any named type (except pointer and interface)
		 2. define methods on slices. e.g.:

				 type ByteSlice []byte
				 func (slice *ByteSlice) Write(data []byte) (n int, err error) {
					 slice := *p
					 // same as above
					 *p = slice
					 return len(data), nil
				 }
			3. the rule about pointers vs values for receivers is taht value methods can be invoked on
				 pointers and values, but pointer methods can only be invoked on pointers.
			4.
*/

/*
INTERFACES and OTHER TYPES:
			1. A type can implement multiple interfaces
			2. Example:

				 type Sequence []int
				 // Methods required by sort.Interface
				 func (s Sequence) Len() int {
					 return len(s)
				 }
				 func (s Sequence) Less(i, j int) bool {
					 return s[i] < s[j]
				 }
				 func (s Sequence) Swap(i, j int) {
					 s[i], s[j] = s[j], s[i]
				 }

				 // Method for printing - sorts the elements before printing
				 func (s Sequence) String() string {
					 sort.Sort(s)
					 return fmt.Sprint([]int(s)) // <-- []int(s): converts named type to plain slice
				 }

CONVERSIONS:
			1. It's an idiom in Go programs to convert the type of an expression to access a different
				 set of methods.

INTERFACE CONVERSIONS and TYPE ASSERTIONS:
			1. e.g.:
				 if str, ok := value.(string); ok {
					 return str
				 } else if str, ok := value.(Stringer); ok {
					 return str.String()
				 }

GENERALITY:
			1. If a type exists only to implement an interface and will never have exported methods
				 beyond that interface, there is no need to export the type itself. In such cases
				 the constructor (implementation) should return an interface value rather than the implementing type.

INTERFACES AND METHODS:
			1. Implementation of a handler to count the number of times a page is visited.
				 // Simple counter server
				 type Counter struct {
					 n int
				 }

				 func (ctr *Counter) ServerHTTP(w http.ResponseWriter, req *http.Request) {
					 ctr.n++
					 fmt.Fprintf(w, "counter = %d\n", ctr.n)
				 }
			2. what if your program has some internal state that needs to be notified that a page
				 has been visited? Tie a channel to the web page. e.g.:
				 // A channel tht sends a notification on each visit.
				 // (Probably want the channel to be buffered.)
				 type Chan chan *http.Request

				 func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
					 ch <- req
					 fmt.Fprint(w, "notification sent")
				 }

				 // Print the server args
				 fun ArgServer() {
					 fmt.Println(os.Args)
				 }

				 // The HandlerFunc type is an adapter to allow the use of
				 // ordinary functions as HTTP handlers. If f is a function
				 // with the appropriate signature, HandlerFunc(f) is a
				 // Handler object that calls f.
				 type HandlerFunc func(ResponseWriter, *Request)

				 // ServeHTTP calls f(w, req)
				 func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
					 f(w, req)
				 }

*/

func main() {
	fmt.Println("Effective Go")
}
