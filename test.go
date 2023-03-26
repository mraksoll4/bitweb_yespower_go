package main

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
    "encoding/hex"
    "fmt"
    "unsafe"
)

func main() {
    inputHex := "00000020adb7da25618975d45e2c47746afffd09b52b748b262dec914efc3b2d95000000175bfb78b38be48d47ce7f357288bc7a38b02828d606594f99792afb5f24a9f2e7c62064b6a7001e20002dfe01010000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff4e0351200f04e7c6206408fabe6d6d00000000000000000000000000000000000000000000000000000000000000000100000000000000f7fa34c7a7c0ae140f706f6f6c2e72706c616e742e78797a00000000020000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf900f902950000000016001494344516ae82e165e54ca8b4e7279ffe2d3a0ed70120000000000000000000000000000000000000000000000000000000000000000000000000"
    inputBytes, _ := hex.DecodeString(inputHex)
    input := string(inputBytes)

    expectedOutputHex := "6123f3f1c4e95d052dbe2f0238426f192a9592c55e55a3394054be399a000000"
    expectedOutput, _ := hex.DecodeString(expectedOutputHex)

    output := make([]byte, 32)
    cInput := C.CString(input)
    defer C.free(unsafe.Pointer(cInput))
    cOutput := (*C.char)(unsafe.Pointer(&output[0]))
    C.yespower_hash(cInput, cOutput)
    fmt.Printf("Input: %s\n", inputHex)
    fmt.Printf("Expected Output: %x\n", expectedOutput)
    fmt.Printf("Actual Output: %x\n", output)
    if string(output) == string(expectedOutput) {
        fmt.Println("Test Passed!")
    } else {
        fmt.Println("Outputs do not match.\nTest Failed!")
    }
}