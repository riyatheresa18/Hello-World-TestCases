// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// // type Hello struct {
// // 	Value string `json:"value"`
// // }

// func newRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/hello", handler).Methods("GET")
// 	//r.HandleFunc("/post", posthandler).Methods("POST")
// 	return r
// }

// func main() {
// 	r := newRouter()
// 	http.ListenAndServe(":8080", r)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "123456")
// }

// // func posthandler(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")

// // 	// vars := mux.Vars(r)
// // 	// str :=r.Body
// // 	json.NewDecoder(r.Body).Decode(&hello)
// // 	json.NewEncoder(w).Encode(hello)
// // 	log.Println(hello)
// // }
