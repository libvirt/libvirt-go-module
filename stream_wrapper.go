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
#include <assert.h>
#include "stream_wrapper.h"


int
virStreamAbortWrapper(virStreamPtr stream,
                      virErrorPtr err)
{
    int ret = virStreamAbort(stream);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamEventAddCallbackWrapper(virStreamPtr stream,
                                 int events,
                                 virStreamEventCallback cb,
                                 void *opaque,
                                 virFreeCallback ff,
                                 virErrorPtr err)
{
    int ret = virStreamEventAddCallback(stream, events, cb, opaque, ff);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamEventRemoveCallbackWrapper(virStreamPtr stream,
                                    virErrorPtr err)
{
    int ret = virStreamEventRemoveCallback(stream);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamEventUpdateCallbackWrapper(virStreamPtr stream,
                                    int events,
                                    virErrorPtr err)
{
    int ret = virStreamEventUpdateCallback(stream, events);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamFinishWrapper(virStreamPtr stream,
                       virErrorPtr err)
{
    int ret = virStreamFinish(stream);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamFreeWrapper(virStreamPtr stream,
                     virErrorPtr err)
{
    int ret = virStreamFree(stream);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamRecvWrapper(virStreamPtr stream,
                     char *data,
                     size_t nbytes,
                     virErrorPtr err)
{
    int ret = virStreamRecv(stream, data, nbytes);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamRecvAllWrapper(virStreamPtr stream,
                        virStreamSinkFunc handler,
                        void *opaque,
                        virErrorPtr err)
{
    int ret = virStreamRecvAll(stream, handler, opaque);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamRecvFlagsWrapper(virStreamPtr stream,
                          char *data,
                          size_t nbytes,
                          unsigned int flags,
                          virErrorPtr err)
{
#if LIBVIR_VERSION_NUMBER < 3004000
    assert(0); // Caller should have checked version
#else
    int ret = virStreamRecvFlags(stream, data, nbytes, flags);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
#endif
}


int
virStreamRecvHoleWrapper(virStreamPtr stream,
                         long long *length,
                         unsigned int flags,
                         virErrorPtr err)
{
#if LIBVIR_VERSION_NUMBER < 3004000
    assert(0); // Caller should have checked version
#else
    int ret = virStreamRecvHole(stream, length, flags);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
#endif
}


int
virStreamRefWrapper(virStreamPtr stream,
                    virErrorPtr err)
{
    int ret = virStreamRef(stream);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamSendWrapper(virStreamPtr stream,
                     const char *data,
                     size_t nbytes,
                     virErrorPtr err)
{
    int ret = virStreamSend(stream, data, nbytes);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamSendAllWrapper(virStreamPtr stream,
                        virStreamSourceFunc handler,
                        void *opaque,
                        virErrorPtr err)
{
    int ret = virStreamSendAll(stream, handler, opaque);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
}


int
virStreamSendHoleWrapper(virStreamPtr stream,
                         long long length,
                         unsigned int flags,
                         virErrorPtr err)
{
#if LIBVIR_VERSION_NUMBER < 3004000
    assert(0); // Caller should have checked version
#else
    int ret = virStreamSendHole(stream, length, flags);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
#endif
}


int
virStreamSparseRecvAllWrapper(virStreamPtr stream,
                              virStreamSinkFunc handler,
                              virStreamSinkHoleFunc holeHandler,
                              void *opaque,
                              virErrorPtr err)
{
#if LIBVIR_VERSION_NUMBER < 3004000
    assert(0); // Caller should have checked version
#else
    int ret = virStreamSparseRecvAll(stream, handler, holeHandler, opaque);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
#endif
}


int
virStreamSparseSendAllWrapper(virStreamPtr stream,
                              virStreamSourceFunc handler,
                              virStreamSourceHoleFunc holeHandler,
                              virStreamSourceSkipFunc skipHandler,
                              void *opaque,
                              virErrorPtr err)
{
#if LIBVIR_VERSION_NUMBER < 3004000
    assert(0); // Caller should have checked version
#else
    int ret = virStreamSparseSendAll(stream, handler, holeHandler, skipHandler, opaque);
    if (ret < 0) {
        virCopyLastError(err);
    }
    return ret;
#endif
}


*/
import "C"
