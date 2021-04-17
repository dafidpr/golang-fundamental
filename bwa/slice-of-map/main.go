package main

import "fmt"

func main() {

	/* Declaration dan initialization of slice map */
	students := []map[string]string{
		{"name": "Dafid Prasteyo", "score": "A"},
		{"name": "Yusup Supriyanto", "score": "C"},
		{"name": "Taufik Hidayat", "score": "B"},
	}

	for _, student := range students {

		fmt.Println("Nama Lengkap : ", student["name"], " - ", " Nilai : ", student["score"])
	}

	/* Quiz looking for an average score */
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var total int

	for _, score := range scores {
		total = total + score
	}
	average := float64(total) / float64(len(scores))
	fmt.Println("Avreage : ", average)

	fmt.Println("======================")

	var goodScores []int
	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println(goodScores)

}
