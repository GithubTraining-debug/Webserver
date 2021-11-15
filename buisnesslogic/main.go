package main

import (
	"backend/database"
	"backend/logic/getallbyname"
	"backend/logic/getallmovies"
	"backend/logic/savemovie"
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func backend() {
	http.HandleFunc("/movie", handleMovieReq)
	http.HandleFunc("/movies", getAllMovies)
	http.ListenAndServe(":8900", nil)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	movies, err := getallmovies.Setup()
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	jMovies, err := json.Marshal(movies)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	fmt.Fprintf(w, string(jMovies))
}

func handleMovieReq(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Innan response")
	setupResponse(&w, r)
	switch r.Method {
	case "GET":
		name := r.URL.Query()["name"][0]
		println(name)
		movies, err := getallbyname.Setup(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jMovies, err := json.Marshal(movies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, string(jMovies))

	case "POST":
		var movie models.Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("movie: %v\n", movie)
		err = savemovie.SetupSaveIt(movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	//os.Setenv("DATABASEPATH", "172.0.0.0")
	//fmt.Printf("os.Getenv(\"HELLO\"): %v\n", os.Getenv("HELLO"))
	//backend()
	database.StartDB()
}
