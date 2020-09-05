package graphql

var Schema = `
	schema {
		query: Query
	}

	# All Query api
	type Query {
		hello: String!
	}
`

type Resolver struct{}

func (_ *Resolver) Hello() string { return "Hello, world!" }
