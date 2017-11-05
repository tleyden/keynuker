package keynuker

import (
	"github.com/google/go-github/github"
	"github.com/golang-collections/collections/stack"
)


type EventStack struct {
	*stack.Stack
}

func NewEventStack() *EventStack {
	return &EventStack{
		Stack: stack.New(),
	}
}

func (this *EventStack) Pop() *github.Event {
	val := this.Stack.Pop()
	if val == nil {
		return nil
	}
	return val.(*github.Event)
}

func (this *EventStack) PopAll() []*github.Event {
	events := []*github.Event{}
	for {
		popped := this.Pop()
		if popped == nil {
			break
		}
		events = append(events, popped)
	}
	return events
}


func (this *EventStack) Push(value *github.Event) {
	if value == nil {
		return
	}
	this.Stack.Push(value)
}