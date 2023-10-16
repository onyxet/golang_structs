package main

import (
	"context"
	"os"
	"os/signal"
	"structs/pkg/second"
	"time"
)

func main() {
	fakePlayers := second.Players{
		second.Player{
			Player: "fakePlayer1",
			Answer: []string{"4", "Finland", "Canberra", "Delta", "Norway", "Calligraphy"},
		},
		second.Player{
			Player: "fakePlayer2",
			Answer: []string{"4", "Finland", "Canberra", "Delta", "Norway", "Calligraphy"},
		},
		second.Player{
			Player: "fakePlayer3",
			Answer: []string{"3", "Scotland", "Canberra", "Alpha", "Britain", "Calligraphy"},
		},
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
			os.Exit(0)
		case <-ctx.Done():
		}
	}()

	answers := make(chan second.Player)
	ticker := time.NewTicker(2 * time.Second)
	playerRounds := make([]chan second.Round, len(fakePlayers))
	for i := range playerRounds {
		playerRounds[i] = make(chan second.Round)
		go second.FakePlayer(ctx, playerRounds[i], answers, fakePlayers[i], fakePlayers[i].Answer[i])
		go second.FakeResults(ctx, answers, playerRounds[i])
	}

	second.FakeGenerateRoundGame(ctx, playerRounds, *ticker)
}
