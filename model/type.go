package model

import (
	proto "github.com/abaeve/sde-service/proto"
	"strconv"
)

type Type struct {
	TypeID        string `yaml:"typeID" bson:"typeID"`
	Capacity      string `yaml:"capacity" bson:"capacity"`
	Description   string `yaml:"description" bson:"description"`
	GraphicId     string `yaml:"graphicID" bson:"graphicID"`
	GroupId       string `yaml:"groupID" bson:"groupID"`
	IconId        string `yaml:"iconID" bson:"iconID"`
	MarketGroupId string `yaml:"marketGroupID" bson:"marketGroupID"`
	Mass          string `yaml:"mass" bson:"mass"`
	Name          string `yaml:"name" bson:"name"`
	PortionSize   string `yaml:"portionSize" bson:"portionSize"`
	Published     string `yaml:"published" bson:"published"`
	Radius        string `yaml:"radius" bson:"radius"`
	Volume        string `yaml:"volume" bson:"volume"`
}

func ProtoTypeFromModelType(theType Type, locale proto.Locale) *proto.Type {
	pType := &proto.Type{
		Name:        theType.Name,
		Description: theType.Description,
	}
	//This shit really needs cleaned up but to do that I need to unfuck the importer
	//this is what I get for stealing some code off of stack overflow :P.
	iC, _ := strconv.Atoi(theType.TypeID)
	pType.TypeId = int32(iC)

	f64C, _ := strconv.ParseFloat(theType.Volume, 64)
	pType.Volume = float32(f64C)

	f64C, _ = strconv.ParseFloat(theType.Radius, 64)
	pType.Radius = float32(f64C)

	bC, _ := strconv.ParseBool(theType.Published)
	pType.Published = bC

	iC, _ = strconv.Atoi(theType.PortionSize)
	pType.PortionSize = int32(iC)

	f64C, _ = strconv.ParseFloat(theType.Mass, 64)
	pType.Mass = float32(f64C)

	iC, _ = strconv.Atoi(theType.MarketGroupId)
	pType.MarketGroupId = int32(iC)

	iC, _ = strconv.Atoi(theType.IconId)
	pType.IconId = int32(iC)

	iC, _ = strconv.Atoi(theType.GroupId)
	pType.GroupId = int32(iC)

	iC, _ = strconv.Atoi(theType.GraphicId)
	pType.GraphicId = int32(iC)

	f64C, _ = strconv.ParseFloat(theType.Capacity, 64)
	pType.Capacity = float32(f64C)

	return pType
}

//{
//"_id": "593",
//"radius": "29",
//"volume": "26500",
//"groupID": "25",
//"capacity": "140",
//"description": "Often nicknamed The Fat Man this nimble little frigate is mainly used by the Federation in escort duties or on short-range patrols. The Tristan has been very popular throughout Gallente space for years because of its versatility. It is rather expensive, but buyers will definitely get their money's worth, as the Tristan is one of the more powerful frigates available on the market.",
//"mass": "1044000",
//"portionSize": "1",
//"typeID": "593",
//"soundID": "20074",
//"sofFactionName": "gallentebase",
//"published": "true",
//"marketGroupID": "77",
//"graphicID": "60",
//"name": "Tristan",
//"factionID": "500004",
//"raceID": "8"
//}
