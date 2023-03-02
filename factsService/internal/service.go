package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/Diflapuna/fuckingSlave/factsService/models"
)

const URL = "https://official-joke-api.appspot.com/random_joke"

func New() {
	mux := http.NewServeMux()
	mux.HandleFunc("/jokes", DadJoke)
	err := http.ListenAndServe(":6969", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func GetJoke(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)
	jsonResp := &models.Response{}
	json.Unmarshal(resBody, jsonResp)

	return jsonResp.Setup + "\n" + jsonResp.Punchline
}

func DadJoke(w http.ResponseWriter, r *http.Request) {
	joke := &models.Joke{}
	joke.Joke = GetJoke(URL)
	response, err := json.Marshal(joke)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(response)
}
									