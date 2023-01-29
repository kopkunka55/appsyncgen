package schema

import (
	"log"
	"strings"

	"github.com/kopkunka55/appsyncgen/codegen/directive"
)

type ObjectType struct {
	Name string
}

func (o ObjectType) IsConnection() bool {
	return o.Name == "Connection"
}

func (o ObjectType) IsPayload() bool {
	return o.Name == "Payload"
}

func (o ObjectType) IsConnectionPayload() bool {
	return o.Name == "ConnectionPayload"
}

func NewObjectType(name string) ObjectType {
	if name == "Input" || name == "Mutation" || name == "Query" || name == "Subscription" {
		return ObjectType{name}
	}
	if strings.Contains(name, "Connection") {
		if strings.HasSuffix(name, "Payload") {
			return ObjectType{"ConnectionPayload"}
		} else {
			return ObjectType{"Connection"}
		}
	}
	if strings.HasSuffix(name, "Payload") {
		return ObjectType{"Payload"}
	}
	return ObjectType{"Node"}
}

type Object struct {
	Name      string                  `json:"name"`
	AuthRules *directive.AuthRuleList `json:"-"`
	Fields    FieldList               `json:"-"`
	Type      ObjectType              `json:"type"`
}

func (o Object) HasAuthRule() bool {
	if o.AuthRules == nil {
		return false
	}
	return len(*o.AuthRules) > 0
}

func (o Object) ConnectedNames() (string, string) {
	if o.Type.Name != "Connection" {
		log.Fatalln("given object is not Connection type, but got invoked as Connection type")
	}
	woConnection := o.Name[:len(o.Name)-len("Connection")]
	keys := strings.Split(woConnection, "_")
	return keys[0], keys[1]
}

type ObjectList []*Object

func (o ObjectList) ForName(name string) *Object {
	for _, it := range o {
		if it.Name == name {
			return it
		}
	}
	return nil
}
