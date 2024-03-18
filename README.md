# go eclai

## what?

A simple API client for [Acinq's Eclair](https://github.com/ACINQ/eclair) [lightning node JSON API](https://acinq.github.io/eclair/#introduction)



### Examples

Some helpful examples are available in [the examples folder](./examples)

### Create Invoice

```go
client = client.WithBaseURL("http://localhost:8282")

invoice, err := client.CreateInvoice({
    Description: "Drugqs"
    Amount: 100000,
})
if err != nil {
    panic(err)
}

fmt.Println(invoice.Serialized)

```


#### Watching Events

It is possible to listen for a handful of events that the node emits. This is done over websocket. Events that be subscribed to:

| Event Type                | Description                                                   |
|---------------------------|---------------------------------------------------------------|
| `payment-received`        | A payment has been received                                   |
| `payment-relayed`         | A payment has been successfully relayed                       |
| `payment-sent`            | A payment has been successfully sent                          |
| `payment-settling-onchain`| A payment wasn't fulfilled and its HTLC is being redeemed on-chain |
| `payment-failed`          | A payment failed                                              |
| `channel-created`         | A channel opening flow has started                            |
| `channel-opened`          | A channel opening flow has completed                          |
| `channel-state-changed`   | A channel state changed (e.g. going from offline to connected)|
| `channel-closed`          | A channel has been closed                                     |
| `onion-message-received`  | An onion message was received                                 |

To open a channel and watch for events

```go
    client = client.WithBaseURL("http://localhost:8282") // if you are using polar for a local setup

	channel, err := client.Subscribe()
	if err != nil {
		panic(err)
		return
	}

    // example fo handling events
	for message := range channel {
		switch message.Type {
            case eclair.ChannelOpened:
                // handle channel opened event
            case eclair.ChannelClosed:
                // handle channel closed event
            case eclair.PaymentReceived
                // handle channel closed event
		}
	}
```
