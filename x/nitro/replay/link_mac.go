//go:build darwin && !sys_nitro_replayer
// +build darwin,!sys_nitro_replayer

package replay

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lnitro_replayer
import "C"
