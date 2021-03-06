/*
 * This file is part of the libvirt-go-module project
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * Copyright (c) 2013 Alex Zorin
 * Copyright (C) 2016 Red Hat, Inc.
 *
 */

package libvirt

import (
	"fmt"
	"testing"
	"time"
)

func TestNodeDeviceEventRegister(t *testing.T) {

	if true {
		return
	}
	callbackId := -1

	conn := buildTestConnection()
	defer func() {
		if callbackId >= 0 {
			if err := conn.NodeDeviceEventDeregister(callbackId); err != nil {
				t.Errorf("got `%v` on NodeDeviceEventDeregister instead of nil", err)
			}
		}
		if res, _ := conn.Close(); res != 0 {
			t.Errorf("Close() == %d, expected 0", res)
		}
	}()

	defName := time.Now().String()

	nbEvents := 0

	callback := func(c *Connect, d *NodeDevice, event *NodeDeviceEventLifecycle) {
		if event.Event == NODE_DEVICE_EVENT_CREATED {
			domName, _ := d.GetName()
			if defName != domName {
				t.Fatalf("Name was not '%s': %s", defName, domName)
			}
		}
		eventString := fmt.Sprintf("%s", event)
		expected := "NodeDevice event=\"started\" detail=\"booted\""
		if eventString != expected {
			t.Errorf("event == %q, expected %q", eventString, expected)
		}
		nbEvents++
	}

	callbackId, err := conn.NodeDeviceEventLifecycleRegister(nil, callback)
	if err != nil {
		t.Error(err)
		return
	}

	// Test a minimally valid xml
	xml := `<device>
		<name>` + defName + `</name>
	</device>`
	dom, err := conn.DeviceCreateXML(xml, 0)
	if err != nil {
		t.Error(err)
		return
	}

	// This is blocking as long as there is no message
	EventRunDefaultImpl()
	if nbEvents == 0 {
		t.Fatal("At least one event was expected")
	}

	defer func() {
		dom.Destroy()
		dom.Free()
	}()

	// Check that the internal context entry was added, and that there only is
	// one.
	goCallbackLock.Lock()
	if len(goCallbacks) != 1 {
		t.Errorf("goCallbacks should hold one entry, got %+v", goCallbacks)
	}
	goCallbackLock.Unlock()

	// Deregister the event
	if err := conn.NodeDeviceEventDeregister(callbackId); err != nil {
		t.Fatalf("Event deregistration failed with: %v", err)
	}
	callbackId = -1 // Don't deregister twice

	// Check that the internal context entries was removed
	goCallbackLock.Lock()
	if len(goCallbacks) > 0 {
		t.Errorf("goCallbacks entry wasn't removed: %+v", goCallbacks)
	}
	goCallbackLock.Unlock()
}
