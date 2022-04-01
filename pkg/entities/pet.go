package entities

type Pet struct {
	Id        int64    `json:"id,omitempty"`
	Category  Category `json:"category,omitempty"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags,omitempty"`
	Status    string   `json:"status,omitempty"`
}

func (p *Pet) IsValid() bool {
	return len(p.Name) > 0 && len(p.PhotoUrls) > 0
}

type Category struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Tag struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
