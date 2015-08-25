// You can use runtime.Stack to get the stack trace of all goroutines:
buf := make([]byte, 1<<16)
runtime.Stack(buf, true)
fmt.Printf("%s", buf)

