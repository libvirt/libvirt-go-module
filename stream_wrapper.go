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
#include <stdint.h>
#include <assert.h>
#include "stream_wrapper.h"

int streamSourceCallback(virStreamPtr st, char *cdata, size_t nbytes, int callbackID);
int streamSourceHoleCallback(virStreamPtr st, int *inData, long long *length, int callbackID);
int streamSourceSkipCallback(virStreamPtr st, long long length, int callbackID);

int streamSinkCallback(virStreamPtr st, const char *cdata, size_t nbytes, int callbackID);
int streamSinkHoleCallback(virStreamPtr st, long long length, int callbackID);

struct CallbackData {
    int callbackID;
    int holeCallbackID;
    int skipCallbackID;
};

static int streamSourceCallbackHelper(virStreamPtr st, char *data, size_t nbytes, void *opaque)
{
    struct CallbackData *cbdata = opaque;

    return streamSourceCallback(st, data, nbytes, cbdata->callbackID);
}

static int streamSourceHoleCallbackHelper(virStreamPtr st, int *inData, long long *length, void *opaque)
{
    struct CallbackData *cbdata = opaque;

    return streamSourceHoleCallback(st, inData, length, cbdata->holeCallbackID);
}

static int streamSourceSkipCallbackHelper(virStreamPtr st, long long length, void *opaque)
{
    struct CallbackData *cbdata = opaque;

    return streamSourceSkipCallback(st, length, cbdata->skipCallbackID);
}

static int streamSinkCallbackHelper(virStreamPtr st, const char *data, size_t nbytes, void *opaque)
{
    struct CallbackData *cbdata = opaque;

    return streamSinkCallback(st, data, nbytes, cbdata->callbackID);
}

static int streamSinkHoleCallbackHelper(virStreamPtr st, long long length, void *opaque)
{
    struct CallbackData *cbdata = opaque;

    return streamSinkHoleCallback(st, length, cbdata->holeCallbackID);
}

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


void
streamEventCallback(virStreamPtr st, int events, int callbackID);

static void
streamEventCallbackHelper(virStreamPtr st, int events, void *opaque)
{
    streamEventCallback(st, events, (int)(intptr_t)opaque);
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
virStreamEventAddCallbackHelper(virStreamPtr stream,
                                int events,
                                int callbackID,
                                virErrorPtr err)
{
    return virStreamEventAddCallbackWrapper(stream, events, streamEventCallbackHelper,
                                            (void *)(intptr_t)callbackID, NULL, err);
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
virStreamRecvAllHelper(virStreamPtr stream,
                       int callbackID,
                       virErrorPtr err)
{
    struct CallbackData cbdata = { .callbackID = callbackID };
    return virStreamRecvAllWrapper(stream, streamSinkCallbackHelper, &cbdata, err);
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
virStreamSendAllHelper(virStreamPtr stream,
                       int callbackID,
                       virErrorPtr err)
{
    struct CallbackData cbdata = { .callbackID = callbackID };
    return virStreamSendAllWrapper(stream, streamSourceCallbackHelper, &cbdata, err);
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
virStreamSparseRecvAllHelper(virStreamPtr stream,
                             int callbackID,
                             int holeCallbackID,
                             virErrorPtr err)
{
    struct CallbackData cbdata = { .callbackID = callbackID,
                                   .holeCallbackID = holeCallbackID };
    return virStreamSparseRecvAllWrapper(stream, streamSinkCallbackHelper,
                                         streamSinkHoleCallbackHelper,
                                         &cbdata, err);
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


int
virStreamSparseSendAllHelper(virStreamPtr stream,
                             int callbackID,
                             int holeCallbackID,
                             int skipCallbackID,
                             virErrorPtr err)
{
    struct CallbackData cbdata = { .callbackID = callbackID,
                                   .holeCallbackID = holeCallbackID,
                                   .skipCallbackID = skipCallbackID };
    return virStreamSparseSendAllWrapper(stream, streamSourceCallbackHelper,
                                         streamSourceHoleCallbackHelper,
                                         streamSourceSkipCallbackHelper,
                                         &cbdata, err);
}


*/
import "C"
