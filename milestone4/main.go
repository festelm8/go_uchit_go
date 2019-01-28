package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "math/rand"
    "net/http"
    "strconv"
    "time"

    "github.com/dgrijalva/jwt-go"
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


type User struct {
  ID string `json:"id"`
  Name string `json:"name"`
}

type UserJWTClaims struct {
    UID string `json:"uid"`
    jwt.StandardClaims
}

type AuthData struct {
  Phone string `json:"phone_number"`
  Pass string `json:"password"`
}

type AuthToken struct {
  Token string `json:"auth_token"`
}

type ErrorResponse struct {
  Msg string `json:"msg"`
}


var books []Book
var mySigningKey = []byte("cockkekkok")


func UserCtx(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header["Authorization"] != nil {

            token, err := jwt.ParseWithClaims(r.Header["Authorization"][0], &UserJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
                return mySigningKey, nil
            })

            if err != nil {
				fmt.Fprintf(w, err.Error())
				return
			}

            if claims, ok := token.Claims.(*UserJWTClaims); ok && token.Valid {
                currentUser := User{ID:"123", Name:"CockUser"}
                if currentUser.ID == claims.UID{
                    fmt.Println(claims.UID)
                    fmt.Println(currentUser)
                    fmt.Println("mmm")
                    ctx := context.WithValue(r.Context(), "currentUser", currentUser)
                    next.ServeHTTP(w, r.WithContext(ctx))
                    return
                }
            }
        }

        fmt.Println("excess")
        ctx := context.WithValue(r.Context(), "currentUser", nil)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func GenerateJWT() (string, error) {
    claims := UserJWTClaims{
        "123",
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)

	return tokenString, err
}

func authLogin(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var authData AuthData
    _ = json.NewDecoder(r.Body).Decode(&authData)
    if authData.Pass != "kok" {
        w.WriteHeader(http.StatusForbidden)
        json.NewEncoder(w).Encode(ErrorResponse{ Msg: "wrong password" })
        return
    }

    validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
        return
	}
    json.NewEncoder(w).Encode(AuthToken{Token: validToken})
}


func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    ctx := r.Context()
    currentUser, _ := ctx.Value("currentUser").(*User)
    fmt.Println(currentUser)
    //if ok {
    //    fmt.Println(currentUser.Name)
    //}
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    for _, item := range books {
       if item.ID == chi.URLParam(r, "id") {
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
    for index, item := range books {
        if item.ID == chi.URLParam(r, "id") {
            books = append(books[:index], books[index+1:]...)
            var book Book
            _ = json.NewDecoder(r.Body).Decode(&book)
            book.ID = chi.URLParam(r, "id")
            books = append(books, book)
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    for index, item := range books {
        if item.ID == chi.URLParam(r, "id") {
            books = append(books[:index], books[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(books)
}



func main() {
    books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
    books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
    r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Post("/login", authLogin)

    r.Route("/books", func(r chi.Router) {
        r.Use(UserCtx)
        r.Get("/", getBooks)
        r.Post("/", createBook)

        r.Route("/{id}", func(r chi.Router) {
          r.Get("/", getBook)
          r.Put("/", updateBook)
          r.Delete("/", deleteBook)
        })
    })

    fmt.Println(">> Here we go! Server is run on :8000")
    http.ListenAndServe(":8000", r)
}