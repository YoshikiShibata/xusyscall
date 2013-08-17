// Copyright (c) 2013 Yoshiki Shibata. All rights reserved.

package xusyscall_test

import "fmt"
import . "."

func Example() {
	shmid, err := Shmget(keyOfShm, sizeOfShm, IPC_CREAT|IPC_EXCL|0777)
	if err != nil {
		fmt.Printf("shmget error = " + err.Error())
		return
	}

	defer func() {
		err := Shmrm(shmid)
		if err != nil {
			fmt.Printf("Shmrd error = " + err.Error())
			return
		}
	}()

	var data []byte
	data, err = Shmat(shmid, false)

	if err != nil {
		fmt.Printf("Shmat error = " + err.Error())
		return
	}

	defer func() {
		err := Shmdt(data)
		if err != nil {
			fmt.Printf("Shmdt error = " + err.Error())
			return
		}
	}()

	fmt.Printf("len(data) = %d\n", len(data))

	// Write and Read
	for i := 0; i < len(data); i++ {
		b := byte(i & 0xff)
		data[i] = b
		if data[i] != b {
			fmt.Printf("incorrect read\n")
		}
	}
	// Output:
	// len(data) = 1048576
}
