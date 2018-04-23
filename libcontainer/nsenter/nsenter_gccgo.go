// +build linux,gccgo

package nsenter

/*
#cgo CFLAGS: -Wall
extern void nsexec();
extern void ns_init_mount_rootfs();
extern void ns_init_pivot_rootfs();
extern void ns_init_finalize_rootfs();
void __attribute__((constructor)) init(void) {
	nsexec();
	ns_init_mount_rootfs();
	ns_init_pivot_rootfs();
	ns_init_finalize_rootfs();
}
*/
import "C"

// AlwaysFalse is here to stay false
// (and be exported so the compiler doesn't optimize out its reference)
var AlwaysFalse bool

func init() {
	if AlwaysFalse {
		// by referencing this C init() in a noop test, it will ensure the compiler
		// links in the C function.
		// https://gcc.gnu.org/bugzilla/show_bug.cgi?id=65134
		C.init()
	}
}
