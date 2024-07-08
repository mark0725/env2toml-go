# env2toml-go
Convert env vars to toml text.

## Installation
```
go get https://github.com/mark0725/env2toml-go
```

## Syntax

__ split to .

```
APP_TITLE='TOML Example'
APP_OWNER__NAME='Tom Preston-Werner'
APP_DATABASE__ENABLED=true
APP_DATABASE__PORTS='[ 8000, 8001, 8002 ]'
APP_SERVERS__ALPHA__IP=10.0.0.1
APP_SERVERS__ALPHA__ROLE=frontend
APP_SERVERS__BETA__IP=10.0.0.2
APP_SERVERS__BETA__ROLE=backend
```    

PRIFIX: APP_

RESULT：
```toml
title="TOML Example" 

[owner]
name="Tom Preston-Werner" 

[database]
enabled=true 
ports=[ 8000, 8001, 8002 ] 

[servers]

[servers.alpha]
ip="10.0.0.1" 
role="frontend" 

[servers.beta]
ip="10.0.0.2" 
role="backend" 
```   

## Usage
```go
package main

import (
        "fmt"
        "github.com/joho/godotenv"
        env2toml "github.com/mark0725/env2toml-go"
)

func main() {
        err := godotenv.Load()
        if err != nil {
                fmt.Println("Error loading .env file")
        }

        result, err := env2toml.Parse("APP_")
        if err != nil {
                fmt.Println("Error:", err)
                return
        }
        fmt.Println(result)
}
```

## License

This project is licensed under the MIT license.
