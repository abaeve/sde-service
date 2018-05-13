package handler

import (
	"fmt"
	"github.com/abaeve/sde-service/model"
	proto "github.com/abaeve/sde-service/proto"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type typeQueryHandler struct {
	s *mgo.Session
}

func (tqh *typeQueryHandler) FindTypesByTypeIds(context.Context, *proto.TypeIdRequest, *proto.TypeResponse) error {
	return nil
}

func (tqh *typeQueryHandler) SearchTypesByTypeNames(ctx context.Context, req *proto.TypeNameSearchRequest, rsp *proto.TypesResponse) error {
	//{ $or: [{ "name":  { $regex: /150mm railgun.*/i } }, { "name":  { $regex: /250mm rail.*/i } }] }
	s := tqh.s.Copy()
	defer s.Close()

	c := s.DB("sde").C("types")

	var results []model.Type

	var searches []bson.M

	for _, typeName := range req.TypeName {
		searches = append(searches, bson.M{
			"name": bson.RegEx{Pattern: "^" + typeName + ".*", Options: "i"},
		})
	}

	c.Find(
		bson.M{
			"$or":       searches,
			"published": "true",
		},
	).All(&results)

	fmt.Printf("%v\n", results)

	for _, result := range results {
		rsp.Type = append(rsp.Type, model.ProtoTypeFromModelType(result, req.Locale))
	}

	return nil
}

func (tqh *typeQueryHandler) FindTypeByTypeName(ctx context.Context, req *proto.TypeNameRequest, rsp *proto.TypeResponse) error {
	//{ $or: [{ "name":  { $regex: /150mm railgun.*/i } }, { "name":  { $regex: /250mm rail.*/i } }] }
	s := tqh.s.Copy()
	defer s.Close()

	c := s.DB("sde").C("types")

	var result model.Type

	c.Find(
		bson.M{
			"name":      req.TypeName,
			"published": "true",
		},
	).One(&result)

	rsp.Type = model.ProtoTypeFromModelType(result, req.Locale)

	return nil
}

func (tqh *typeQueryHandler) SearchForTypes(context.Context, *proto.TypeNameRequest, *proto.TypeNameAndIdResponse) error {
	return nil
}

func (tqh *typeQueryHandler) Load(typeIdsYamlFile string) {

}

func NewTypeQueryHandler(session *mgo.Session) proto.TypeQueryHandler {
	return &typeQueryHandler{
		s: session,
	}
}
