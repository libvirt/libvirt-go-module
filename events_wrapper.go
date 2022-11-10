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
#cgo pkg-config: libvirt
#include "events_wrapper.h"


void
virEventRegisterImplWrapper(virEventAddHandleFunc addHandle,
                            virEventUpdateHandleFunc updateHandle,
                            virEventRemoveHandleFunc removeHandle,
                            virEventAddTimeoutFunc addTimeout,
                            virEventUpdateTimeoutFunc updateTimeout,
                            virEventRemoveTimeoutFunc removeTimeout)
{
    virEventRegisterImpl(addHandle,
                         updateHandle,
                         removeHandle,
                         addTimeout,
                         updateTimeout,
                         removeTimeout);
}


int
virEventAddHandleWrapper(int fd,
                         int events,
                         virEventHandleCallback cb,
                         void *opaque,
                         virFreeCallback ff,
                         virErrorPtr err)
{
    int ret = virEventAddHandle(fd, events, cb, opaque, ff);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virEventAddTimeoutWrapper(int timeout,
                          virEventTimeoutCallback cb,
                          void *opaque,
                          virFreeCallback ff,
                          virErrorPtr err)
{
    int ret =  virEventAddTimeout(timeout, cb, opaque, ff);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virEventRegisterDefaultImplWrapper(virErrorPtr err)
{
    int ret = virEventRegisterDefaultImpl();
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


void
virEventUpdateHandleWrapper(int watch,
                            int events)
{
    virEventUpdateHandle(watch, events);
}


void
virEventUpdateTimeoutWrapper(int timer,
                             int timeout)
{
    virEventUpdateTimeout(timer, timeout);
}


int
virEventRemoveHandleWrapper(int watch,
                            virErrorPtr err)
{
    int ret = virEventRemoveHandle(watch);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virEventRemoveTimeoutWrapper(int timer,
                             virErrorPtr err)
{
    int ret = virEventRemoveTimeout(timer);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virEventRunDefaultImplWrapper(virErrorPtr err)
{
    int ret = virEventRunDefaultImpl();
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


*/
import "C"
