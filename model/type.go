package model

type Type struct {
	BasePrice     float32 `yaml:"basePrice"`
	Capacity      float32 `yaml:"capacity"`
	Description   map[string]string `yaml:"description"`
	GraphicId     int32 `yaml:"graphicID"`
	GroupId       int `yaml:"groupID"`
	IconId        int `yaml:"iconID"`
	MarketGroupId int `yaml:"marketGroupID"`
	Mass          float32 `yaml:"mass"`
	Name          map[string]string `yaml:"name"`
	PortionSize   int `yaml:"portionSize"`
	Published     bool `yaml:"published"`
	Radius        float32 `yaml:"radius"`
	FactionName   string `yaml:"sofFactionName"`
	Volume        float32 `yaml:"volume"`
}

//func ProtoTypeFromModelType(theType Type, locale proto.Locale) proto.Type {
//	return nil
//}
