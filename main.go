package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	if err := main1(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main1() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	token := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: os.Getenv("GITHUB_TOKEN"),
	})
	hc := oauth2.NewClient(ctx, token)
	c := githubv4.NewClient(hc)

	var printedHeader bool
	variables := map[string]interface{}{
		"cursor": (*githubv4.String)(nil),
	}
	for {
		var query struct {
			Viewer struct {
				Login        githubv4.String
				Repositories struct {
					PageInfo struct {
						EndCursor   githubv4.String
						HasNextPage githubv4.Boolean
					}
					Nodes []struct {
						URL     githubv4.URI
						License *struct {
							Typename githubv4.String `graphql:"__typename"`
						} `graphql:"license: object(expression: \"HEAD:LICENSE\")"`
						LicenseMd *struct {
							Typename githubv4.String `graphql:"__typename"`
						} `graphql:"licensemd: object(expression: \"HEAD:LICENSE.md\")"`
					}
				} `graphql:"repositories(first: 100, after: $cursor, privacy: PUBLIC, affiliations: OWNER, orderBy: {field: NAME, direction: ASC})"`
			}
		}
		if err := c.Query(ctx, &query, variables); err != nil {
			return err
		}

		if !printedHeader {
			fmt.Fprintf(os.Stderr, "The following public repositories of %s are missing a LICENSE file:\n", query.Viewer.Login)
			printedHeader = true
		}

		for _, node := range query.Viewer.Repositories.Nodes {
			if node.License == nil && node.LicenseMd == nil {
				fmt.Println(node.URL)
			}
		}

		if !query.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = githubv4.NewString(query.Viewer.Repositories.PageInfo.EndCursor)
	}

	return nil
}
