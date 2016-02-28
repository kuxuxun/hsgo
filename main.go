package main

import "net/http"

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8000", noCacheHeader)
	if err != nil {
		panic(err)
	}
}

func noCacheHeader(w http.ResponseWriter, r *http.Request) {

}
