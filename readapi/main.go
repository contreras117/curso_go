//Ejemplo para consumir un api que devuelve un json con 100 elementos.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Post es una estructura para recibir onbjetos json de la api jsonplaceholder.typicode
type Post struct {
	UserId int    `json: "userId"` //De esta manera se matchea el nombre del valor en json con el nombre de la variable en go. Tambien funcioa XML y SQL.
	ID     int    `json: "id"`
	Title  string `json: "title"`
	Body   string `json: "body"`
}

func main() {
	var client = &http.Client{Timeout: 10 * time.Second}
	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := client.Get(url)

	if err != nil {
		panic(err.Error())
	}

	var post []Post

	err = json.NewDecoder(resp.Body).Decode(&post)

	if err != nil {
		panic(err.Error())
	}

	//fmt.Println(post)
	//fmt.Println(post[23])
	fmt.Printf("UserId: %d \n", post[3].UserId)
	fmt.Printf("ID: %d \n", post[42].ID)
	fmt.Printf("Title: %s \n", post[56].Title)
	fmt.Printf("Body: %s \n", post[83].Body)

}
