package band

import "strconv"

type band struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

var bandsList = []band{
	{ID: 1, Name: "Katatonia"},
	{ID: 2, Name: "Opeth"},
}

func bands() []band {
	return bandsList
}

func loadBands() map[string]band {
	bands := bands()
	res := make(map[string]band, len(bands))

	for _, x := range bands {
		res[strconv.Itoa(int(x.ID))] = x
	}

	return res
}
