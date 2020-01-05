package mockit

import (
	"testing"
)

// Mock contains methods to mock a call with specified arguments, and verify it
type Mock interface {

	// Mock mocks the call for the specified arguments
	Mock(t *testing.T, in []interface{}, out []interface{})

	// Restore restore the mock configuration, unmocked by UnMock
	Restore()

	// UnMock disable the mock configuration, allowing to call the real method
	UnMock()

	// Verify fails the test if a call with the specified arguments wasn't made
	Verify(t *testing.T, in []interface{})
}
