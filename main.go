Summary of Effective-Go: https://golang.org/doc/effective_go.html


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

	5. Long names don't automatically make things more readable. A helpful doc comment
	can often be more valuable than an extra long name.

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

FUNCTIONS:

	1. Multiple return values: func nextInt(b []byte, i int) (int, int) {...
	2. Named result params: funct nextInt(b []byte, i int) (value, nextPos int) {...
	3. Go "defer" statement schedules a function call (the deferred function) to be run
	   immediately before the function executing the "defer" function returns.

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

ARRAYS:
	1. Assignning one array to another copies all the elements
	2. Passing an array to a function will receive a copy of the array, not a ponter
	3. The size of the array is part of the type.

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

TWO-DIMENSIONAL SLICES:
	1. Sometimes its necessary to allocate a 2D slice, for example when processing
	   scan lines of pixels.

	   // Allocate the top-level slice, the same as before
	   picture := make([][]uint8, YSize) // one row per unit of Y
	   // allocate one large slice to hold all the pixels.
	   pixels := make([]uint8, XSize*YSize) // has type []uint8 even though picture is [][]uint8
	   // loop over the rows, slicing each row from the front of the remaining pixels slice
	   for i:= range picture {
		   picture[i], pixels = pixels[:XSize], pixels[XSize:]
	   }

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

APPEND:
	1. func append(slice []T, elements ...T) []T {...

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

VARIABLES:
	var (
		home = os.Getenv("HOME")
		user = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)

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
"
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
			1. Its an idiom in Go programs to convert the type of an expression to access a different
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

				 // Make ArgServer into an HTTP server
				 func ArgServer(w http.ResponseWriter, req *http.Request) {
					 fmt.Fprintln(w, os.Args)
				 }

				 // ArgServe now has the same signature as HandlerFunc so it can be converted
				 // to that type to access its methods.
				 http.Handle("/args", http.HandlerFunc(ArgServer))
				 // HTTP server will invoke ServeHTTP with ArgServer as receiver
				 // which then calls f(w, req) where f = ArgServer which will print the args:
				 // fmt.Fprintln(w, os.Args)

THE BLANK IDENTIFIER (_):
				 1. acts as a placeholder empty null value
				 2. do not use it for error values, always check error returns

UNUSED IMPORTS and VARIABLES:
				1. its an error to import a package without using it
				2. unused imports create unneessary bloat
				3. to silence complaints about unused imports use the blank identifier:
					import (
							"fmt"
							"io"
						)
					var _ = fmt.Printf // for debugging; delete when done
					var _ io.Reader // for debugging; delete when done

				4. global declarations to silence import errors should come right after
					the imports and be commented to remind yourself to clean things up later

IMPORT FOR SIDE EFFECT:
				1. sometimes its good to import a package only for its side effects without
					 any explicit use. For example, during its "init" function.
				2. the "net/http/pprof" package registers HTTP handlers that provide debugging
					 information.
				3. to import a pakage only for its side effects, rename the package to the blank
					 identifier:
							 import _ "net/http/pprof"
				4. if you import it with a name the compiler would reject the program
					 becuase named imports need to be used.

INTERFACE CHECKS:
				1. most interface conversions are static and therefore checked at compile time.
				2. some interface checks do happen at run-time. e.g.: in the "encoding/json"
					 package, which defineds a "Marshaler" interface.
					 When the JSON encoder receives a value that implements that interface,
					 the encoder invokes the value marshaling method to convert it to JSON
					 instead of doing standard conversion. The encoder checks this property
					 at runtime with a type assertion like:
				 			m, ok := val.(json.Marshaler)
				3. to ask whether a type implements an interface, use the blank identifier to
					 ignore the value:
				 			if _, ok := val.(json.Marshaler); ok {
								 //...
							}
			4. sometimes is a good idea to check if a type implements an interface
						var _ json.Marshaler = (*RawMessage)(nil)
								// - the assignment requires that *RawMessage implements Marshaler,
								// which will be checked at compile time.
								- The appearance of the blank identifier indicates that the declaration
									exists only for the type checking.
			5. Do not do this for every type that satisfies an interface, though.
			6. this declarations are only used when there are no static conversions already
					present in the code (which is rare).

EMBEDDING:
			1. Go does not have subclassing. but has ability to "borrow" pieces of an implementation
					by embedding types within a struct or interface.
			2. Interface embedding:
							type Reader interface {
								Read(p []byte) (n int, err error)
							}
							type Writer interface {
								Write(p []byte) (n int, err error)
							}
			3. io package exports interfaces that implement several methods.
					We could list both Read/Write methods but embedding interfaces is easier and
					more evocative.
					// ReadWriter is the interface that combines the Reader and Writer interfaces
						type ReadWriter interface {
							Reader
							Writer
						}

			4. only interfaces can be embedded within interfaces.
			5. the same idea with structs but far more reaching implications:
					A:
							// ReadWriter stores pointers to a Reader and a Writer struct
							// It implements io.ReadWriter
							type ReadWriter struct {
								*Reader // *bufio.Reader
								*Writer // *bufio.Writer
							}
							// the embedded elements are pointers and must be initialized to
							// valid structs before used.

							// Could be rewritten as:
							// but to promote the methods of the fields and satisfy io interfaces
							// we would also need to provide forwarding methods.
							// By embedding the structs directly, we avoid this bookkeeping.
							type ReadWriter struct {
							reader *Reader
								writer *Writer
							}

					B:
							The methods of embedded types come along for free. which means that
							bufio.ReadWriter not only has methods of bufio.Reader and bufio.Writer
							it also satisfies all three interfaces: io.Reader, io.Writer, and io.ReadWriter

					C:
							Difference between "embedding" and "subclassing":
								- when we embed a type, the methods of that type become methods of teh outer type.
								- when those methods are invoked, the receiver of the method is the inner type
									no the outer one. e.g.: when the "Read" method of a "bufio.ReadWriter" is invoked
									is has exactly the same effect as the forwarding method approach; the receiver
									is the "reader" field of the "ReadWriter", not the "ReadWriter" itself.

					D:
							Embedding can also be a simple convenience.
							// Job has now "Log", "Logf" and other methods of *log.Logger
							type Job struct {
								Command string
								*log.Logger
						}
								- "Logger" is a regular field of the Job struct. we can initialize it in the usual way
									inside the constructor for Job. e.g.:
											func NewJob(command string, logger *log.Logger) *Job {
												return &Job{command, logger}
											}
											// or with a composite literal
											job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
					E:
							Embedding types introduces the problem of name conflicts, but the rules
							to resolve them are simple:
										1. a field or method x hides any other item x in a more deeply nested part
										of the type. If "log.Logger" contained a field or method named "command"
										the "command" field of Job would dominate it.
										2. if the name apears at the same nesting level, its usually an error.
										itd be erronous to embed "log.Logger" if the struct had another field or
										method named "Logger".
										3. however, if the duplicate name is never mentioned in the program outside the
										type definition is OK.
										4. theres no problem if a field added that conflicts with another field if neither
										field is ever used.

CONCURRENCY:
	1. Share by communicating:
				- "Do not communicate by sharing memory; instead, share memory by communicating"
				- "Go's approach to concurrency originates in Hoare's Communicating Sequential Processes (CSP)"

	2. Goroutines:
			- A function executing concurrently with other goroutines in the same address space.
			- Allocated in stack space. Start small, so they are cheap, and grow by allocating
				(and freeing) heap storage as required.
			- goroutines are multiplexed onto multiple OS threads so they dont block eath other.
			- Prefix a funtion with "go" to run the call in a new goroutine. when it completes
				it exists silently (similar to a Unix shell & notation for running a command in the
				background)
			- e.g.:
						go list.Sort() // run list.Sort concurrently dont' wait for it

						// A function literal (closure) can be handy in a goroutine invocation
						func Announce(message string, delay time.Durarion) {
							go func() {
								time.Sleep(delay)
								fmt.Println(message)
							}() // Note the parentheses - must call the function
						}
			- These examples arent practical because the function above has no way of signaling
				completion. For that, we need channels.

CHANNELS:
	1. channels are alloated with "make()". Optional buffer size, default is zero.
		 ci := make(chan int) // Unbuffered channel of integers
		 cj := make(chan int, 0) // Unbuffered channel of integers
		 ck := make(chan *os.File, 100) // Buffered channel of pointers to Files

	2. Unbuffered channels combine "communication" (exchange of a value) with "synchronization"
		 guaranteeing that 2 calculations (goroutines) are in a known state.

	3. A channel can allow the launching goroutine to wait for the sort to complete.

		 c := make(chan int) // Allocate a channel
		 go func() { // Start the sort in a goroutine; when it completes, signal on the channel
			 list.Sort()
			 c <- 1 // Send a signal; value does not matter.
		 }
		 doSomethingForAWhile()
		 <- c // Wait for sort to finish; discard sent value

	4. Receivers always block until theres data to receive.
		 - if channel is Unbuffered, sender blocks until the receiver has received the value
		 - if channel has a buffer, sender blocks only until the value has been copied to
			 the buffer; if the buffer is full, this means waiting until some receiver has
			 retrieved a value.

	5. A buffered channel can be used like a semaphore, for instance to limit throughput.
		 - the capacity of the channel buffer limits the number of simultaneous calls to process.

			var sem = make(chan int, MaxOutstanding)

			func handle(r *Request) {
				sem <- 1 	// wait for active queue to drain.
				process(r) // May take a long time.
				<- sem			// Done; enable next request to run.
			}

			func Serve(queue chan *Request) {
				for req := range queue {
					sem <- 1
					go func(req *Request) {
						process(req)
						<- sem
					}(req) // we need to pass it to each closure so it's unique on each loop
				}
			}

		- Another approach that manages resources well is to start a fixed number of "handle"
			goroutines all reading from the request channel. The number of goroutines limits the
			number of simultaneous calls to "process". This "Serve" function also accepts
			a channel on which it will be told to exit; after launching the goroutines it blocks
			receiving from that channel.

			func handle(queue chan *Request) {
				for r := range queue {
					process(r)
				}
			}

			func Serve(clientRequests chan *Request, quit chan bool) {
				// Start handlers
				for i := 0; i < MaxOutstanding; i++ {
					go handle(clientRequests)
				}
				<- quit // Wait to be told to exit
			}

CHANNELS OF CHANNELS:
	1. Channels are first class values so they can be passed around like any other value.

	2. A common use of this property is to implement safe, parallel demultiplexing.

	3.  In the example in the previous section, handle was an idealized handler for a request but we didnt define the type it was handling. If that type includes a channel on which to reply, each client can provide its own path for the answer. Heres a schematic definition of type Request.

			type Request struct {
				args 				[]int
				f 					func([]int) int
				resultChan	chan int
			}

			// The client provides a function and its arguments, as well as a channel inside the request
			// object on qhich to receive the answer.
			func sum(a []int) (s int) {
				for _, v := range a {
					s += v
				}
				return
			}

			request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
			// Send request
			clientRequests <- request
			// Wait for response.
			fmt.Printf("answer: %d\n", <- request.resultChan)