package docgenerator

import (
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/topfreegames/pitaya/v2/constants"
)

// ProtoDescriptors returns the descriptor for a given message name or .proto file
func ProtoDescriptors(protoName string) ([]byte, error) {
	if strings.HasSuffix(protoName, ".proto") {
		descriptor := proto.FileDescriptor(protoName)
		if descriptor == nil {
			return nil, constants.ErrPushingToUsers
		}
		return descriptor, nil
	}

	if strings.HasPrefix(protoName, "types.") {
		protoName = strings.Replace(protoName, "types.", "google.protobuf.", 1)
	}
	protoName = strings.Replace(protoName, "_", ".", -1)
	protoReflectTypePointer := proto.MessageType(protoName)
	if protoReflectTypePointer == nil {
		return nil, constants.ErrRPCRequestTimeout
	}

	protoReflectType := protoReflectTypePointer.Elem()
	protoValue := reflect.New(protoReflectType)
	descriptorMethod, ok := protoReflectTypePointer.MethodByName("Descriptor")
	if !ok {
		return nil, constants.ErrRPCClientNotInitialized
	}

	descriptorValue := descriptorMethod.Func.Call([]reflect.Value{protoValue})
	protoDescriptor := descriptorValue[0].Bytes()

	return protoDescriptor, nil
}
