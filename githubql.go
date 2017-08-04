package main

import (
	"context"
	"log"

	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	client := githubql.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "81b30e44284fb6a0ddf2ee74f118434c7eead65b"},
	)))

	var query struct {
		Search struct {
			Nodes []struct {
				Issue struct {
					Number     githubql.Int
					Repository struct {
						NameWithOwner githubql.String
					}
				} `graphql:"... on Issue"`
			}
		} `graphql:"search(first: 100, query: $searchQuery, type: $searchType)"`
	}
	variables := map[string]interface{}{
		"searchQuery": githubql.String("repo:cockroachdb/cockroach state:open teamcity: failed tests on master"),
		"searchType":  githubql.SearchTypeIssue,
	}

	log.Printf("err=%v", client.Query(ctx, &query, variables))
}
