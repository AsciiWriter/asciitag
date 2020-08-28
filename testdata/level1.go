package tests

func MyFirstFunction() {
	var a = 1
	a++
}

// TAG Is it really your second function?
func MySecondFunction() {
	// This is a dummy comment
	// TAG This is a tabbed TAG
	// TAG this a second tag in the same comment group
	
	// This is another dummy comment
	var b = 2
	b++
}

// FIXME Please delete me!
var DeleteMe = 5

// tag::mymethod[]
// TAG(asticode) I should be false
var Oops = true

// TAG(asciitag) Something else comes here
var SomethingElse = "Something else"

// TAG: I can use colons to signal the tag.
var WithColons = "Something else"

// TAG(asciitag): It also works with assignee.
var WithColons2 = "Something else"
