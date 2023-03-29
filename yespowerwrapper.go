package bitweb_yespower_go

/*
#cgo CFLAGS: -I${SRCDIR}/yespower/libs/include
#cgo LDFLAGS: -L${SRCDIR}/yespower/libs/lib -lyespower -Wl,-rpath=${SRCDIR}/yespower/libs/lib

#include <stdlib.h>
#include "yespower.h"
#include "sysendian.h"

static const yespower_params_t v1 = {YESPOWER_0_5, 4096, 16, "Client Key", 10};

static const yespower_params_t v2 = {YESPOWER_1_0, 2048, 32, NULL, 0};

int yespower_hash(const char *input, char *output)
{
    uint32_t time = le32dec(&input[68]);
    if (time > 1676761800) {
        return yespower_tls(input, 80, &v1, (yespower_binary_t *) output);
    } else {
        return yespower_tls(input, 80, &v2, (yespower_binary_t *) output);
    }
}
*/
import "C"
import (
    "unsafe"
)

func YespowerHash(input []byte) []byte {
	output := make([]byte, 32)
	cInput := C.CString(string(input))
	defer C.free(unsafe.Pointer(cInput))
	cOutput := (*C.char)(unsafe.Pointer(&output[0]))
	C.yespower_hash(cInput, cOutput)
	return output
}