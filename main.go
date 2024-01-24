package main

import (
	"log"

	"gihtub.com/task4233/protoc/generated"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GetOption(x generated.Color) *generated.ColorOptions {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic with %v", err)
		}
	}()

	return proto.GetExtension(
		x.Descriptor().Values().Get(int(x.Number())).Options().(*descriptorpb.EnumValueOptions),
		generated.E_ColorOptions,
	).(*generated.ColorOptions)
}

func main() {
	log.Printf("optional_value: %s", GetOption(generated.Color_BLACK).GetColorCode())
}
