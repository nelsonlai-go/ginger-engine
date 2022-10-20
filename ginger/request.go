package ginger

import "reflect"

func GetRequest[T any](ctx Context) *T {
	request := new(T)
	tags := parseRequestTags(request)

	for _, key := range tags {
		switch key {
		case "uri":
			if err := ctx.ShouldBindUri(request); err != nil {
				panic(err)
			}
		case "json":
			if err := ctx.ShouldBindJSON(request); err != nil {
				panic(err)
			}
		case "form":
			if err := ctx.ShouldBindQuery(request); err != nil {
				panic(err)
			}
		}
	}

	return request
}

func parseRequestTags[T any](request T) []string {
	m := make(map[string]bool)

	numOfField := reflect.TypeOf(request).Elem().NumField()
	for i := 0; i < numOfField; i++ {
		tag := reflect.TypeOf(request).Elem().Field(i).Tag
		for _, key := range []string{"uri", "json", "form"} {
			check := tag.Get(key)
			if check != "" {
				m[key] = true
				break
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
