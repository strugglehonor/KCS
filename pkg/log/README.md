
### log 

clone from https://github.com/RichardKnop/logging


### Usage
```
 package main
 
 import (
 	"yourproject/infra-go-utils/log"
 )
 
 func main() {
 	log.INFO.Print("log message")
 	log.ERROR.Printf("some err: %+v", err)
 }
```

### TODO
 - set level
 - hook
 - tracing
 - concurrent