package observer

import (
	"sync"

	"github.com/google/uuid"
)

// Observer tracks subscribers for topics
type Observer struct {
	sync.Mutex
	subscribers map[string][]*Subscription
}

// Subscription is a single subscription to a topic.
// It contains a channel that will receive all future messages
// for the address
type Subscription struct {
	ID       string
	Topic    string
	Observer *Observer
	Ch       chan []byte
}

// Unsubscribe unsubscribes the subscription from the observer
func (s *Subscription) Unsubscribe() {
	s.Observer.Lock()
	defer s.Observer.Unlock()
	defer close(s.Ch)
	i := 0
	for _, sub := range s.Observer.subscribers[s.Topic] {
		if sub.ID == s.ID {
			break
		}
		i++
	}

	s.Observer.subscribers[s.Topic] = append(s.Observer.subscribers[s.Topic][:i], s.Observer.subscribers[s.Topic][i+1:]...)
}

// NewObserver returns a new Observer
func New() *Observer {
	return &Observer{
		subscribers: map[string][]*Subscription{},
	}
}

// Subscribe subscribes to an address
func (o *Observer) Subscribe(topic string) *Subscription {
	o.Lock()
	defer o.Unlock()

	_, found := o.subscribers[topic]
	if !found {
		o.subscribers[topic] = []*Subscription{}
	}
	id, _ := uuid.NewV7()
	subscription := &Subscription{
		ID:       id.String(),
		Topic:    topic,
		Observer: o,
		Ch:       make(chan []byte, 1),
	}
	o.subscribers[topic] = append(o.subscribers[topic], subscription)

	return subscription
}

// Notify notifies all subscribers of a particular topic of a message
func (o *Observer) Notify(topic string, msg []byte) {
	o.Lock()
	defer o.Unlock()
	if _, found := o.subscribers[topic]; !found {
		return
	}
	for _, sub := range o.subscribers[topic] {
		sub.Ch <- msg
	}
}
