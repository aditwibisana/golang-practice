package main

import (
	"encoding/json"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan decode dari json menjadi objek dari struct.
// Mengambil kasus pada encode, lengkapi function DecodeToLeaderboard agar json dapat di decode menjadi objek Leaderboard

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

func DecodeToLeaderboard(jsonData []byte) (Leaderboard, error) {
	// TODO: answer here
	var leaderboard Leaderboard
	err := json.Unmarshal(jsonData, &leaderboard)

	if err != nil{
		return Leaderboard{}, err
	}
	return leaderboard, nil
}

// func main(){
// 	leaderBoard := Leaderboard{
// 		[]*UserRank{
// 			{
// 				Name: "Roger",
// 				Rank: 1,
// 			},
// 			{
// 				Name: "Tony",
// 				Rank: 2,
// 			},
// 			{
// 				Name: "Bruce",
// 				Rank: 3,
// 			},
// 			{
// 				Name: "Natasha",
// 				Rank: 4,
// 			},
// 			{
// 				Name: "Clint",
// 				Rank: 5,
// 			},
// 		},
// 	}
	
// 		res, err := DecodeToLeaderboard([]byte(jsonData))
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		for _, v:= range res.Users{
// 			fmt.Println(v.Name)
// 		}

// }