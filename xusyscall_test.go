// Copyright (c) 2013 Yoshiki Shibata. All rights reserved.

package xusyscall

import "testing"

const (
	keyOfShm = 1701
	sizeOfShm = 1024 * 1024
)

func TestPrivateGet(t *testing.T) {
	shmid, err := Shmget(IPC_PRIVATE, sizeOfShm, IPC_CREAT | IPC_EXCL)

	if err != nil {
		t.Errorf("shmget error = " + err.Error())
		t.Fail()
	} 

	t.Logf("shmid = %d\n", shmid)
}

func TestNonPrivateGet(t *testing.T) {
	shmid, err := Shmget(keyOfShm, sizeOfShm, IPC_CREAT | IPC_EXCL)

	if err != nil {
		t.Errorf("shmget error = " + err.Error())
	} 

	t.Logf("shmid = %d\n", shmid)
	cleanUpSharedMemory(shmid, t)
}

func cleanUpSharedMemory(shmid int, t *testing.T) {
	t.Logf("cleanUpSharedMemory shmid = %d\n", shmid)
	err := Shmrm(shmid)

	if err != nil {
		t.Errorf("Shmrm error = " + err.Error())
		t.Fail()
	}
}
