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

/*
#cgo !libvirt_dlopen pkg-config: libvirt
#cgo libvirt_dlopen LDFLAGS: -ldl
#cgo libvirt_dlopen CFLAGS: -DLIBVIRT_DLOPEN
#include <stdlib.h>
#include "libvirt_generated.h"
*/
import "C"

import (
	"unsafe"
)

type NodeDeviceEventID int

const (
	NODE_DEVICE_EVENT_ID_LIFECYCLE = NodeDeviceEventID(C.VIR_NODE_DEVICE_EVENT_ID_LIFECYCLE)
	NODE_DEVICE_EVENT_ID_UPDATE    = NodeDeviceEventID(C.VIR_NODE_DEVICE_EVENT_ID_UPDATE)
)

type NodeDeviceEventLifecycleType int

const (
	NODE_DEVICE_EVENT_CREATED   = NodeDeviceEventLifecycleType(C.VIR_NODE_DEVICE_EVENT_CREATED)
	NODE_DEVICE_EVENT_DELETED   = NodeDeviceEventLifecycleType(C.VIR_NODE_DEVICE_EVENT_DELETED)
	NODE_DEVICE_EVENT_DEFINED   = NodeDeviceEventLifecycleType(C.VIR_NODE_DEVICE_EVENT_DEFINED)
	NODE_DEVICE_EVENT_UNDEFINED = NodeDeviceEventLifecycleType(C.VIR_NODE_DEVICE_EVENT_UNDEFINED)
)

type NodeDeviceCreateXMLFlags int

const (
	NODE_DEVICE_CREATE_XML_VALIDATE = NodeDeviceCreateXMLFlags(C.VIR_NODE_DEVICE_CREATE_XML_VALIDATE)
)

type NodeDeviceDefineXMLFlags int

const (
	NODE_DEVICE_DEFINE_XML_VALIDATE = NodeDeviceDefineXMLFlags(C.VIR_NODE_DEVICE_DEFINE_XML_VALIDATE)
)

type NodeDeviceXMLFlags int

const (
	NODE_DEVICE_XML_INACTIVE = NodeDeviceXMLFlags(C.VIR_NODE_DEVICE_XML_INACTIVE)
)

type NodeDeviceUpdateFlags int

const (
	NODE_DEVICE_UPDATE_AFFECT_CURRENT = NodeDeviceUpdateFlags(C.VIR_NODE_DEVICE_UPDATE_AFFECT_CURRENT)
	NODE_DEVICE_UPDATE_AFFECT_CONFIG  = NodeDeviceUpdateFlags(C.VIR_NODE_DEVICE_UPDATE_AFFECT_CONFIG)
	NODE_DEVICE_UPDATE_AFFECT_LIVE    = NodeDeviceUpdateFlags(C.VIR_NODE_DEVICE_UPDATE_AFFECT_LIVE)
)

