// Code generated by "stringer -type=HashAlgorithm"; DO NOT EDIT.

package sema

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[HashAlgorithmUnknown-0]
	_ = x[HashAlgorithmSHA2_256-1]
	_ = x[HashAlgorithmSHA2_384-2]
	_ = x[HashAlgorithmSHA3_256-3]
	_ = x[HashAlgorithmSHA3_384-4]
}

const _HashAlgorithm_name = "HashAlgorithmUnknownHashAlgorithmSHA2_256HashAlgorithmSHA2_384HashAlgorithmSHA3_256HashAlgorithmSHA3_384"

var _HashAlgorithm_index = [...]uint8{0, 20, 41, 62, 83, 104}

func (i HashAlgorithm) String() string {
	if i >= HashAlgorithm(len(_HashAlgorithm_index)-1) {
		return "HashAlgorithm(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _HashAlgorithm_name[_HashAlgorithm_index[i]:_HashAlgorithm_index[i+1]]
}
