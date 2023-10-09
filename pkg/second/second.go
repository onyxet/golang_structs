package second

import (
	"context"
	"fmt"
	"time"
)

type Player struct {
	Player string
	Answer []string
}

type Players []Player

type Round struct {
	Question string
	Answer   string
}

type Rounds []Round

func FakeGenerateRoundGame(ctx context.Context, rounds []chan Round, ticker time.Ticker) {
	fakeQuestions := Rounds{
		Round{
			Question: "2+2=?",
			Answer:   "4",
		},
		Round{
			Question: "What country drinks the most coffee per capital?",
			Answer:   "Finland",
		},
		Round{
			Question: "What is the capital of Australia?",
			Answer:   "Canberra",
		},
		Round{
			Question: "What is the 4th letter of the Greek alphabet?",
			Answer:   "Delta",
		},
		Round{
			Question: "Roald Amundsen was the first man to reach the South Pole, but where was he from?",
			Answer:   "Norway",
		},
		Round{
			Question: "What art form is described as decorative handwriting or handwritten lettering?",
			Answer:   "Calligraphy",
		},
	}
	for _, question := range fakeQuestions {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Printf("[TICKER]It's time to new question! %s\n", question.Question)
			for _, roundChan := range rounds {
				roundChan <- question
			}
		}
	}
}

func FakePlayer(ctx context.Context, round chan Round, answers chan Player, player Player, answer string) {
	fmt.Printf("\n%s Is in the game!\n", player.Player)
	for {
		select {
		case <-ctx.Done():
			return
		case <-round:
			fmt.Printf("\nAnswer of %s: is %s\n", player.Player, answer)
			player.Answer = player.Answer[:0] // This smells like teen spirit
			player.Answer = append(player.Answer, answer)
			answers <- player
		}
	}
}

func FakeResults(ctx context.Context, answers chan Player, round chan Round) {
	for {
		select {
		case <-ctx.Done():
			return
		case player := <-answers:
			correctAnswer := <-round
			if player.Answer[0] == correctAnswer.Answer {
				fmt.Printf("\nPlayer %s has correct answer and it's %s\n", player.Player, player.Answer)
			} else {
				fmt.Printf("\nPlayer %s has incorrect answer and it's %s\n", player.Player, player.Answer)
			}
		}
	}
}
