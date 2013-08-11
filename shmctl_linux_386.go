// Copyright (c) 2013 Yoshiki Shibata. All rights reserved.

package xusyscall

// #include <sys/shm.h>
// #include <errno.h>
import "C"

import "unsafe"

func shmctl(shmid int, cmd int, shmid_ds *C.struct_shmid_ds) (error) {
	result, errno := C.shmctl(C.int(shmid), C.int(cmd), 
            (*_Ctype_struct___shmid_ds_new)(unsafe.Pointer(shmid_ds)))

	if result == -1 {
		return errno
    }
    return nil
}
