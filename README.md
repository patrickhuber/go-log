# go-log
go logging abstration library

## getting started

```bash
go get github.com/patrickhuber/go-log
```

## usage

```golang
import (
    "github.com/patrickhuber/go-log"
)

func main(){
  logger := log.Default(log.WithLevel(log.Debug))
  logger.Trace("trace")
  logger.Debug("debug")
  logger.Info("info")
  logger.Warn("warn")
  logger.Error("error")
}
```

```bash
debug
info
warn
error
```