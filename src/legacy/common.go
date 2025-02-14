package legacy

/*
#include "defs.h"
*/
import "C"
import noxflags "github.com/opennox/opennox/v1/common/flags"

//export nox_common_setEngineFlag
func nox_common_setEngineFlag(flags C.nox_engine_flag) {
	noxflags.SetEngine(noxflags.EngineFlag(flags))
}

//export nox_common_resetEngineFlag
func nox_common_resetEngineFlag(flags C.nox_engine_flag) {
	noxflags.UnsetEngine(noxflags.EngineFlag(flags))
}

//export nox_common_getEngineFlag
func nox_common_getEngineFlag(flags C.nox_engine_flag) C.bool {
	return C.bool(noxflags.HasEngine(noxflags.EngineFlag(flags)))
}

//export nox_common_randomInt_415FA0
func nox_common_randomInt_415FA0(min, max int) int {
	return GetServer().S().Rand.Logic.IntClamp(min, max)
}

//export nox_common_randomIntMinMax_415FF0
func nox_common_randomIntMinMax_415FF0(min, max int, file *C.char, line int) int {
	return GetServer().S().Rand.Other.Int(min, max)
}

//export nox_common_randomFloat_416030
func nox_common_randomFloat_416030(min, max C.float) C.double {
	return C.double(GetServer().S().Rand.Logic.FloatClamp(float64(min), float64(max)))
}

//export nox_common_randomFloatXxx_416090
func nox_common_randomFloatXxx_416090(min, max C.float) C.double {
	return C.double(GetServer().S().Rand.Other.Float(float64(min), float64(max)))
}
