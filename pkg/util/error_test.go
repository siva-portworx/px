/*
Copyright © 2019 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package util

import (
	"testing"
	"strings"
	"fmt"
	"bytes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

var(
	// gRPC error Message
	grpcMsg = "Volume Not found"
	// gRPC error			 
	grpcErrStatus = status.Error(codes.NotFound, grpcMsg)
	// Volume name 
	volumeName = "volume1"
)

/* TestPxErrorMessage:
 * Testcase to validate the success case of PxErrorMessage().
 * The return value of the PxErrorMessage need to composed meessage of
 * both the gRPX error status message and the provided message.
 */
func TestPxErrorMessage(t * testing.T) {
    // Expected error
	expectedMsg := fmt.Sprintf("%s: %s" ,volumeName, grpcMsg)

	err := PxErrorMessage(grpcErrStatus, volumeName)
    if (strings.Compare(err.Error(), expectedMsg) != 0) {
		t.Errorf("TestPxErrorMessage Failed, Got: %s, Expected: %s",
				err.Error(), expectedMsg)
	}
}

/* TestPxErrorMessagef:
 * Testcase to validate the success case of PxErrorMessagef().
 * The return value of the PxErrorMessage need to composed meessage of
 * both the gRPX error status message and the provided message.
 */
func TestPxErrorMessagef(t * testing.T) {
    // Expected message
	expectedMsg := fmt.Sprintf("[%s]: %s", volumeName, grpcMsg)

	err := PxErrorMessagef(grpcErrStatus, "%s", volumeName)
    if (strings.Compare(err.Error(), expectedMsg) != 0) {
		t.Errorf("TestPxErrorMessagef Failed, Got: %s, Expected: %s",
				err.Error(), expectedMsg)
	}
}

/* TestPrintPxErrorMessagef:
 * Testcase to validate the success case of PrintPxErrorMessagef().
 * PrintPxErrorMessagef will print the Portworx error message to stderr.
 * Make sure that message passed to PrintPxErrorMessagef is appears on
 * stderr.
 */
func TestPrintPxErrorMessagef(t * testing.T) {
	//Excpected message
    expectedMsg := fmt.Sprintf("[%s]: %s\n", volumeName, grpcMsg)
	// Save
	oldStderr := Stderr
	// Create new buffer
	stderr := new(bytes.Buffer)
	// Set buffer
	Stderr = stderr
	PrintPxErrorMessagef(grpcErrStatus, "%s", volumeName)
	if (strings.Compare(stderr.String(), expectedMsg) != 0) {
		t.Errorf("TestPrintPxErrorMessagef Failed, Got: %s, Expected: %s",
				stderr.String(), expectedMsg)
	}
	// Restore
	Stderr = oldStderr
}

/* TestPxError:
 * Testcase to validate the success case of PxError().
 * PxError extract only the message found in the gRPC error status
 */
func TestPxError(t * testing.T) {
    // Expected message
	expectedMsg :=  fmt.Sprintf("%s\n", grpcMsg)

	err := PxError(grpcErrStatus)
    if (strings.Compare(err.Error(), expectedMsg) != 0) {
		t.Errorf("TestPxErrorMessagef Failed, Got: %s, Expected: %s",
				err.Error(), expectedMsg)
	}
}
