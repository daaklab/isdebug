# プログラム実行時にDEBUG実行かを検知

Go言語にはデバッグ実行中かどうかを判断する方法が公式には存在しません。
このパッケージはプログラムがデバッグ実行かどうかを判断する`isdebug.Enabled`変数と、デバッグ実行かを検知するユーティリティパッケージを提供します。


## buildflags `debug`

コンパイル時に`-tags=debug`が設定されているかどうかを検知します。

```go
import _ "github.com/daaklab/isdebug/buildflags"
```


## 環境変数 `DEBUG`

実行時に環境変数`DEBUG`が設定されているかどうかを検知します。

```go
import _ "github.com/daaklab/isdebug/env"
```

```sh
DEBUG=1 /path/to/program
```


## コマンドライン引数 `--debug`

実行時にコマンドライン引数として`--debug`が渡されているかを検知します。

```go
import _ "github.com/daaklab/isdebug/cmdarg"
```

```sh
/path/to/program --debug
```


## `/proc/self/status`を利用した自動検知 (Linux only)

/proc/self/statusにTracerPidが存在するかを検知します。

```go
import _ "github.com/daaklab/isdebug/procselfstatus"
```



# おまけ

`isdebug.Enabled`が`true`時は設定レベルを無視して出力する`sloghandler`もあるよ。

```go
import "github.com/daaklab/isdebug/sloghandler"

logger := slog.New(sloghandler.New(slog.NewJSONHandler(os.Stdout, nil)))
```



