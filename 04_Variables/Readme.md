# 04 Variables

## Variables
A variable is declared by := as seen in Chapter 03_Lists
Type declaration isn't needed with :=
```go
name := "John"
```
But there can be a variable typs such as
```go
var name string
```
## Arrays
Arrays are declared the same way variable with specific datatypes are declared
```go
var myarray []int
```
## List types
In lists types like int or string don't exist.
Once an element is added, it's added.
```go
l := list.New()
	e4 := l.PushBack("Doe")
	e1 := l.PushFront("John")
	l.InsertBefore(3,e4)
	l.InsertAfter(2,e1)
```
