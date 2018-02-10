#ifndef _GONCS_H_
#define _GONCS_H_

#include <stdlib.h>
#include <mvnc.h>

#define NAME_SIZE 100

#ifdef __cplusplus
extern "C" {
#endif

int ncs_GetDeviceName(int idx, char* name);

#ifdef __cplusplus
}
#endif

#endif //_GONCS_H_
