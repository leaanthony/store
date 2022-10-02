# Store

Store is a library for creating reactive state stores. It supports a subscription mechanism to notify state changes.
As it uses generics, it can be used for any datatype. It is a dependency free port of Svelte's state store. 

## Example

```go
package main

import (
    "github.com/leaanthony/store"
)

func main() {
	    // Create a new store
    s := store.New(1)

    // Subscribe to the store
    unsub := s.Subscribe(func(value int) {
        println("Store updated:", value)
    })

    // Update the store
    s.Set(2)
	// Output: "Store updated: 2"
	
	// Update the store value using a function
	s.Update(func(value int) int {
        return value + 1
    })
	
	// Output: "Store updated: 3"
	
	unsub()
	s.Set(4)
	
	// Output: (none)
	
}

```
