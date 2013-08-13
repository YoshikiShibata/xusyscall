// Copyright (c) 2013 Yoshiki Shibata. All rights reserved.

package xusyscall

import "testing"

const (
	keyOfShm  = 1701
	sizeOfShm = 1024 * 1024
)

func TestIllegalKeyForGet(t *testing.T) {
	defer failIfRecoverIsNil(t)

	// negative key
	Shmget(-1, sizeOfShm, IPC_CREAT|IPC_EXCL)
}

func failIfRecoverIsNil(t *testing.T) {
	if x := recover(); x == nil {
		t.Fail()
	}
}

func TestIllegalSizeForGet(t *testing.T) {
	defer failIfRecoverIsNil(t)

	// negative size
	Shmget(IPC_PRIVATE, -1, IPC_CREAT|IPC_EXCL)
}

func TestIllegalFlagForGet(t *testing.T) {
	defer failIfRecoverIsNil(t)

	// negative flag
	Shmget(IPC_PRIVATE, sizeOfShm, -1)
}

func TestPrivateGet(t *testing.T) {
	testGet(IPC_PRIVATE, t)
}

func TestNonPrivateGet(t *testing.T) {
	testGet(keyOfShm, t)
}

func testGet(key int, t *testing.T) {
	shmid, err := Shmget(key, sizeOfShm, IPC_CREAT|IPC_EXCL|0777)

	if err != nil {
		t.Errorf("shmget error = " + err.Error())
		t.Fail()
		return
	}

	t.Logf("shmid = %d\n", shmid)
	removeSharedMemory(shmid, t)
}

func removeSharedMemory(shmid int, t *testing.T) {
	t.Logf("removeSharedMemory shmid = %d\n", shmid)
	err := Shmrm(shmid)

	if err != nil {
		t.Errorf("Shmrm error = " + err.Error())
		t.Fail()
	}
}

func TestReadOnlyAccessToPrivate(t *testing.T) {
	testReadOnlyAccess(IPC_PRIVATE, t)
}

func TestReadOnlyAccessToNonPrivate(t *testing.T) {
	testReadOnlyAccess(keyOfShm, t)
}

func testReadOnlyAccess(key int, t *testing.T) {
	shmid, err := Shmget(key, sizeOfShm, IPC_CREAT|IPC_EXCL|0777)

	if err != nil {
		t.Errorf("Shmget error = " + err.Error())
		return
	}

	defer removeSharedMemory(shmid, t)

	var data []byte
	data, err = Shmat(shmid, SHM_RDONLY)

	if err != nil {
		t.Error("Shmat error = " + err.Error())
		return
	}

	defer detachSharedMemory(data, t)

	if !isExpectedLength(data, t) {
		return
	}

	readData(data, t)

	// Following call causes fatal error which cannnot be recovered.
	// writeData(data)
}

func TestReadWriteAccessToPrivate(t *testing.T) {
	testReadWriteAccess(IPC_PRIVATE, t)
}

func TestReadWriteAccessToNonPrivate(t *testing.T) {
	testReadWriteAccess(keyOfShm, t)
}

func testReadWriteAccess(key int, t *testing.T) {
	shmid, err := Shmget(key, sizeOfShm, IPC_CREAT|IPC_EXCL|0777)

	if err != nil {
		t.Errorf("Shmget error = " + err.Error())
		return
	}

	defer removeSharedMemory(shmid, t)

	var data []byte
	data, err = Shmat(shmid, 0)

	if err != nil {
		t.Error("Shmat error = " + err.Error())
		return
	}

	defer detachSharedMemory(data, t)

	if !isExpectedLength(data, t) {
		return
	}

	writeData(data)
	verifyWrittenData(data, t)
}

func isExpectedLength(data []byte, t *testing.T) bool {
	t.Logf("length of data is %d bytes", len(data))

	if len(data) != sizeOfShm {
		t.Errorf("unexpected length of attached size = %d", len(data))
		return false
	}
	return true
}

func readData(data []byte, t *testing.T) {
	count := 0

	for d := range data {
		count = count + 1 + d
	}
	t.Logf("count = %d", count)
}

func writeData(data []byte) {
	for i := 0; i < len(data); i++ {
		data[i] = byte(i & 0xff)
	}
}

func verifyWrittenData(data []byte, t *testing.T) {
	for i := 0; i < len(data); i++ {
		if data[i] != byte(i&0xff) {
			t.Errorf("data[%d] is expected to be %d, but %d",
				i, byte(i&0xff), data[i])
		}
	}
}

func detachSharedMemory(data []byte, t *testing.T) {
	err := Shmdt(data)

	if err != nil {
		t.Error("Shmdt error = " + err.Error())
	}
}
