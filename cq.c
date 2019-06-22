#include <windows.h>
#include "app.h"
#include "cq.h"

int32_t ac;
CQEVENT Initialize(int32_t access_code)
{
    ac = access_code;
    HMODULE hmod = LoadLibrary("CQP.dll");
    CQ_addLog_Ptr = GetProcAddress(hmod, "CQ_addLog");
    return 0;
}

CQEVENT on_Enable() { return Enable(); }

//apis

int32_t CQ_addLog(GoInt32 priority, const char *type, const char *reason)
{
    int32_t ret = CQ_addLog_Ptr(ac, priority, type, reason);
    free(type); //释放字符串内存
    free(reason);
    return ret;
}
