# modular-iris
> Is a simple example iris framework with modular implementation, use a glide (https://github.com/Masterminds/glide) for package management

### Testing
> functional testing for each modules and http/api testing for whole apps 

### Install Modules
```go

import (
  "modular-iris/modules/hello"
.....
)

func main() {
  app := iris.New()

 ....
 ....

 hello.Install(app) // install module hello

 ....
 ....
}
```

### Structure Modules
- main.go -> main modules handle route or interfacing
- logic.go -> handle a logic of modules
- struct.go -> collection of a struct
