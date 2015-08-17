package main

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
	"strings"
)

var nginx Nginx

func StartAgent(secret string, port int) error {

	m := pat.New()
	m.Post("/Cdn/Create", http.HandlerFunc(Create))
	m.Del("/Cdn/Delete", http.HandlerFunc(Delete))
	m.Get("/Cdn/List", http.HandlerFunc(List))

	http.Handle("/", accessControl(m, secret))

	addr := fmt.Sprintf(":%d", port)
	log.Println("Start Server: " + addr)
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
			err := json.NewEncoder(w).Encode(result)

			if err != nil {
				log.Fatal(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		if key != secret {

			result.Message = "Invalid Credential. Access Denied"
			w.WriteHeader(http.StatusForbidden)
			err := json.NewEncoder(w).Encode(result)

			if err != nil {
				log.Fatal(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		handler.ServeHTTP(w, r)
	})
}

func Create(w http.ResponseWriter, req *http.Request) {
	result := MessageModel{}
	result.Success = false

	isSSLcdn := false
	isFullcdn := false

	domainName := req.FormValue("name")
	ipaddr := req.FormValue("ipaddr")
	port := req.FormValue("port")
	sslReq := req.FormValue("ssl")
	fullReq := req.FormValue("full")

	log.Println("domainName: " + domainName)
	log.Println("ipaddr: " + ipaddr)
	log.Println("port: " + port)

	w.WriteHeader(http.StatusOK)

	if domainName == "" {
		result.Message = "name parameter is null"
		json.NewEncoder(w).Encode(result)
		return
	}

	if ipaddr == "" {
		result.Message = "ipaddr parameter is null"
		json.NewEncoder(w).Encode(result)
		return
	}

	if port == "" {
		result.Message = "port parameter is null"
		json.NewEncoder(w).Encode(result)
		return
	}

	if sslReq != "" {
		sslReq = strings.ToLower(sslReq)
		if sslReq == "true" {
			isSSLcdn = true
		}
	}

	if fullReq != "" {
		fullReq = strings.ToLower(fullReq)
		if fullReq == "true" {
			isFullcdn = true
		}
	}

	result = nginx.Create(domainName, ipaddr, port, isSSLcdn, isFullcdn, config.Api.ConfigRoot)
	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func Delete(w http.ResponseWriter, req *http.Request) {
	result := MessageModel{}
	result.Success = false

	domainName := req.FormValue("name")
	w.WriteHeader(http.StatusOK)

	if domainName == "" {
		result.Message = "name parameter is null"
		json.NewEncoder(w).Encode(result)
		return
	}

	result = nginx.Delete(domainName, config.Api.ConfigRoot)
	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		log.Fatal(err.Error())
	}

}

func List(w http.ResponseWriter, req *http.Request) {
	result := MessageModelList{}
	result.Success = false

	result = nginx.List(config.Api.ConfigRoot)
	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		log.Fatal(err.Error())
	}

}

type MessageModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type MessageModelList struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Vhosts  []string `json:"vhosts"`
}
