package main

import (
	"net/http"
	"github.com/bmizerany/pat"
	"encoding/json"
	"fmt"
)

func StartAgent(secret string, port int) error {
	
	m := pat.New()
	m.Post("/Cdn/Create", http.HandlerFunc(Create))

	//m.Post("/Cdn/Delete", http.HandlerFunc(startSync))
	//m.Post("/Cdn/List", http.HandlerFunc(startSync))
		
	http.Handle("/", accessControl(m,secret))
	
	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, nil)
	
	return err
}

func accessControl(handler http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := MessageModel{}
		result.Success = false
								
		key := r.Header.Get("Authorization")				
		w.Header().Set("Content-Type", "application/json")		
				
		if key == "" {

			result.Message = "Authorization is empty"
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		if key != secret {
					
			result.Message = "Invalid Credential. Access Denied"
			w.WriteHeader(http.StatusForbidden)
			return
		}
		
		handler.ServeHTTP(w, r)
	})	
}

func Create(w http.ResponseWriter, req *http.Request) {
		result := MessageModel{}
		result.Success = false	
	
		json.NewEncoder(w).Encode(result)	
}
	

type MessageModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`	
}