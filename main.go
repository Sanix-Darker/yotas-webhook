package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func logEvent(e *github.PullRequestEvent) {

	log.Printf("=========================================================== \n")
	if *e.PullRequest.Merged {
		log.Printf("   - PR-ID: %d \n", *e.PullRequest.ID)
		log.Printf("[-] url : '%s'\n", *e.PullRequest.URL)

		log.Printf("   - Merged: %t \n", *e.PullRequest.Merged)
		log.Printf("   - Merged By: %s \n", *e.PullRequest.MergedBy.Login)
	} else {
		log.Printf("-----------------------------------------------------------\n")
		log.Printf("[-] dev : '%s'\n", *e.Sender.Login)
		log.Printf("[-] project : '%s'\n", *e.Repo.FullName)
		log.Printf("[-] title : '%s'\n", *e.PullRequest.Title)
		log.Printf("[-] body : '%s'\n", *e.PullRequest.Body)
		log.Printf("[-] url : '%s'\n", *e.PullRequest.URL)
		log.Printf("-----------------------------------------------------------\n")

		log.Printf("   - PR-ID: %d \n", *e.PullRequest.ID)
		log.Printf("   - commits: %d \n", *e.PullRequest.Commits)
		log.Printf("   - additions: %d \n", *e.PullRequest.Additions)
		log.Printf("   - deletions: %d \n", *e.PullRequest.Deletions)
		log.Printf("   - changedfiles: %d \n", *e.PullRequest.ChangedFiles)
	}
	log.Printf("=========================================================== \n")
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {

	var YOTAS_WEBHOOK_SECRET = os.Getenv("YOTAS_WEBHOOK_SECRET")

	payload, err := github.ValidatePayload(r, []byte(YOTAS_WEBHOOK_SECRET))
	if err != nil {
		log.Printf("[x] error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("[x] could not parse webhook: err=%s\n", err)
		return
	}

	switch e := event.(type) {
	case *github.PullRequestEvent:
		Block{
			Try: func() {
				logEvent(e)
			},
			Catch: func(e Exception) {
				fmt.Println("------------------------------------------")
				fmt.Printf("Caught : %v\n", e)
				fmt.Println("------------------------------------------")
			},
			Finally: func() {
				log.Printf("\n")
			},
		}.Do()
	}
}

func main() {

	log.Println("[+] Yotas-WebHook started...")
	log.Println("[-] Running on http: http://127.0.0.1:9091/webhook")

	http.HandleFunc("/webhook", handleWebhook)

	log.Fatal(http.ListenAndServe(":9091", nil))
}
