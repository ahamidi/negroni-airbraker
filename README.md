### Negroni Airbrake Middleware

Negroni middleware that captures errors being returned and sends them to 
[Airbrake](http://airbrake.io)

(Will probably work with most similar frameworks, YMMV)

#### Functionality

1. Capture Panics
1. Capture non `200` returned status codes
1. Parse `error` field from response

#### Usage

```go
// Import the package
import (
    nam "github.com/ahamidi/negroni-airbrake-middleware"
)

// "Use" middleware
n.Use(nam.NewMiddleware("<airbrake_app_id>", "<airbrake_app_key>"))
```

In order for the error message to be automatically parsed from the response,
the response should be in the following format:

```JSON
{
    "data": {
        "error": "Error message goes here"
    }
}
```

#### Building

1. Checkout Code
1. Get dependencies - `go get`
1. Build - `go build`

#### License

The MIT License (MIT)

Copyright (c) 2015 Ali Hamidi

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

