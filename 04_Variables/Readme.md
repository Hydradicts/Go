# 04 Variables

## Variables
A variable is declared by := as seen in Chapter 03_Lists
The variable can take all forms, as there is no type declaration.
```go
name := "John"
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
