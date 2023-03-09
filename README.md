# enumstr
`enumstr` provides an easy way to map between enumerated values and their string representations.

### Installation
```shell
go get github.com/semichkin-gopkg/enumstr
```

### Usage
```go
package main

import (
	"github.com/semichkin-gopkg/enumstr"
	"log"
)

type (
	UserStatusFromProtobuf int32
	UserStatus             string
)

const (
	ProtoUserStatus_Unknown    UserStatusFromProtobuf = 0
	ProtoUserStatus_Active     UserStatusFromProtobuf = 1
	ProtoUserStatus_Inactive   UserStatusFromProtobuf = 2
	ProtoUserStatus_MovedToBin UserStatusFromProtobuf = 3
)

var UserStatusFromProtobuf_name = map[int32]string{
	0: "UNKNOWN",
	1: "ACTIVE",
	2: "INACTIVE",
	3: "MOVED_TO_BIN",
}

func main() {
	mapper := enumstr.New[UserStatusFromProtobuf, UserStatus](
		UserStatusFromProtobuf_name,
	)

	log.Println(mapper.ToStr(ProtoUserStatus_Unknown))    // unknown
	log.Println(mapper.ToStr(ProtoUserStatus_Active))     // active
	log.Println(mapper.ToStr(ProtoUserStatus_Inactive))   // inactive
	log.Println(mapper.ToStr(ProtoUserStatus_MovedToBin)) // moved_to_bin
}
```