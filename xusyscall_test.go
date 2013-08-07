// Copyright (c) 2013 Yoshiki Shibata. All rights reserved.

package xusyscall

import "testing"

func TestPrivateGet(t *testing.T) {
	shmid, err := Shmget(IPC_PRIVATE, 1024, IPC_CREAT | IPC_EXCL)

	if err != nil {
		t.Errorf("shmget error = " + err.Error())
		t.Fail()
	} 

	t.Logf("shmid = %d\n", shmid)
}

func TestNonPrivateGet(t *testing.T) {
	shmid, err := Shmget(1701, 1024 * 4, IPC_CREAT | IPC_EXCL)

	if err != nil {
		t.Errorf("shmget error = " + err.Error())
		t.Fail()
	} 

	t.Logf("shmid = %d\n", shmid)
}
