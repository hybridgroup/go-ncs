#include "core.h"

int ncs_GetDeviceName(int idx, char* name) {
    //char devName[NAME_SIZE];
    mvncStatus r = mvncGetDeviceName(idx, name, NAME_SIZE);
    return int(r);
}

int ncs_OpenDevice(const char* name, void* deviceHandle) {
    mvncStatus r = mvncOpenDevice(name, &deviceHandle);
    return int(r);
}

int ncs_CloseDevice(void* deviceHandle) {
    mvncStatus r = mvncCloseDevice(deviceHandle);
    return int(r);
}
