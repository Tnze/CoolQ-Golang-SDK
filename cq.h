#include <stdint.h>

//events
extern int32_t __stdcall Initialize(int32_t p0);
extern int32_t __stdcall onEnable();

//apis
extern int32_t __stdcall CQ_addLog(int32_t ac, int32_t priority, const char* type, const char* reason);
extern int32_t __stdcall CQ_sendPrivateMsg(int32_t ac, int64_t qq, const char* msg);