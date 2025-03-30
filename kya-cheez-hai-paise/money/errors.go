package money

// NOTE:
// Question: Are we able to create a chain
// such that an explicit typecast from top to bottom is possible?
type Error string

// NOTE:
// 1. Any type that has this method
// is an error
func (e Error) Error() string {
	return string(e)
}
