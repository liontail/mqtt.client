package test

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/liontail/mqtt.client"
)

func TestConnect(t *testing.T) {
	g := Goblin(t)
	g.Describe("Mqtt Connect", func() {
		client, err := mqttclient.Connect("", "", "localhost:1883")
		g.It("Should equal to nil", func() {
			g.Assert(err).Equal(nil)
		})
		g.It("Client should connect", func() {
			g.Assert(client.IsConnected()).Equal(true)
		})
		g.It("GetMessageFromBeginning", func() {
			data, err := mqttclient.GetMessageFromBeginning(client, "test", "central_accounts")
			g.Assert(err).Equal(nil)
			g.Assert(len(data) > 0).Equal(true)
		})
	})
}
