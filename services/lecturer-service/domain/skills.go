package domain

type Skill struct {
	NameMap string `json:"name_map" bson:"name_map"`
	Value   string `json:"value" bson:"value"`
}
