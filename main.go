package main

import (
	"api-todo/models"
	service "api-todo/services"
	"encoding/json"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/google/uuid"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	server := service.App{}
	server.Init()
	r := mux.NewRouter()

	//AttachProfiler(r)

	r.HandleFunc("/todos", CorsHandler).Methods("OPTIONS")
	r.HandleFunc("/todos/{id}", CorsHandler).Methods("OPTIONS")

	r.HandleFunc("/todos", createHandler(server.GetTodos)).Methods("GET")
	r.HandleFunc("/todos/{id}", createHandler(server.GetTodo)).Methods("GET")
	r.HandleFunc("/todos", createHandler(server.CreateTodo)).Methods("POST")
	r.HandleFunc("/todos/{id}", createHandler(server.UpdateTodo)).Methods("PUT")
	r.HandleFunc("/todos/{id}", createHandler(server.DeleteTodo)).Methods("DELETE")
	r.HandleFunc("/todos", createHandler(server.DeleteCompleted)).Methods("DELETE")

	r.Use(loggingMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	http.Handle("/", r)
	log.Println("Listening port :8080")
	http.ListenAndServe(":8080", r)
}

func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.New().String()
		context.Set(r, "reqID", reqID)
		//log.Println(r.Method + " " + r.RequestURI + " " + reqID)
		log.Print(r.Method + " " + r.RequestURI + " ")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}

func CorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == http.MethodOptions {
		return
	}
}

func createHandler(serviceAction models.ServiceAction) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//create empty structure for response (we will pass it by reference)
		resp := models.ServiceResponse{}
		//create and fill structure for request
		req := models.ApiRequest{Vars: mux.Vars(r), Query: r.URL.Query(), Body: r.Body}

		err := serviceAction(req, &resp)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(http.StatusBadRequest)
			context.Clear(r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		apiResponse := models.ApiResponse{Data: resp.Response, ReqID: context.Get(r, "reqID")}
		json.NewEncoder(w).Encode(apiResponse)

		//log action result
		log.Println(http.StatusOK)
		context.Clear(r)
	}
	return http.HandlerFunc(fn)
}
