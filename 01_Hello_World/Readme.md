# Part 01 Hello World 🌎



## Hello World
Fmt package is needed, <br>
FMT gives us Println option, <br>
And this should look like this:
```go
package main

import "fmt"

func main() {
	fmt.Println("Something")
}
```
## FMT package
func Append(b []byte, a ...any) []byte <br>
func Appendf(b []byte, format string, a ...any) []byte <br>
func Appendln(b []byte, a ...any) []byte <br>
func Errorf(format string, a ...any) error <br>
func FormatString(state State, verb rune) string <br>
func Fprint(w io.Writer, a ...any) (n int, err error) <br>
func Fprintf(w io.Writer, format string, a ...any) (n int, err error) <br>
func Fprintln(w io.Writer, a ...any) (n int, err error) <br>
func Fscan(r io.Reader, a ...any) (n int, err error) <br>
func Fscanf(r io.Reader, format string, a ...any) (n int, err error)v
func Fscanln(r io.Reader, a ...any) (n int, err error) <br>
func Print(a ...any) (n int, err error) <br>
func Printf(format string, a ...any) (n int, err error) <br>
func Println(a ...any) (n int, err error) <br>
func Scan(a ...any) (n int, err error) <br>
func Scanf(format string, a ...any) (n int, err error) <br>
func Scanln(a ...any) (n int, err error) <br>
func Sprint(a ...any) string <br>
func Sprintf(format string, a ...any) string <br>
func Sprintln(a ...any) string <br>
func Sscan(str string, a ...any) (n int, err error) <br>
func Sscanf(str string, format string, a ...any) (n int, err error) <br>
func Sscanln(str string, a ...any) (n int, err error) <br>
type Formatter <br>
type GoStringer <br>
type ScanState <br>
type Scanner <br>
type State <br>
type Stringer <br>

## Use of a few FMT Packages
### Append()
```go
package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 10)
	fmt.Println(slice)
}

```

```
[1 2 3 4 5 10]
```
### Scan
```go
package main

import "fmt"

func main() {
	fmt.Println("Please give in your age")
	var a int
	fmt.Scan(&a)
}
```
