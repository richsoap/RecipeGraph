// Code generated by "stringer -type errorCode"; DO NOT EDIT.

package errors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NOT_INITIALIZED-0]
}

const _errorCode_name = "NOT_INITIALIZED"

var _errorCode_index = [...]uint8{0, 15}

func (i errorCode) String() string {
	if i >= errorCode(len(_errorCode_index)-1) {
		return "errorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _errorCode_name[_errorCode_index[i]:_errorCode_index[i+1]]
}
