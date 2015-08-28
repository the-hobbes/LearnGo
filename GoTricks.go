// You can use runtime.Stack to get the stack trace of all goroutines:
buf := make([]byte, 1<<16)
runtime.Stack(buf, true)
fmt.Printf("%s", buf)

// In general, when you want to build up a string over a series of statements, 
// it's much more efficient to use a bytes.Buffer than to repeatedly concatenate:

func catFoos(foos []string) string {
  var b bytes.Buffer
  for _, foo := range foos {
    b.WriteString(foo)
  }
  return b.String()
}

