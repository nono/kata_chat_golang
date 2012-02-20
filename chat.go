package chat

import (
	"fmt"
)

type Chat struct {
	messages chan string
	joiners  chan *User
	users    []*User
}

type User struct {
	login string
	inbox chan string
}

func NewUser(login string) *User {
	inbox := make(chan string)
	return &User{login, inbox}
}

func NewChat() *Chat {
	c := &Chat{}
	c.messages = make(chan string)
	c.joiners = make(chan *User)
	c.users = make([]*User, 0)
	go c.Run()
	return c
}

func (c *Chat) Run() {
	var msg string
	var user *User
	for {
		select {
		case user = <-c.joiners:
			c.users = append(c.users, user)
			user.inbox <- "Bonjour"
		case msg = <-c.messages:
			for _, user = range c.users {
				user.inbox <- msg
			}
		}
	}
}

func (c *Chat) Join(user *User) {
	c.joiners <- user
}

func (u *User) Shout(c *Chat, m string) {
	msg := fmt.Sprintf("%s says %s", u.login, m)
	c.messages <- msg
}
