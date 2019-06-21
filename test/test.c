#include <windows.h>
#include <stdio.h>
#include <stdint.h>

typedef char* (__stdcall *AppInfo)();
typedef int32_t (__stdcall *Initialize)(int32_t token);
int main()
{
    printf("prog start\n");
    HMODULE hMod = LoadLibrary("app.dll");
    if (hMod == NULL)
    {
        printf("%d\n", GetLastError());
        return 1;
    }

    AppInfo fp = (AppInfo)GetProcAddress(hMod, "AppInfo");
    if (fp != NULL){
        char *c = fp();
        printf("%s\n", c);
    }else{
        printf("func not found\n");
    }

    Initialize init = (Initialize)GetProcAddress(hMod, "Initialize");
    if (init != NULL){
        int32_t ret = init(1234);
        printf("%d\n", ret);
    }else{
        printf("func not found\n");
    }

    FreeLibrary(hMod);
    printf("prog exit\n");
}