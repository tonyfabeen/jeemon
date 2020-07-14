== Jeemon
Control your daemons!

This project intends to provide an easy way to manage or daemonize your programs with Go.


# Usage

```
pid, err := daemon.Start("bundle exec rails server")
if err != nil {
    fmt.Printf("your program is running under pid: %d", pid)
}

running, _ = daemon.IsRunning(pid)
fmt.Printf("Is my program running? %b", running)

pid, _ = daemon.Stop(pid)
fmt.Printf("Program stopped %d", pid)
```

# Tests

`$ go test ./...`


# Author
[Tony Fabeen Oreste](mailto: tony.fabeen@gmail.com)
