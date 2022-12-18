package main

import (
	"context"
	"fmt"
	"log"

	"github.com/awangga/gopi"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

var apiscope = []string{"https://www.googleapis.com/auth/spreadsheets", "https://www.googleapis.com/auth/documents", "https://www.googleapis.com/auth/drive", "https://www.googleapis.com/auth/blogger", "https://www.googleapis.com/auth/gmail.send", "https://www.googleapis.com/auth/gmail.readonly"}

const jsonsecfile = "credentials.json"
const tokenfile = "token.json"

func main() {
	client := gopi.GetClient(jsonsecfile, tokenfile, apiscope...)
	ctx := context.Background()
	srv, err := docs.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Docs client: %v", err)
	}

	// Prints the title of the requested doc:
	// https://docs.google.com/document/d/195j9eDD3ccgjQRttHhJPymLJUCOUjs-jmwTrekvdjFE/edit
	docId := "195j9eDD3ccgjQRttHhJPymLJUCOUjs-jmwTrekvdjFE"
	doc, err := srv.Documents.Get(docId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from document: %v", err)
	}
	fmt.Printf("The title of the doc is: %s\n", doc.Title)
}
