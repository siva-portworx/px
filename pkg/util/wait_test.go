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
	"time"
	"strings"
	"fmt"
)

const (
    oneSecond = 1 * time.Second
	twoSecond = 2 * time.Second
	fiveSecond = 5 * time.Second
	//Timeout error message
	timeoutErrMsg = "Timed out"
	// Dummy error message
	dummyErrMsg = "Error Message"
)

/* TestWaitForTimeout:
 * Test the timeout case of WaitFor() function.
 * Set the duration parameter value of WaitFor() to 2 sec.
 * Set the period parameter value of WaitFor() to 1 sec.
 * In the input function wait for 5 sec which greater than duration value.
 * Expected behaviour: WaitFor should return timeout as the input function
 *                     takes more than 2 sec.
 */
func TestWaitForTimeoutCase(t * testing.T) {
	if err := WaitFor(twoSecond, oneSecond, func() (bool, error) {
		time.Sleep(fiveSecond)
		return true, nil
	}); err == nil {
        // Got err as nil but expected is "Timed Out" message
		t.Errorf("TestWaitForTimeoutCase Failed, got: nil, expected: %s ",
				timeoutErrMsg)
	} else if err != nil && (strings.Compare(err.Error(), timeoutErrMsg) != 0) {
		// Got err as non nil but the message is not "Timed Out" message
		t.Errorf("TestWaitForTimeoutCase Failed, got: %s, expected: %s ",
				err.Error(), timeoutErrMsg)
	}
}

/* TestWaitForInputFunctionReturnFalseNil:
 * Test the case, when the input function return false and nil
 * Expected behaviour: WaitFor should return nil.
 */
func TestWaitForInputFunctionReturnFalseNil(t * testing.T) {
	if err := WaitFor(fiveSecond, oneSecond, func() (bool, error) {
		time.Sleep(twoSecond)
		return false, nil
	}); err != nil {
		// Got unexpected err but expected err was nil
		t.Errorf("TestWaitForInputFunctionReturnFalseNil Failed, got: %s, expected: nil", err.Error())
    }
}

/* TestWaitForInputFunctionReturnFalseNonNil:
 * Test the case, when the input function return false and nil
 * Expected behaviour: WaitFor should return nil.
 */
func TestWaitForInputFunctionReturnFalseNonNil(t * testing.T) {
	if err := WaitFor(fiveSecond, oneSecond, func() (bool, error) {
		time.Sleep(twoSecond)
		return false, fmt.Errorf(dummyErrMsg)
	}); err == nil {
		// Got nil but expected err as dummyErrMsg
		t.Errorf("TestWaitForInputFunctionReturnFalseNonNil Failed, got: nil , expected: \"%s\"", dummyErrMsg)
    } else if err != nil && (strings.Compare(err.Error(), dummyErrMsg) != 0) {
		// Got unexpected error but expected err as dummyErrMsg
		t.Errorf("TestWaitForInputFunctionReturnFalseNonNil Failed, got: %s, expected: \"%s\"", err.Error(), dummyErrMsg)
    }
}

/* TestWaitForInputFunctionReturnTrueNil:
 * Test the case, when the input function return true and nil value
 * Expected behaviour: WaitFor should return timeoutErrMsg.
 */
func TestWaitForInputFunctionReturnTrueNil(t * testing.T) {
	if err := WaitFor(fiveSecond, oneSecond, func() (bool, error) {
		time.Sleep(twoSecond)
		return true, nil
	}); err == nil {
		// Got err as nil but expected err as timeoutErrMsg
		t.Errorf("TestWaitForInputFunctionReturnTrueNil Failed, got: nil, expected: %s", timeoutErrMsg)
	} else if err != nil && (strings.Compare(err.Error(), timeoutErrMsg) != 0) {
		//Got unexpected err but expected err as timeoutErrMsg
		t.Errorf("TestWaitForInputFunctionReturnTrueNil Failed, got: %s, expected: %s ", err.Error(), timeoutErrMsg)
    }
}

/* TestWaitForInputFunctionReturnTrueNonNil:
 * Test the case, when the input function return true and non nil value
 * Expected behaviour: WaitFor should return nil.
 */
func TestWaitForInputFunctionReturnTrueNonNil(t * testing.T) {
	if err := WaitFor(fiveSecond, oneSecond, func() (bool, error) {
		time.Sleep(twoSecond)
		return true, fmt.Errorf(dummyErrMsg)
	}); err == nil {
		//Got err as nil but expected err as dummyErrMsg
		t.Errorf("TestWaitForInputFunctionReturnTrueNonNil Failed, got: nil, expected: %s", dummyErrMsg)
	} else if err != nil && (strings.Compare(err.Error(), dummyErrMsg) != 0) {
		//Gor unexpected err but expected err as nil
		t.Errorf("TestWaitForInputFunctionReturnTrueNonNil Failed, got: %s, expected: %s", err.Error(), dummyErrMsg)
    }
}
