package handler

import (
	"github.com/abaeve/sde-service/model"
	proto "github.com/abaeve/sde-service/proto"
	"golang.org/x/net/context"
)

var TypeQueryHandler proto.TypeQueryHandler

type typeQueryHandler struct {
	theTypes map[int32]*model.Type
}

func (tqh *typeQueryHandler) FindTypesByTypeIds(context.Context, *proto.TypeIdRequest, *proto.TypeResponse) error {
	return nil
}

func (tqh *typeQueryHandler) FindTypesByTypeNames(context.Context, *proto.TypeNameRequest, *proto.TypeResponse) error {
	return nil
}

func (tqh *typeQueryHandler) SearchForTypes(context.Context, *proto.TypeNameRequest, *proto.TypeNameAndIdResponse) error {
	return nil
}

func (tqh *typeQueryHandler) Load(typeIdsYamlFile string) {

}
