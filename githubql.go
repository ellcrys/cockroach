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
			First githubql.Int
			Query githubql.String
			Type  githubql.SearchType
		}
		Nodes []struct {
			Issue struct {
				Number githubql.Int
			}
		}
	}
	query.Search.First = 100
	query.Search.Query = "repo:cockroachdb/cockroach state:open teamcity: failed tests on master"
	query.Search.Type = githubql.Issue

	log.Printf("err=%v", client.Query(ctx, &query, nil))
}
