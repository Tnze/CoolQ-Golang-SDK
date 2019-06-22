#include <stdint.h>
#include <stdlib.h>

#define CQEVENT extern int32_t __stdcall

#define CQAPI(RetType, Name, ...)                            \
    RetType(__stdcall *Name##_Ptr)(int32_t ac, __VA_ARGS__); \
    RetType Name(__VA_ARGS__)

extern char* __stdcall AppInfo();
//events
CQEVENT Initialize(int32_t p0);
CQEVENT EVENT_ON_ENABLE();

//apis
CQAPI(int32_t, CQ_addLog, int32_t priority, char *type, char *reason);
CQAPI(int32_t, CQ_sendPrivateMsg, int64_t qq, char *msg);