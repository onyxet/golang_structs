package post

import "fmt"

type Box struct {
	SenderAddr, RecipientAddr string
}

func (b *Box) Send() {
	if b.RecipientAddr == "" {
		panic("no recipient address provided")
	}
	if b.SenderAddr == "" {
		panic("no sender address provided")
	}
	fmt.Println("Sending box to", b.RecipientAddr)
}
