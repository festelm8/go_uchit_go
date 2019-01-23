package main

import (
  "fmt"
  //"io/ioutil"
  "os"
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "encoding/json"
  "github.com/gorilla/mux"
)

type Book struct {
  ID      string `json:"id"`
  Title   string `json:"title"`
  Author *Author `json:"author"`
}

type Author struct {
  Firstname string `json:"firstname"`
  Lastname  string `json:"lastname"`
}

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range books {
       if item.ID == params["id"] {
          json.NewEncoder(w).Encode(item)
          return
        }
    }
   json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = strconv.Itoa(rand.Intn(1000000))
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"] {
            books = append(books[:index], books[index+1:]...)
            var book Book
            _ = json.NewDecoder(r.Body).Decode(&book)
            book.ID = params["id"]
            books = append(books, book)
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"] {
            books = append(books[:index], books[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(books)
}

func chipokomon(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
        os.Exit(1)
	}

	var responseObject Response
    _ = json.NewDecoder(response.Body).Decode(&responseObject)
    for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println("Chipokomon name: " + responseObject.Pokemon[i].Species.Name)
	}
}

func main() {
    r := mux.NewRouter()
    books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
    books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/books", createBook).Methods("POST")
    r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
    r.HandleFunc("/chipokomon", chipokomon).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", r))
}