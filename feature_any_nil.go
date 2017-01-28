package jsoniter

type nilAny struct {
	baseAny
}

func (any *nilAny) LastError() error {
	return nil
}

func (any *nilAny) ValueType() ValueType {
	return Nil
}

func (any *nilAny) ToBool() bool {
	return false
}

func (any *nilAny) ToInt() int {
	return 0
}

func (any *nilAny) ToInt32() int32 {
	return 0
}

func (any *nilAny) ToInt64() int64 {
	return 0
}

func (any *nilAny) ToFloat32() float32 {
	return 0
}

func (any *nilAny) ToFloat64() float64 {
	return 0
}

func (any *nilAny) ToString() string {
	return "nil"
}

func (any *nilAny) WriteTo(stream *Stream) {
	stream.WriteNil()
}

func (any *nilAny) Parse() *Iterator {
	return nil
}

func (any *nilAny) GetInterface() interface{} {
	return nil
}