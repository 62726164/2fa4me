package main

import (
	"encoding/xml"
	"net/http"
)

type Play struct {
	Digits string `xml:"digits,attr"`
}

type Response struct {
	Play Play `xml:"Play"`
}

func mfa(w http.ResponseWriter, r *http.Request) {
	// w means wait for 0.5 seconds (ww is 1 second).
	// If you wanted to wait 2 seconds and then press 9, change this to "wwww9".
	play := Play{Digits: "w1w2w3w4"}
	resp := Response{Play: play}
	x, err := xml.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func main() {
	// Todo: Make this a server struct and set timeouts
	http.HandleFunc("/mfa", mfa)
	http.ListenAndServe(":3000", nil)
}
