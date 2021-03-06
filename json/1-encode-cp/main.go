package main

import (
	"encoding/json"
	"fmt"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan encode struct menjadi json.
// Lengkapi function EncodeToJson agar dapat mengembalikan nilai byte hasil dari encode objek Leaderboard.
// Modifikasi struct UserRank sehingga field Name menjadi name ,field Rank menjadi rank, dan field Email tidak ikut untuk diencode.

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	// TODO: answer here
	Name string `json:"name"`
	Email string `json:"-"`
	Rank int `json:"rank"`
}

type Leaderboard struct {
	Users []*UserRank
}

func EncodeToJson(leaderboard Leaderboard) ([]byte, error) {
	// TODO: answer here
	jsonData,err := json.Marshal(leaderboard)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func main(){
	leaderBoard := Leaderboard{
		[]*UserRank{
			{
				Name : "Dewi",
				Email : "test",
				Rank:1,
			},
			{
				Name : "Ghifa",
				Email : "test",
				Rank:2,

			},
		},
	}
	res, _ := EncodeToJson(leaderBoard)
	fmt.Println(string(res))

}
