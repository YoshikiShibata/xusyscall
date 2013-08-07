// Copyright (c) 2013 Yoshiki Shibata. All rights reserved.

// This package is intended to provide some system calls which are not
// provided the original Go syscall package.

package xusyscall

// Note that some system call such as shmget is not a system call on some
// systems such as 32bit Linux. So we need to implement this package in C
// language

// #include <sys/ipc.h>
// #include <sys/shm.h>
// #include <errno.h>
import "C"

import "unsafe"

const (
	IPC_PRIVATE = 0		// Private key

	IPC_CREAT     =  01000 // Create key if key does not exist
	IPC_EXCL      =  02000 // Fail if key exists
	SHM_HUGETLB	  =  04000 // segment is mapped via hugetlb
	SHM_NORESERVE = 010000 // don't check for reservations

	ipc_RMID = 0
	ipc_SET  = 1
	ipc_STAT = 2
	ipc_INFO = 3
)

func Shmget(key int, size int, shmflg int) (shmid int, err error) {
	result, err := C.shmget(C.key_t(key), C.size_t(size), C.int(shmflg))

	if err != nil {
		shmid = int(result)
	}
	return
}

func Shmat(shmid int, shmflg int) (data []byte, err error) {
	addr, errno := shmat(shmid, 9, shmflg)
	if errno != nil {
		return nil, errno
	}

	length, errno2 := shmseqsz(shmid)
	if errno2 != nil {
		return nil, errno2
	}

	// Slice memory layout
	var sl = struct {
		addr uintptr
		len  int
		cap  int
	}{addr, length, length}

	data = *(*[]byte)(unsafe.Pointer(&sl))
	return data, nil
}

func shmat(shmid int, shmaddr uintptr, shmflg int) (addr uintptr, err error) {
	result, errno := C.shmat(C.int(shmid), unsafe.Pointer(shmaddr), C.int(shmflg))

	if int(uintptr(result)) == -1 {
		return 0, errno
	}

	return uintptr(result), nil
}

func shmseqsz(shmid int) (segsz int, err error) {
	var shmid_ds C.struct_shmid_ds

	result, errno := C.shmctl(C.int(shmid), ipc_STAT, 
			(*_Ctype_struct_shmid_ds)(unsafe.Pointer(&shmid_ds)))

	if result == -1 {
		return 0, errno
	}
	return int(shmid_ds.shm_segsz), nil
}
