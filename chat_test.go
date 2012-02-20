package chat

import (
	"fmt"
	"testing"
)

func TestNewChat(t *testing.T) {
	chat := NewChat()
	if chat == nil {
		t.Fail()
	}
}

func TestJoinChat(t *testing.T) {
	chat := NewChat()
	john := NewUser("John")
	chat.Join(john)
	join := <-john.inbox
	if join != "Bonjour" {
		t.Fail()
	}
}

func TestShout(t *testing.T) {
	chat := NewChat()
	john := NewUser("John")
	chat.Join(john)
	<-john.inbox
	john.Shout(chat, "something")
	msg := <-john.inbox
	if msg != "John says something" {
		t.Fail()
	}
}

func TestTwoUsers(t *testing.T) {
	chat := NewChat()
	john := NewUser("John")
	jane := NewUser("Jane")
	users := []*User{john, jane}
	for _, user := range users {
		chat.Join(user)
		<-user.inbox
	}
	for i, user := range users {
		msg := fmt.Sprintf("message %d", i)
		user.Shout(chat, msg)
		for _, u := range users {
			m := <-u.inbox
			expected := fmt.Sprintf("%s says %s", user.login, msg)
			if m != expected {
				t.Fail()
			}
		}
	}
}
