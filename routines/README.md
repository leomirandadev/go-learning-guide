# GO ROUTINES
If you want that your software run many functions in the same time, you can use the go routine feature.

For execute routine in background without wait the final you just have to put go in front of function call. Example:

```go
package main

func main(){
  go runAnything()
}

func runAnything() {
  // process
}
```

But you want that your software wait the final you can use the waitGroups. WaitGroups is simple to use and is in "sync" package.

The code with wait group will be:

```go
package main


func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // this is inform that 1 routine is in use
		go sayHello(i, &wg)
	}

	wg.Wait() // wait until the waitGroup counter be 0

}

func sayHello(i int, wg *sync.WaitGroup) {
  // "defer" indicate that the function will be execute on the final of this block of code.
	defer wg.Done() // wg.Done() will be take off 1 from the counter
	fmt.Println("Go routine", i)
}

```

