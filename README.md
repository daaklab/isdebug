# Detect whether a program is running under a debugger

Go has no official way to determine whether a program is currently being debugged.
This package provides the `isdebug.Enabled` variable to check whether the program is running in debug mode, along with utility packages to detect it.


## buildflags `debug`

Detects whether `-tags=debug` was set at compile time.

```go
import _ "github.com/daakdev/isdebug/buildflags"
```


## Environment variable `DEBUG`

Detects whether the `DEBUG` environment variable is set at runtime.

```go
import _ "github.com/daakdev/isdebug/env"
```

```sh
DEBUG=1 /path/to/program
```


## Command-line argument `--debug`

Detects whether `--debug` was passed as a command-line argument at runtime.

```go
import _ "github.com/daakdev/isdebug/cmdarg"
```

```sh
/path/to/program --debug
```


## Automatic detection via `/proc/self/status` (Linux only)

Detects whether `TracerPid` is present in `/proc/self/status`.

```go
import _ "github.com/daakdev/isdebug/procselfstatus"
```



# Bonus

There's also a `sloghandler` that ignores the configured level and always logs when `isdebug.Enabled` is `true`.

```go
import "github.com/daakdev/isdebug/sloghandler"

logger := slog.New(sloghandler.New(slog.NewJSONHandler(os.Stdout, nil)))
```
