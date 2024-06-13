package api

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/messages", handleMessages)
	mux.HandleFunc("/messages/", handleMessageByID)

	return mux
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		CreateMessage(w, r)
	case "GET":
		GetAllMessages(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleMessageByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetMessageByID(w, r)
	case "PUT":
		UpdateMessage(w, r)
	case "DELETE":
		DeleteMessage(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
