package todo

import (
	"io"

	"github.com/ymotongpoo/goltsv"
)

type Decoder struct {
	reader io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{reader: r}
}

func (d *Decoder) Decode() ([]Todo, error) {
	reader := goltsv.NewReader(d.reader)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	todos := []Todo{}
	for _, record := range records {
		var id, parentID, title string
		var done bool

		for k, v := range record {
			switch k {
			case "id":
				id = v
			case "parent_id":
				parentID = v
			case "title":
				title = v
			case "done":
				done = (v == "true")
			}
		}

		todo := Todo{
			ID:       id,
			ParentID: parentID,
			Title:    title,
			Done:     done,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
