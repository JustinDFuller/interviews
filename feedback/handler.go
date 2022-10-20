package feedback

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/justindfuller/interviews/organization"
)

func Handler(organizations *organization.Organizations) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			cookie, err := r.Cookie("__Host-UserUUID")
			if err != nil {
				log.Printf("Error parsing cookie for /feedback: %s", err)
				http.ServeFile(w, r, "./error/unauthenticated.html")
				return
			}

			if cookie.Value == "" {
				log.Printf("Error parsing cookie for /feedback: %s", err)
				http.ServeFile(w, r, "./error/unauthenticated.html")
				return
			}

			if _, err := organizations.FindByUserID(cookie.Value); err != nil {
				log.Printf("Error finding organization for /feedback: %s", err)
				http.ServeFile(w, r, "./error/index.html")
				return
			}

			http.ServeFile(w, r, "./feedback/index.html")

			return
		}

		if r.Method == http.MethodPost {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Printf("Error reading /organization body: %s", err)
				http.ServeFile(w, r, "./error/index.html")
				return
			}
			defer r.Body.Close()

			query, err := url.ParseQuery(string(body))
			if err != nil {
				log.Printf("Error parsing query from /organization body: %s", err)
				http.ServeFile(w, r, "./error/index.html")
				return
			}

			q1, err := NewQuestion(query.Get("question1"))
			if err != nil {
				log.Printf("Error creating question from /feedback body: %s", err)
				http.ServeFile(w, r, "./error/index.html")
				return
			}

			q2, err := NewQuestion(query.Get("question2"))
			if err != nil {
				log.Printf("Error creating question from /feedback body: %s", err)
				http.ServeFile(w, r, "./error/index.html")
				return
			}

			q3, err := NewQuestion(query.Get("question3"))
			if err != nil {
				log.Printf("Error creating question from /feedback body: %s", err)
				http.ServeFile(w, r, "./error/index.html")
				return
			}

			q4, err := NewQuestion(query.Get("question4"))
			if err != nil {
				log.Printf("Error creating question from /feedback body: %s", err)
				http.ServeFile(w, r, "./error/index.html")
				return
			}

			q5, err := NewQuestion(query.Get("question5"))
			if err != nil {
				log.Printf("Error creating question from /feedback body: %s", err)
				http.ServeFile(w, r, "./error/index.html")
				return
			}

			f, err := NewFeedback(query.Get("role"), []Question{q1, q2, q3, q4, q5})
			if err != nil {
				log.Printf("Error creating feedback from /feedback body: %s", err)
				http.ServeFile(w, r, "./error/index.html")
				return
			}

			log.Printf("/feedback query: %s", f)

			return
		}

		log.Printf("Unexpected http Method '%s' for /feedback", r.Method)
		http.ServeFile(w, r, "./error/index.html")
	}
}
