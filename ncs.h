#ifndef _GONCS_H_
#define _GONCS_H_

#ifdef __APPLE__
#include <sys/types.h>
#endif // __APPLE__
#include <stdlib.h>
#include <mvnc.h>

#define NAME_SIZE 100

typedef unsigned int uint;

typedef struct ResultData{
  void* data;
  uint length;
} ResultData;

#ifdef __cplusplus
extern "C" {
#endif

int ncs_GetDeviceName(int idx, char* name);
int ncs_OpenDevice(const char* name, void** deviceHandle);
int ncs_CloseDevice(void* deviceHandle);
int ncs_AllocateGraph(void* deviceHandle, void** graphHandle, void* graphData, unsigned int graphDataLen);
int ncs_DeallocateGraph(void* graphHandle);
int ncs_LoadTensor(void* graphHandle, void* tensorData, unsigned int tensorDataLen);
int ncs_GetResult(void* graphHandle, ResultData* resultData);

#ifdef __cplusplus
}
#endif

#endif //_GONCS_H_
