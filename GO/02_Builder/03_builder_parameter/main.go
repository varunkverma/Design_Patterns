package main

import (
	"fmt"
	"strings"
)

// the internal type, whose access is needed to be prevented
type email struct {
	from, to, subject, body string
}

// The Builder that builds an email type
type EmailBuilder struct {
	email *email
}

// fluent utilty methods of the Builder
func (eb *EmailBuilder) from(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @ character")
	}
	eb.email.from = from
	return eb
}

func (eb *EmailBuilder) to(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("email should contain @ character")
	}
	eb.email.to = to
	return eb
}

func (eb *EmailBuilder) subject(subject string) *EmailBuilder {
	eb.email.subject = subject
	return eb
}

func (eb *EmailBuilder) body(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}

// The hidden implementation
func sendEmailImplementation(email *email) {
	fmt.Printf("\nsending email \n to: %s\n from: %s\n subject: %s\n body: %s\n",
		email.to, email.from, email.subject, email.body)
}

type build func(*EmailBuilder)

// The Func Exposed
// Here the consumer sends an action/func, who is capable of receiving an EmailBuilder Type, in that they can interact with.
func SendEmail(action build) {
	eb := &EmailBuilder{
		email: &email{},
	}
	action(eb) // this action is like a callback func, it will populate the email builder
	sendEmailImplementation(eb.email)
}

func main() {
	SendEmail(
		func(eb *EmailBuilder) {
			eb.
				to("tom@test.com").
				from("micheal@test.com").
				subject("Greetings").
				body("Howdy")
		},
	)
}
