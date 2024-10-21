# 06 Work with Loops
## For Loop
For loop are the primary loops used in go. <br>
Other "loops" don't exist. <br>
Here is an example of a for loop
```go
	for i:=1; i<=5; i++ {
		fmt.Println("This loop is " + strconv.Itoa(i) +" out of 5")
	}
```
## While loop
While loops do exist, but not with the while keyword:
```go
	for i<5{
		fmt.Println("This loop is " + strconv.Itoa(i) +" out of 5")
		i++
	}
```
