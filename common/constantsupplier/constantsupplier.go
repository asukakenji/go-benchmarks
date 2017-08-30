package constantsupplier

// Int is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-2
type Int struct {
	value int
}

// NewInt allocates and returns a new Int.
func NewInt(value int) *Int {
	return &Int{
		value: value,
	}
}

// Next returns the constant value.
func (s *Int) Next() int {
	return s.value
}

// Int8 is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-3
type Int8 struct {
	value int8
}

// NewInt8 allocates and returns a new Int8.
func NewInt8(value int8) *Int8 {
	return &Int8{
		value: value,
	}
}

// Next returns the constant value.
func (s *Int8) Next() int8 {
	return s.value
}

// Int16 is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-4
type Int16 struct {
	value int16
}

// NewInt16 allocates and returns a new Int16.
func NewInt16(value int16) *Int16 {
	return &Int16{
		value: value,
	}
}

// Next returns the constant value.
func (s *Int16) Next() int16 {
	return s.value
}

// Int32 is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-5
type Int32 struct {
	value int32
}

// NewInt32 allocates and returns a new Int32.
func NewInt32(value int32) *Int32 {
	return &Int32{
		value: value,
	}
}

// Next returns the constant value.
func (s *Int32) Next() int32 {
	return s.value
}

// Int64 is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-6
type Int64 struct {
	value int64
}

// NewInt64 allocates and returns a new Int64.
func NewInt64(value int64) *Int64 {
	return &Int64{
		value: value,
	}
}

// Next returns the constant value.
func (s *Int64) Next() int64 {
	return s.value
}

// Uint is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-7
type Uint struct {
	value uint
}

// NewUint allocates and returns a new Uint.
func NewUint(value uint) *Uint {
	return &Uint{
		value: value,
	}
}

// Next returns the constant value.
func (s *Uint) Next() uint {
	return s.value
}

// Uint8 is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-8
type Uint8 struct {
	value uint8
}

// NewUint8 allocates and returns a new Uint8.
func NewUint8(value uint8) *Uint8 {
	return &Uint8{
		value: value,
	}
}

// Next returns the constant value.
func (s *Uint8) Next() uint8 {
	return s.value
}

// Uint16 is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-9
type Uint16 struct {
	value uint16
}

// NewUint16 allocates and returns a new Uint16.
func NewUint16(value uint16) *Uint16 {
	return &Uint16{
		value: value,
	}
}

// Next returns the constant value.
func (s *Uint16) Next() uint16 {
	return s.value
}

// Uint32 is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-10
type Uint32 struct {
	value uint32
}

// NewUint32 allocates and returns a new Uint32.
func NewUint32(value uint32) *Uint32 {
	return &Uint32{
		value: value,
	}
}

// Next returns the constant value.
func (s *Uint32) Next() uint32 {
	return s.value
}

// Uint64 is a type for providing a constant value to be used in test cases
// or benchmarks.
//
// ID: CNG-11
type Uint64 struct {
	value uint64
}

// NewUint64 allocates and returns a new Uint64.
func NewUint64(value uint64) *Uint64 {
	return &Uint64{
		value: value,
	}
}

// Next returns the constant value.
func (s *Uint64) Next() uint64 {
	return s.value
}
