package slice

import "github.com/clipperhouse/typewriter"

var indexByT = &typewriter.Template{
	Name: "IndexBy",
	Text: `
// IndexBy{{.TypeParameter.LongName}} groups elements into a map keyed by {{.TypeParameter}}. Returns an error if more than one element satisfies the func. See: http://clipperhouse.github.io/gen/#IndexBy
func (rcv {{.SliceName}}) IndexBy{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (map[{{.TypeParameter}}]{{.Type}}, error) {
	result := make(map[{{.TypeParameter}}]{{.Type}})
	for _, v := range rcv {
		key := fn(v)
		if _, ok := result[key]; ok {
			return nil, errors.New("multiple {{.SliceName}} elements returned true for passed func")
		}
		result[key] = v
	}
	return result, nil
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, and it must be comparable
		{Comparable: true},
	},
}
