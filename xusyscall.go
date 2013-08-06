// Copyright (c) 2013 Yoshiki Shibata. All rights reserved.

// This package is intended to provide some system calls which are not
// provided the original Go syscall package.

package uxsyscall

// Note that some system call such as shmget is not a system call on some
// systems such as 32bit Linux. So we need to implement this package in C
// language
