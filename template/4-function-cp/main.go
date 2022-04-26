package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan function pada template.
// Lengkapi function CalculateScore yang berfungsi untuk menjumlahkan total score dari users
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Nama: Score", contoh: "Roger: 1000"
// Pada bagian terakhir, cetak "Total Score: total", contoh: "Total Score: 50000"

type UserRank struct {
	Name  string
	Email string
	Rank  int
	Score int
}

type Leaderboard struct {
	Users []*UserRank
}

func CalculateScore(leaderboard Leaderboard) int {
	// TODO: answer here
	var sum int
	for _, i:=range leaderboard.Users{
		sum += i.Score
	}
	return sum
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
}

func main(){
	leaderboard := Leaderboard{
		[]*UserRank{
			{
				Name: "Roger",
				Rank: 1,
			},
			{
				Name: "Tony",
				Rank: 2,
			},
			{
				Name: "Bruce",
				Rank: 3,
			},
			{
				Name: "Natasha",
				Rank: 4,
			},
			{
				Name: "Clint",
				Rank: 5,
			},
		},
	}
	res, _ := ExecuteToByteBuffer(Leaderboard)
	fmt.Println(string(res))
}
