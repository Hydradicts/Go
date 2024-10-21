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
## Data types
<table class="table table-bordered">
	<thead class="thead-light">
		<tr>
			<th>Data Types</th>
			<th>Usage</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>bool</td>
			<td>true or false</td>
		</tr>
		<tr>
			<td>int</td>	
			<td>number</td>
		</tr>
		<tr>
			<td>byte</td>	
			<td>alias for uint8</td>
		</tr>
		<tr>
			<td>float32</td>
			<td>decimal</td>
		</tr>
		<tr>
			<td>string</td>	
			<td>a string, text</td>
		</tr>
		<tr>
			<td>complex</td>
			<td>complex number</td>
		</tr>
		<tr>
			<td>rune</td>	
			<td>alias for int32, unicode code point</td>
		</tr>
		<tr>
			<td>uint</td>	
			<td>also a number</td>
		</tr>
	</tbody>
</table>
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
