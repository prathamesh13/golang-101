package main

import "fmt"

func main() {

	// Refer https://design215.com/toolbox/ascii-utf8.php
	// or
	// https://utf8-chartable.de/unicode-utf8-table.pl
	// for character codes.

	fmt.Println("Exploring runes, bytes and encoding in go.")

	// ASCII characters. represented using 0 to 127

	// ASCII code for space is 32
	spaceCode := 32

	// Prints Hello There.
	fmt.Printf("Hello%cThere.\n", spaceCode)

	fmt.Printf("Data type of variable spaceCode is %T\n", spaceCode)

	var spaceCodeInByte byte = 32

	// Even though spaceCodeInByte is declared with data type byte, when printed it would
	// indicate that the datatype is uint8. This is because byte is and alias for uint8.
	fmt.Printf("Hello%cThere printed using spaceCodeInByte of type %T\n", spaceCodeInByte, spaceCodeInByte)

	// Declaring and assigning a character variable

	var dollarSymbol = '$'
	var pipeCharacter = '|'

	// character variable is defined with int32. rune is alias for int32
	fmt.Printf("Data Type of dollarSymbol %c is %T\n", dollarSymbol, dollarSymbol)
	fmt.Printf("Data Type of pipeCharacter %c is %T\n", pipeCharacter, pipeCharacter)

	// explicitly the variables can also be declared as byte
	var dollarSymbolInByte byte = '$'
	var pipeCharacterInByte byte = '|'

	fmt.Printf("Data Type of dollarSymbol %c is %T\n", dollarSymbolInByte, dollarSymbolInByte)
	fmt.Printf("Data Type of dollarSymbol %c is %T\n", pipeCharacterInByte, pipeCharacterInByte)

	// try to define a character outside of ASCII character set to a byte
	// var outOfRangeCharacter byte = 256 // not allowed

	// Either of following is allowed.
	var outOfRangeCharacter = 256 // declared with data type int implicitly

	var anotherOutOfRangeCharacter rune = 257

	fmt.Printf("Printed %c to console with variable outOfRangeCharacter of type %T\n", outOfRangeCharacter, outOfRangeCharacter)
	fmt.Printf("Printed %c to console with variable outOfRangeCharacter of type %T\n", anotherOutOfRangeCharacter, anotherOutOfRangeCharacter)

	// declare character variable within ascii range and outside and note data types assigned.
	var pilCrow = '¶'     // ASCII code: 182
	var caret = '^'       // ASCII code: 94
	var copyright = '©'   // non ascii character. 2 byte utf-8 character
	var _3lookalike = 'Ʒ' // non ascii character. 2 byte utf-8 character. 439
	// var cappedA = 'ȃ'     // non ascii character. 2 byte utf-8 character. 439

	fmt.Printf("Printed %c to console using character variable of type %T ::  decimal -> %v, hex -> %x\n", pilCrow, pilCrow, pilCrow, pilCrow)
	fmt.Printf("Printed %c to console using character variable of type %T ::  decimal -> %v, hex -> %x\n", caret, caret, caret, caret)
	fmt.Printf("Printed %c to console using character variable of type %T ::  decimal -> %v, hex -> %x\n", copyright, copyright, copyright, copyright)
	fmt.Printf("Printed %c to console using character variable of type %T ::  decimal -> %v, hex -> %x\n", _3lookalike, _3lookalike, _3lookalike, _3lookalike)

	// 3 byte character
	// var sh byte = 'ष' - not allowed since it doesn't fit in the range for byte.
	sh := 'ष'
	fmt.Printf("Printed %c to console using character variable of type %T ::  decimal -> %v, hex -> %x\n", sh, sh, sh, sh)

	// 4 byte unicode character
	unicorn := '🦄'
	fmt.Printf("Printed %c to console using character variable of type %T ::  decimal -> %v, hex -> %x\n", unicorn, unicorn, unicorn, unicorn)

	fmt.Printf("Unicorn %c unicode character's numeric representation (decimal) is %d\n", unicorn, unicorn)
	fmt.Printf("Unicode %c unicode character's numeric representation(binary) is %b\n", unicorn, unicorn)
	fmt.Printf("Unicode %c unicode character's numeric representation(hex) is %x\n", unicorn, unicorn)
	fmt.Printf("Unicode %c unicode character's numeric representation(octal) is %o\n", unicorn, unicorn)
}
