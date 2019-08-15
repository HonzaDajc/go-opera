package ordering

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter"
)

func TestEventBuffer(t *testing.T) {
	nodes := inter.GenNodes(5)

	var ordered []*inter.Event
	r := rand.New(rand.NewSource(time.Now().Unix()))
	_ = inter.ForEachRandEvent(nodes, 10, 3, r, inter.ForEachEvent{
		Process: func(e *inter.Event, name string) {
			ordered = append(ordered, e)
		},
		Build: func(e *inter.Event, name string) *inter.Event {
			e.Epoch = 1
			return e
		},
	})

	processed := make(map[hash.Event]*inter.Event)
	push, _ := EventBuffer(Callback{

		Process: func(e *inter.Event) error {
			if _, ok := processed[e.Hash()]; ok {
				t.Fatalf("%s already processed", e.String())
				return nil
			}
			for _, p := range e.Parents {
				if _, ok := processed[p]; !ok {
					t.Fatalf("got %s before parent %s", e.String(), p.String())
					return nil
				}
			}
			processed[e.Hash()] = e
			return nil
		},

		Drop: func(e *inter.Event, peer string, err error) {
			t.Fatalf("%s unexpectedly dropped with %s", e.String(), err)
		},

		Exists: func(e hash.Event) *inter.Event {
			return processed[e]
		},
	})

	for _, rnd := range rand.Perm(len(ordered)) {
		e := ordered[rnd]
		push(e, "")
	}
}
