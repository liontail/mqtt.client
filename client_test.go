package mqttclient

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestConnect(t *testing.T) {
	g := Goblin(t)
	g.Describe("Mqtt Connect", func() {
		client, err := Connect("", "", "localhost:1883")
		g.It("Should equal to nil", func() {
			g.Assert(err).Equal(nil)
		})
		g.It("Client should connect", func() {
			g.Assert(client.IsConnected()).Equal(true)
		})
	})
}
