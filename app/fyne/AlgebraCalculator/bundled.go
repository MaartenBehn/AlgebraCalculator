// auto-generated
// Code generated by '$ fyne bundle'. DO NOT EDIT.

package main

import "fyne.io/fyne/v2"

var resourceSimpRulesExpandTxt = &fyne.StaticResource{
	StaticName: "simpRulesExpand.txt",
	StaticContent: []byte(
		"all_0 + var_1 = all_0 + 1 * var_1\r\n\r\n// Fix wrong Brace order\r\nall_0i + ( all_1i + all_2i ) = all_0i + all_1i + all_2i\r\nall_0i * ( all_1i * all_2i ) = all_0i * all_1i * all_2i\r\n\r\n// Remove Division\r\nall_0 / vec_1 = 1 / vec_1 * all_0\r\nall_0 / all_1 = all_0 * all_1 pow -1\r\n\r\n//Remove Subtraction\r\nall_0 - all_1 = all_0 + -1 * all_1\r\n\r\n// Brackte Rules:\r\n- All the same\r\n( all_0 + all_0 ) * ( all_0 + all_0 ) = 4 * all_0 pow 2\r\n\r\n// - Two\r\n( all_0 + all_1 ) pow 2 = ( all_0 + all_1 ) * ( all_0 + all_1 )\r\n( all_0 - all_1 ) pow 2 = ( all_0 - all_1 ) * ( all_0 - all_1 )\r\n\r\n( all_0i + all_1i ) * ( all_0i + all_1i ) = all_0 pow 2 + all_1 pow 2 + 2 * all_0 * all_1\r\n( all_0i + all_1i ) * ( all_1i + all_0i ) = all_0 pow 2 + all_1 pow 2 + 2 * all_0 * all_1\r\n\r\n( all_0i + all_1i ) * ( all_0i - all_1i ) = all_0 pow 2 - all_1 pow 2\r\n( all_0i + all_1i ) * ( all_1i - all_0i ) = all_1 pow 2 - all_0 pow 2\r\n\r\n( all_0i - all_1i ) * ( all_0i + all_1i ) = all_0 pow 2 - all_1 pow 2\r\n( all_0i - all_1i ) * ( all_1i + all_0i ) = all_1 pow 2 - all_0 pow 2\r\n\r\n( all_0i - all_1i ) * ( all_0i - all_1i ) = all_0 pow 2 + all_1 pow 2 - 2 * all_0 * all_1\r\n( all_0i - all_1i ) * ( all_1i - all_0i ) = all_0 pow 2 + all_1 pow 2 - 2 * all_0 * all_1\r\n\r\n// - Four\r\n( all_0i + all_1i ) * ( all_2i + all_3i ) = all_0 * all_2 + all_1 * all_2 + all_0 * all_3 + all_1 * all_3\r\n( all_0i + all_1i ) * ( all_2i - all_3i ) = all_0 * all_2 + all_1 * all_2 - all_0 * all_3 - all_1 * all_3\r\n( all_0i - all_1i ) * ( all_2i + all_3i ) = all_0 * all_2 - all_1 * all_2 + all_0 * all_3 - all_1 * all_3\r\n( all_0i - all_1i ) * ( all_2i - all_3i ) = all_0 * all_2 - all_1 * all_2 - all_0 * all_3 - all_1 * all_3\r\n\r\n// - Three\r\nall_0i * ( all_1i + all_2i ) = all_0 * all_1 + all_0 * all_2\r\nall_0i * ( all_1i - all_2i ) = all_0 * all_1 - all_0 * all_2\r\n( all_0i + all_1i ) * all_2i = all_2 * all_0 + all_2 * all_1\r\n( all_0i - all_1i ) * all_2i = all_2 * all_0 - all_2 * all_1\r\n\r\n// Delete Rules:\r\n0 + all_0 = all_0\r\n0 * all_0 = 0\r\nall_0 - all_0 = 0\r\nall_0 / 1 = all_0\r\n"),
}
var resourceSimpRulesSumUpTxt = &fyne.StaticResource{
	StaticName: "simpRulesSumUp.txt",
	StaticContent: []byte(
		"// Fix wrong Brace order\r\nall_0i + ( all_1i + all_2i ) = all_0i + all_1i + all_2i\r\nall_0i * ( all_1i * all_2i ) = all_0i * all_1i * all_2i\r\n\r\n// Remove Division\r\nall_0 / vec_1 = 1 / vec_1 * all_0\r\nall_0 / all_1 = all_0 * all_1 pow -1\r\n\r\n//Remove Subtraction\r\nall_0 - all_1 = all_0 + -1 * all_1\r\n\r\n// Merge Pow\r\nall_0 pow vec_1 * all_0 = all_0 pow ( vec_1 + 1 )\r\nall_0 pow vec_1 pow vec_1 = all_0 pow ( vec_1 * 2 )\r\nall_0 pow vec_1 pow vec_2 = all_0 pow ( vec_1 * vec_2 )\r\n\r\n// Merge Mul\r\nall_0 * all_0 = all_0 pow 2\r\nall_0 * all_1 * all_1 = all_0 * all_1 pow 2\r\n\r\n// Merge Addition\r\nvec_0 * all_1 + all_1 = ( vec_0 + 1 ) * all_1\r\nall_0 + vec_1 * all_2 + all_2 = all_0 + ( vec_1 + 1 ) * all_2\r\n\r\nall_0 + all_0 = 2 * all_0\r\nall_0 + all_1 + all_1 = all_0 + 2 * all_1\r\n\r\n( vec_0 * all_1 ) + ( vec_2 * all_1 ) = ( vec_0 + vec_2 ) * all_1\r\nall_0 + ( vec_1 * all_2 ) + ( vec_3 * all_2 ) = all_0 + ( vec_1 + vec_3 ) * all_2\r\n\r\n// Delete Rules:\r\n0 + all_0 = all_0\r\n0 * all_0 = 0\r\n1 * !var_0 = var_0\r\nall_0 - all_0 = 0\r\nall_0 / 1 = all_0\r\n"),
}
