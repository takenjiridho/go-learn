package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lib/pq"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	var error Error
	error.Message = message
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

var db *sql.DB

func main() {
	pgURL, err := pq.ParseURL("postgres://ygodvjeq:xQ6G-I9hggzFOuN6-uf67Rx9-XmmvjL2@elmer.db.elephantsql.com:5432/ygodvjeq")

	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgURL)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(ProtectedEndpoint)).Methods("GET")

	log.Println("Listen on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GenerateToken(user User) (string, error) {
	var err error
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Email is Missing.")
		return
	}

	if user.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Password is Missing.")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hash)

	err = db.QueryRow("insert into users (email, password) values($1, $2) RETURNING id;",
		user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server error.")
		return
	}

	user.Password = ""

	responseJSON(w, user)
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		var authorizationHeader string

		if bearerToken != "" {
			authorizationHeader = strings.Split(bearerToken, " ")[1]
		}

		if authorizationHeader != "" {
			if len(authorizationHeader) > 2 {
				token, error := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					respondWithError(w, http.StatusUnauthorized, error.Error())
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					respondWithError(w, http.StatusUnauthorized, error.Error())
					return
				}
			}
		} else {
			respondWithError(w, http.StatusUnauthorized, "No Authorization header provided.")
		}
	})
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, "Yes")
}

func ComparePasswords(hashedPssword string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPssword), password)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func login(w http.ResponseWriter, r *http.Request) {
	var user User
	var jwt JWT

	json.NewDecoder(r.Body).Decode(&user)

	password := user.Password

	rows := db.QueryRow("select * from users where email=$1", user.Email)
	err := rows.Scan(&user.ID, &user.Email, &user.Password)

	hashedPassword := user.Password

	if err != nil {
		log.Fatal(err)
	}

	token, err := GenerateToken(user)

	if err != nil {
		log.Fatal(err)
	}

	isValidPassword := ComparePasswords(hashedPassword, []byte(password))

	if isValidPassword {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Authorization", token)

		jwt.Token = token
		responseJSON(w, jwt)
	} else {
		respondWithError(w, http.StatusUnauthorized, "Invalid Password.")
	}
}
