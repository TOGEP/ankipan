package graphql

import (
  "github.com/graphql-go/graphql"
)

var cardtype = graphql.NewObject(
  graphql.ObjectConfig {
    Name: "Card",
    Fields: graphql.Fields {
      "problem": &graphql.Field {
        Type: graphql.String,
      },
      "anser": &graphql.Field{
        Type: graphql.String,
      },
      "note": &graphql.Field{
        Type: graphql.String,
      },
    },
  },
)
