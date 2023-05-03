// Test2 Zyon Morter 
package main

import (
	"log"
	"mime"
	"net/http"
)

//code from article 
func enforceJsonHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}
			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

//2
func middlewareA(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareA ")

		next.ServeHTTP(w, r) 

		log.Println("Executing middlewareA again (after the handler is finished)")
	})
}

func middlewareB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareB ")
		if r.URL.Path == "/tryout" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareB again ")
	})
}
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing a handler currently...")
	w.Write([]byte("It was Executed"))
}

func main() {
	mux := http.NewServeMux() //

	
	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", enforceJsonHandler(finalHandler))

	mux.Handle("/running", middlewareA(middlewareB(http.HandlerFunc(Handler)))) 
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}