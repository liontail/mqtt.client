Mqtt Client Singleton 
======
## Example (How to use)

```go
package main

import (
    client "github.com/liontail/mqtt.client"
)

mqClient ,err := client.Connect(userName, password, url)
...

// or use this anywhere to get client 
mqClient := client.GetClient()

...

// To Subscribe topic

mqClient.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
})


// or use func ListenTo

f := func(msg mqtt.Message){
    fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
}

client.ListenTo(mqClient, topic, f)

data , err := client.GetMessageFromBeginning(client, "client_name", "db_name"))

// data = []byte
...
```