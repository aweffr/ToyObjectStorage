package objects

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
}

func get(w http.ResponseWriter, r *http.Request) {

}

func put(w http.ResponseWriter, r *http.Request) {

}
