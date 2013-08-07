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

func TestAttachNonPrivate(t *testing.T) {
	shmid, err := Shmget(keyOfShm, sizeOfShm, IPC_CREAT | IPC_EXCL | 0777)

	if err != nil {
		t.Errorf("shmget error = " + err.Error())
	} 
	t.Logf("shmid = %d\n", shmid)

	var data []byte
	data, err = Shmat(shmid, 0)
	if err != nil {
		t.Errorf("Shmat error = " + err.Error())
	} 

	t.Logf("len(data) = %d\n", len(data))
	if len(data) != sizeOfShm {
		t.Errorf("len(data) = %d\n", len(data))
	}
	for i := 0; i < len(data); i++ {
		data[i] = byte(i & 0xff) 
	}

	err = Shmdt(data)
	if  err != nil {
		t.Errorf("Shmdt error = " + err.Error())
	} 

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
