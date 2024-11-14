package test

import (
	"fmt"
	"github.com/hdt3213/godis/lib/logger"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

const (
	BITS_PER_BYTE = 8
	COMPLEMENT_8  = 0xFF
	COMPLEMENT_64 = 0xFFFFFFFFFFFFFFFF
)

type ldEnum interface {
	getString() string
	CheckValid() bool
}

type STYPE uint8

const (
	AUC_RQST     STYPE = 0x41
	AUC_RESP     STYPE = 0x42
	AUC_KEY_EXEC STYPE = 0x43
)

func (f STYPE) getString() string {
	return [...]string{
		"AUC_RQST",
		"AUC_RESP",
		"AUC_KEY_EXEC",
	}[f-AUC_RQST]
}

func (f STYPE) CheckValid() bool {
	return f >= AUC_RQST && f <= AUC_KEY_EXEC
}

type PID uint8

const (
	PID_RESERVED PID = 0x0
	PID_SIGN     PID = 0x1
	PID_MAC      PID = 0x2
	PID_BOTH     PID = 0x3
)

func (f PID) getString() string {
	return [...]string{
		"PID_RESERVED",
		"PID_SIGN",
		"PID_MAC",
		"PID_BOTH",
	}[f-PID_RESERVED]
}

func (f PID) CheckValid() bool {
	return f <= PID_BOTH
}

type AucRqst struct {
	Stype STYPE  `ldacs:"name:S_TYPE; size:8; type:enum"`
	Ver   uint8  `ldacs:"name:VER; size:3; type:set"`
	PID   PID    `ldacs:"name:PID; size:2; type:enum"`
	N1    []byte `ldacs:"name:N1; bytes_size: 3; type:fbytes"`
	//N1 []byte `ldacs:"name:N1; type:dbytes"`
	//ASSac  uint16 `ldacs:"required;minlength:8"`
	//GSSac  uint16 `ldacs:"required;minlength:8"`
	//MacLen uint8  `ldacs:"required;minlength:8"`
	//AuthID uint8  `ldacs:"required;minlength:8"`
	//EncID  uint8  `ldacs:"required;minlength:8"`
}

type encodePkt struct {
	bytes    []byte
	currByte uint64
}

type structEncoder struct {
	fields structFields
}

// A Field represents a single Field found in a struct.
type Field struct {
	name      string
	nameBytes []byte // []byte(name)

	size      uint
	bytesSize uint
	typ       reflect.Type
	ltype     string
}

type structFields struct {
	list []Field
}

type encoderFunc func(e *encodePkt, v reflect.Value)

func typeFields(t reflect.Type) structFields {
	var fields []Field
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)

		tag := fieldType.Tag.Get("ldacs")
		if tag == "" {
			continue
		}

		kv, _, err := splitRules(tag)
		if err != nil {
			break
		}
		field := Field{
			name:      kv["name"],
			size:      uint(Atoi(kv["size"])),
			bytesSize: uint(Atoi(kv["bytes_size"])),
			typ:       fieldType.Type,
			ltype:     kv["type"],
		}
		field.nameBytes = []byte(field.name)

		fields = append(fields, field)
	}

	return structFields{fields}
}

func (e *encodePkt) appendBits(toShift *int, n uint64) {
	for *toShift > (^BITS_PER_BYTE + 1) {
		if *toShift >= 0 {
			e.bytes[e.currByte] |= uint8((n >> uint(*toShift)) & COMPLEMENT_8)
			e.currByte++
		} else {
			e.bytes[e.currByte] |= uint8((n << (^(*toShift) + 1)) & COMPLEMENT_8)
		}
		*toShift -= BITS_PER_BYTE
	}
}

func (e *encodePkt) detachBits(toShift *int, n *uint64, preUnavil int) {

	var specific uint8

	for *toShift > (^BITS_PER_BYTE + 1) {
		if preUnavil != 0 {
			specific = e.bytes[e.currByte] & (0xFF >> preUnavil)
			preUnavil = 0
		} else {
			specific = e.bytes[e.currByte]
		}

		if *toShift >= 0 {
			*n |= (uint64(specific) << *toShift) & COMPLEMENT_64
			e.currByte++
		} else {
			*n |= (uint64(specific) >> (^(*toShift) + 1)) & COMPLEMENT_64
		}
		*toShift -= BITS_PER_BYTE
	}
}

// An UnsupportedValueError is returned by Marshal when attempting
// to encode an unsupported value.
type UnsupportedValueError struct {
	Value reflect.Value
	Str   string
}

func (e *UnsupportedValueError) Error() string {
	return "json: unsupported value: " + e.Str
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "ldacs: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "ldacs: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "ldacs: Unmarshal(nil " + e.Type.String() + ")"
}

