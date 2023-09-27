package post

import "fmt"

type Envelope struct {
	SenderAddr, RecipientAddr string
}

func (e *Envelope) Send() {
	if e.RecipientAddr == "" {
		panic("no recipient address provided")
	}
	if e.SenderAddr == "" {
		panic("no sender address provided")
	}
	fmt.Println("Sending envlope to", e.RecipientAddr)
}
