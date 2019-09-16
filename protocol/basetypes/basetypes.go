package basetypes

import "fmt"

// JDWPString represents string in JWDP wire format
type JDWPString struct {
	Length     uint32
	ByteString []byte `struct:"sizefrom=Length"`
}

func (j *JDWPString) String() string {
	return (string)(j.ByteString)
}

// EmptyJWDPString returns the empty string
func EmptyJWDPString() JDWPString {
	return JDWPString{
		Length:     0,
		ByteString: make([]byte, 0),
	}
}

// TODO we need to extend the serialiser to allow these sizes
// to be changed at runtime (IDSizes command)

// JWDPObjectID represents objectID
type JWDPObjectID struct {
	ObjectID uint64
}

func (j *JWDPObjectID) String() string {
	return fmt.Sprintf("0x%X", j.ObjectID)
}

// JWDPFrameID represents frameID
type JWDPFrameID struct {
	FrameID uint64
}

func (j *JWDPFrameID) String() string {
	return fmt.Sprintf("0x%X", j.FrameID)
}

// JWDPFieldID represents fieldID
type JWDPFieldID struct {
	FieldID uint64
}

func (j *JWDPFieldID) String() string {
	return fmt.Sprintf("0x%X", j.FieldID)
}

// JWDPRefTypeID represents refTypeID
type JWDPRefTypeID struct {
	RefTypeID uint64
}

func (j *JWDPRefTypeID) String() string {
	return fmt.Sprintf("0x%X", j.RefTypeID)
}

// JWDPMethodID respresents methodID
type JWDPMethodID struct {
	MethodID uint64
}

func (j *JWDPMethodID) String() string {
	return fmt.Sprintf("0x%X", j.MethodID)
}

// JWDPTypeTag represents type tag
type JWDPTypeTag byte

const (
	// JWDPTypeTagClass - class
	JWDPTypeTagClass = 1
	// JWDPTypeTagInterface - interface
	JWDPTypeTagInterface = 2
	// JWDPTypeTagArray - array
	JWDPTypeTagArray = 3
)

func (j JWDPTypeTag) String() string {
	switch j {
	case JWDPTypeTagClass:
		return "Class"
	case JWDPTypeTagInterface:
		return "Interface"
	case JWDPTypeTagArray:
		return "Array"
	default:
		return "Unknown"
	}
}