type NodeDevice struct {
	ptr C.virNodeDevicePtr
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceFree
func (n *NodeDevice) Free() error {
	var err C.virError
	ret := C.virNodeDeviceFreeWrapper(n.ptr, &err)
	if ret == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceRef
func (c *NodeDevice) Ref() error {
	var err C.virError
	ret := C.virNodeDeviceRefWrapper(c.ptr, &err)
	if ret == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceDestroy
func (n *NodeDevice) Destroy() error {
	var err C.virError
	result := C.virNodeDeviceDestroyWrapper(n.ptr, &err)
	if result == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceReset
func (n *NodeDevice) Reset() error {
	var err C.virError
	result := C.virNodeDeviceResetWrapper(n.ptr, &err)
	if result == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceDettach
func (n *NodeDevice) Detach() error {
	var err C.virError
	result := C.virNodeDeviceDettachWrapper(n.ptr, &err)
	if result == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceDetachFlags
func (n *NodeDevice) DetachFlags(driverName string, flags uint32) error {
	cDriverName := C.CString(driverName)
	defer C.free(unsafe.Pointer(cDriverName))
	var err C.virError
	result := C.virNodeDeviceDetachFlagsWrapper(n.ptr, cDriverName, C.uint(flags), &err)
	if result == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceReAttach
func (n *NodeDevice) ReAttach() error {
	var err C.virError
	result := C.virNodeDeviceReAttachWrapper(n.ptr, &err)
	if result == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceGetName
func (n *NodeDevice) GetName() (string, error) {
	var err C.virError
	name := C.virNodeDeviceGetNameWrapper(n.ptr, &err)
	if name == nil {
		return "", makeError(&err)
	}
	return C.GoString(name), nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceGetXMLDesc
func (n *NodeDevice) GetXMLDesc(flags NodeDeviceXMLFlags) (string, error) {
	var err C.virError
	result := C.virNodeDeviceGetXMLDescWrapper(n.ptr, C.uint(flags), &err)
	if result == nil {
		return "", makeError(&err)
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceGetParent
func (n *NodeDevice) GetParent() (string, error) {
	var err C.virError
	result := C.virNodeDeviceGetParentWrapper(n.ptr, &err)
	if result == nil {
		return "", makeError(&err)
	}
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result), nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceNumOfCaps
func (p *NodeDevice) NumOfCaps() (int, error) {
	var err C.virError
	result := int(C.virNodeDeviceNumOfCapsWrapper(p.ptr, &err))
	if result == -1 {
		return 0, makeError(&err)
	}
	return result, nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceListCaps
func (p *NodeDevice) ListCaps() ([]string, error) {
	const maxCaps = 1024
	var names [maxCaps](*C.char)
	namesPtr := unsafe.Pointer(&names)
	var err C.virError
	numCaps := C.virNodeDeviceListCapsWrapper(
		p.ptr,
		(**C.char)(namesPtr),
		maxCaps, &err)
	if numCaps == -1 {
		return nil, makeError(&err)
	}
	goNames := make([]string, numCaps)
	for k := 0; k < int(numCaps); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceCreate
func (p *NodeDevice) Create(flags uint32) error {
	var err C.virError
	result := C.virNodeDeviceCreateWrapper(p.ptr, C.uint(flags), &err)
	if result == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-nodedev.html#virNodeDeviceUndefine
func (p *NodeDevice) Undefine(flags uint32) error {
	var err C.virError
	result := C.virNodeDeviceUndefineWrapper(p.ptr, C.uint(flags), &err)
	if result == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-network.html#virNodeDeviceGetAutostart
func (n *NodeDevice) GetAutostart() (bool, error) {
	var out C.int
	var err C.virError
	result := C.virNodeDeviceGetAutostartWrapper(n.ptr, (*C.int)(unsafe.Pointer(&out)), &err)
	if result == -1 {
		return false, makeError(&err)
	}
	switch out {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

// See also https://libvirt.org/html/libvirt-libvirt-network.html#virNodeDeviceSetAutostart
func (n *NodeDevice) SetAutostart(autostart bool) error {
	var cAutostart C.int
	switch autostart {
	case true:
		cAutostart = 1
	default:
		cAutostart = 0
	}
	var err C.virError
	result := C.virNodeDeviceSetAutostartWrapper(n.ptr, cAutostart, &err)
	if result == -1 {
		return makeError(&err)
	}
	return nil
}

// See also https://libvirt.org/html/libvirt-libvirt-network.html#virNodeDeviceIsActive
func (n *NodeDevice) IsActive() (bool, error) {
	var err C.virError
	result := C.virNodeDeviceIsActiveWrapper(n.ptr, &err)
	if result == -1 {
		return false, makeError(&err)
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

// See also https://libvirt.org/html/libvirt-libvirt-network.html#virNodeDeviceIsPersistent
func (n *NodeDevice) IsPersistent() (bool, error) {
	var err C.virError
	result := C.virNodeDeviceIsPersistentWrapper(n.ptr, &err)
	if result == -1 {
		return false, makeError(&err)
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

// See also https://libvirt.org/html/libvirt-libvirt-network.html#virNodeDeviceUpdate
func (n *NodeDevice) Update(xml string, flags NodeDeviceUpdateFlags) error {
	cXml := C.CString(xml)
	defer C.free(unsafe.Pointer(cXml))
	var err C.virError
	result := C.virNodeDeviceUpdateWrapper(n.ptr, cXml, C.uint(flags), &err)
	if result == -1 {
		return makeError(&err)
	}
	return nil
}
