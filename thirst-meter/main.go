package main

import (
	"fmt"
	"math"
)

func main() {
	var temp float64
	fmt.Print("Enter water temperature in °C: ")
	fmt.Scanln(&temp)

	score, msg := thirstScore(temp)
	if score == 0 {
		fmt.Println("Thirst Satisfaction Score: 0")
		fmt.Println(msg)
	} else {
		fmt.Printf("Thirst Satisfaction Score: %d\n", score)
		fmt.Println(getScoreMessage(score))
	}
}

func thirstScore(temp float64) (int, string) {
	const (
		minDrinkable = 4.0
		maxDrinkable = 55.0
		idealTemp    = 20.0
	)

	if temp < minDrinkable {
		return 0, "🧊 Too cold! This is less water, more ice dagger."
	}
	if temp > maxDrinkable {
		return 0, "🔥 Too hot! Bro you boiling your soul."
	}

	// Calculate score: closer to 20°C = higher score
	diff := math.Abs(temp - idealTemp)
	score := 10 - int(diff/2)
	if score < 1 {
		score = 1
	}
	return score, ""
}

func getScoreMessage(score int) string {
	messages := map[int]string{
		10: "🐐 Peak hydration. You just drank the essence of life.",
		9:  "🔥 Smooth and satisfying. Your cells are clapping.",
		8:  "💧 Very nice. Thirst? What thirst?",
		7:  "😌 Clean hit. Respectable sip.",
		6:  "😐 Not bad. A bit off, but we vibin’.",
		5:  "🤔 Mid. Feels like licking a rain-soaked window.",
		4:  "😬 Hmm... warm-ish disappointment.",
		3:  "🥴 Getting weird. Your mouth is confused.",
		2:  "😖 Yikes. This ain't it, chief.",
		1:  "💀 You're drinking lava. Why tho?",
	}
	return messages[score]
}

