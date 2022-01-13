/* Include files */

#include "blackbox_cgxe.h"
#include "m_vPQnrEvXbKz2rGpHnRdrkD.h"
#include "m_TTLCUDajK4f7O0CW8HSrh.h"

unsigned int cgxe_blackbox_method_dispatcher(SimStruct* S, int_T method, void
  * data)
{
  if (ssGetChecksum0(S) == 2816651782 &&
      ssGetChecksum1(S) == 19382575 &&
      ssGetChecksum2(S) == 4047763129 &&
      ssGetChecksum3(S) == 980478256) {
    method_dispatcher_vPQnrEvXbKz2rGpHnRdrkD(S, method, data);
    return 1;
  }

  if (ssGetChecksum0(S) == 3980587522 &&
      ssGetChecksum1(S) == 2885085520 &&
      ssGetChecksum2(S) == 1186614747 &&
      ssGetChecksum3(S) == 2353469679) {
    method_dispatcher_TTLCUDajK4f7O0CW8HSrh(S, method, data);
    return 1;
  }

  return 0;
}
