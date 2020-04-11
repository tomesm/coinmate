package coinmate

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/iancoleman/strcase"
)

func createURLValues(body interface{}) (values url.Values) {
	data, err := query.Values(body)
	values = url.Values{}
	if err != nil {
		fmt.Println("ERROR creating request body")
		return nil
	}
	for k, v := range data {
		values.Set(strcase.ToLowerCamel(k), v[0])
	}
	return values
}
