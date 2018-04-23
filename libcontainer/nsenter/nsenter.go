// +build linux,!gccgo

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
