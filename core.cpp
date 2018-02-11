#include "core.h"
#include <stdio.h>

int ncs_GetDeviceName(int idx, char* name) {
    mvncStatus r = mvncGetDeviceName(idx, name, NAME_SIZE);
    return int(r);
}

int ncs_OpenDevice(const char* name, void** deviceHandle) {
    mvncStatus r = mvncOpenDevice(name, deviceHandle);
    return int(r);
}

int ncs_CloseDevice(void* deviceHandle) {
    mvncStatus r = mvncCloseDevice(deviceHandle);
    return int(r);
}

int ncs_AllocateGraph(void* deviceHandle, void** graphHandle, void* graphData, unsigned int graphDataLen) {
    mvncStatus r = mvncAllocateGraph(deviceHandle, graphHandle, graphData, graphDataLen);
    return int(r);
}

int ncs_DeallocateGraph(void* graphHandle) {
    mvncStatus r = mvncDeallocateGraph(graphHandle);
    return int(r);
}

int ncs_LoadTensor(void* graphHandle, void* tensorData, unsigned int tensorDataLen) {
    mvncStatus r = mvncLoadTensor(graphHandle, tensorData, tensorDataLen, NULL);
    return int(r);
}
