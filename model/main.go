package model

type Definition struct {
	Name       string             `json:name`
	Adapter    string             `json:adapter`
	Properties map[string]PropDef `json:properties`
}

type PropDef struct {
	Type     string `json:type`
	Required bool   `json:required`
}
