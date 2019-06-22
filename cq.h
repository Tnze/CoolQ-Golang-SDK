#include <stdint.h>
#include <stdlib.h>

#define CQEVENT extern int32_t __stdcall

#define CQAPI(RetType, Name, ...)                \
    RetType(__stdcall *Name##_Ptr)(__VA_ARGS__); \
    RetType Name(__VA_ARGS__)

//events
CQEVENT Initialize(int32_t p0);
CQEVENT EVENT_ON_ENABLE();

//apis
CQAPI(int32_t, CQ_addLog, int32_t ac, int32_t priority, const char *type, const char *reason);
CQAPI(int32_t, CQ_sendPrivateMsg, int32_t ac, int64_t qq, const char *msg);