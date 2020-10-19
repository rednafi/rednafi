<!-- Zero width character is used to put extra blank lines before and after code -->

<h3>
    
```go
​
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	Stack := map[string][]string{
		"languages"	: {"Python", "Bash"},
		"databases"	: {"PostgreSQL", "Mongo", "Redis"},
		"misc"		: {"Docker", "Celery"},
		"ongoing"	: {"Django", "GraphQL", "Go"},
	}
	b, err := json.MarshalIndent(Stack, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
​
```
</h3>
