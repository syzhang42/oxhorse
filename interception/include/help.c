#include <windows.h>

UINT MapVirtualKey(UINT code,UINT mode){
    return MapVirtualKey(code,MAPVK_VK_TO_VSC);
};