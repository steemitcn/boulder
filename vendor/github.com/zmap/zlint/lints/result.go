/*
 * ZLint Copyright 2017 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package lints

import "encoding/json"

// LintStatus is an enum returned by lints inside of a LintResult.
type LintStatus int

// Known LintStatus values
const (
	// Unused / unset LintStatus
	Reserved LintStatus = 0

	// Not Applicable
	NA LintStatus = 1

	// Not Effective
	NE LintStatus = 2

	Pass   LintStatus = 3
	Notice LintStatus = 4
	Warn   LintStatus = 5
	Error  LintStatus = 6
	Fatal  LintStatus = 7
)

// LintResult contains a LintStatus, and an optional human-readable description.
// The output of a lint is a LintResult.
type LintResult struct {
	Status  LintStatus `json:"result"`
	Details string     `json:"details,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface.
func (e LintStatus) MarshalJSON() ([]byte, error) {
	s := e.String()
	return json.Marshal(s)
}

// String returns the canonical representation of a LintStatus as a string.
func (e LintStatus) String() string {
	switch e {
	case NA:
		return "NA"
	case NE:
		return "NE"
	case Pass:
		return "pass"
	case Notice:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	default:
		return ""
	}
}