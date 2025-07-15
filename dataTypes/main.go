package main

import "fmt"

/**
	Data Types in Go
	Go has several built-in data types, including:
	- int: Represents integer values.
	- float64: Represents floating-point numbers.
	- string: Represents a sequence of characters.
	- bool: Represents a boolean value (true or false).
	- array: Represents a fixed-size sequence of elements of the same type.
	- slice: Represents a dynamically-sized sequence of elements of the same type.
	- map: Represents a collection of key-value pairs.
**/

func main() {
	// Section 1: Integer Types
	// Section 1.1: Signed Integers
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	i = -42
	i8 = -127
	i16 = -32768
	i32 = 2147483647
	i64 = -9223372036854775807
	fmt.Println("Signed Integers:" + fmt.Sprintf("%d, %d, %d, %d, %d", i, i8, i16, i32, i64))

	// Section 1.2: Unsigned Integers
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	u = 42
	u8 = 255
	u16 = 65535
	u32 = 4294967295
	u64 = 18446744073709551615
	fmt.Println("Unsigned Integers:" + fmt.Sprintf("%d, %d, %d, %d, %d", u, u8, u16, u32, u64))

	// Section 2: Floating-Point Types

	// Section 2.1: Float32
	var f32 float32
	f32 = 3.14
	fmt.Println("Float32:", f32)

	// Section 2.2: Float64
	var f64 float64
	f64 = 3.141592653589793
	fmt.Println("Float64:", f64)

	// difference between the float32 and float64
	var HP float64 = 1.00123456789012345
	var LP = float32(HP)
	fmt.Println("High Precision Float64:", HP)
	fmt.Println("Low Precision Float32:", LP)

	// Section 3: boolean Type
	var isActive bool = false
	fmt.Println("Boolean:", isActive)

	// Section4: Complex numbers
	var c64 complex64 = complex(1, 2)   // 1 + 2i
	var c128 complex128 = complex(1, 2) // 1.0 + 2.0i
	fmt.Println("Complex64:", c64)
	fmt.Println("Complex128:", c128)
}
