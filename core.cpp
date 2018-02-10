#include "core.h"

int ncs_GetDeviceName(int idx, char* name) {
    //char devName[NAME_SIZE];
    mvncStatus r = mvncGetDeviceName(idx, name, NAME_SIZE);
    return int(r);
}
