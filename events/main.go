// --- Directions
// Create an 'eventing' library out of the
// Events class.  The Events class should
// have methods 'on', 'trigger', and 'off'.

package main

import "fmt"

func main(){
   event := Events{}

   // registering an event with handler(s)
   event.On("user_create", sendEmail, sendSMS)

   // triggering an event
   event.Trigger("user_create")

   // removing
	event.Off("user_create")

   fmt.Println()
	// triggering an event
	event.Trigger("user_create")
}

type callbackFunc func()

type Events struct {
	EventStore map[string][]callbackFunc
}


// Register an event handler
func (e *Events) On(eventName string, callback ...callbackFunc) {
	if _, ok := e.EventStore[eventName]; !ok {
		e.EventStore = make(map[string][]callbackFunc)
		e.EventStore[eventName] = append(e.EventStore[eventName], callback...)
	}else{
		e.EventStore[eventName] = append(e.EventStore[eventName], callback...)
	}
}

// Trigger all callbacks associated
// with a given eventName
func (e *Events) Trigger(eventName string) {
	if _, ok := e.EventStore[eventName]; !ok {
		return
	}else{
		for _, callback := range e.EventStore[eventName]{
			callback()
		}
	}
}

// Remove all event handlers associated
// with the given eventName
func (e *Events) Off(eventName string) {
	if _, ok := e.EventStore[eventName]; !ok {
		return
	}else{
		delete(e.EventStore, eventName)
	}
}


func sendEmail(){
	fmt.Println("an email has been sent to the new user!")
}

func sendSMS(){
	fmt.Println("an SMS has been sent to the new user!")
}