package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan looping pada template.
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Peringkat ke-n: Nama", contoh: "Peringkat ke-1: Roger"

type UserRank struct {
	Name  string
	Email string
	Rank  int
}

type Leaderboard struct {
	Users []*UserRank
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
	textTemplate = `{{range .Users }}Peringkat ke-{{ .Rank }}: {{ .Name }}
	{{else}} 
	Invalid "struct" Users harus berupa array!
	{{end}}`

	tmpl, err := template.New("template1").Parse(textTemplate)

	if err != nil{
		panic(err)
	}

	var b bytes.Buffer
	err =  tmpl.Execute(&b,leaderboard )
	if err != nil{
		panic(err)
	}
	return b.Bytes(),nil
}

func main(){
	Leaderboard := Leaderboard{
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
