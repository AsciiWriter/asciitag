package level1

var dumdum int /*
	First line of comment not part of the tag
	TAG Here is a
	multi line tag
*/
func MyFirstFunction() {
	var a = 1
	a++
	// tag(my.weird-email_address+1@email.com) This is a named TAG
	var c = 3
	c++
}


// tag::mymethod[]
// multi line tag
func MySecondFunction() {
	var b = 2
	b++
}
// tag(test) something