func (se structEncoder) encode(e *encodePkt, v reflect.Value) error {
	var bitsize = 0

	for i := 0; i < v.NumField(); i++ {
		//fmt.Println(se.fields.list[i], v.Field(i))
		f := &se.fields.list[i]
		fv := v.Field(i)
		switch f.ltype {
		case "enum", "set":
			sz := int(f.size)
			var n = fv.Uint()

			toShift := sz - (BITS_PER_BYTE - (bitsize % BITS_PER_BYTE))

			switch f.ltype {
			case "enum":
				if fv.Interface() != nil {
					if ldenum, ok := fv.Interface().(ldEnum); ok {
						//fmt.Println(ldenum.CheckValid(), ldenum, fv.Interface())
						if ldenum.CheckValid() == false {
							return &UnsupportedValueError{v, ""}
						}
					}
				}

			default:

			}

			e.appendBits(&toShift, n)
			bitsize += sz

		case "fbytes":
			sz := int(f.bytesSize)
			if bitsize%BITS_PER_BYTE != 0 {
				// Pad
				e.currByte++
			}
			copy(e.bytes[e.currByte:], fv.Bytes()[:sz])
		case "dbytes":
			if bitsize%BITS_PER_BYTE != 0 {
				// Pad
				e.currByte++
			}
			copy(e.bytes[e.currByte:], fv.Bytes()[:])
		default:

		}
	}

	return nil
}

func (se structEncoder) decode(e *encodePkt, v reflect.Value) error {

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	var bitsize = 0

	for i := 0; i < v.NumField(); i++ {
		f := &se.fields.list[i]
		fv := v.Field(i)
		switch f.ltype {
		case "enum", "set":
			sz := int(f.size)
			var n uint64 = 0
			preOctectUnavil := bitsize % BITS_PER_BYTE
			toShift := sz + preOctectUnavil - BITS_PER_BYTE
			e.detachBits(&toShift, &n, preOctectUnavil)

			fv.SetUint(n)

			switch f.ltype {
			case "enum":
				if fv.Interface() != nil {
					if ldenum, ok := fv.Interface().(ldEnum); ok {
						//fmt.Println(ldenum.CheckValid(), ldenum, fv.Interface())
						if ldenum.CheckValid() == false {
							return &UnsupportedValueError{v, ""}
						}
					} else {
						logger.Warn(f.name, "is NOT enum =>", fv.Type())
					}
				}

			default:

			}
			bitsize += sz
		case "fbytes":
			sz := uint64(f.bytesSize)
			if bitsize%BITS_PER_BYTE != 0 {
				// Pad
				e.currByte++
			}
			fv.SetBytes(e.bytes[e.currByte : e.currByte+sz])
		case "dbytes":
			sz := uint64(len(fv.Bytes()))
			logger.Warn(sz)
			if bitsize%BITS_PER_BYTE != 0 {
				// Pad
				e.currByte++
			}
			fv.SetBytes(e.bytes[e.currByte : e.currByte+sz])
		default:

		}
	}

	return nil
}

func (e *encodePkt) marshal(v any) error {
	rv := reflect.ValueOf(v)
	se := structEncoder{fields: typeFields(rv.Type())}
	return se.encode(e, rv)
}

func (e *encodePkt) unmarshal(v any) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	se := structEncoder{fields: typeFields(rv.Type())}
	return se.decode(e, rv)

}

func MarshalLdacsPkt(v any) ([]byte, error) {
	e := encodePkt{
		bytes:    make([]byte, 512),
		currByte: 0,
	}

	err := e.marshal(v)
	if err != nil {
		return nil, err
	}

	return e.bytes, nil
}

func UnmarshalLdacsPkt(data []byte, v any) error {
	e := encodePkt{
		bytes:    data,
		currByte: 0,
	}

	return e.unmarshal(v)
}

func TestTag(t *testing.T) {
	rqst := AucRqst{
		Stype: 0x41,
		Ver:   0x1,
		PID:   0x1,
		N1:    []byte{0x01, 0x02, 0x03},
	}
	bytes, err := MarshalLdacsPkt(rqst)

	if err != nil {
		logger.Error("Validation failed:", err)
	} else {
		logger.Info("Validation succeeded")
	}

	logger.Warn("Marshaled:", bytes)

	rqst2 := AucRqst{
		N1: make([]byte, 3),
	}
	err = UnmarshalLdacsPkt(bytes, &rqst2)
	if err != nil {
		logger.Error(err)
	}

	logger.Warn("Unmarshaled: ", rqst2)
}

func Atoi(in string) (out int) {
	out, _ = strconv.Atoi(in)
	return
}

func TrimSpace(str string) string {
	return strings.TrimSpace(str)
}

// splitRules 函数用于拆分标签中的规则
func splitRules(tag string) (kvPairs map[string]string, booltags []string, err error) {
	basic_rules := SplitString(tag, ";")

	kvPairs = make(map[string]string)
	booltags = []string{}

	for _, rule := range basic_rules {
		trimmed := TrimSpace(rule) //去除首尾空格
		if trimmed == "" {
			continue
		}

		//检查键值对类型还是bool类型
		if strings.Contains(trimmed, ":") {
			kv := strings.SplitN(trimmed, ":", 2)
			if len(kv) != 2 {
				return nil, nil, fmt.Errorf("invalid key-value pair: %s", trimmed)
			}
			kvPairs[TrimSpace(kv[0])] = TrimSpace(kv[1])
		} else {
			booltags = append(booltags, trimmed)
		}
	}

	return kvPairs, booltags, nil
}

func SplitString(tag, sep string) []string {
	return strings.Split(tag, sep)
}
