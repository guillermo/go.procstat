/*
Package procstat provides an interface for /proc/:pid/stat

See Stats structure for knowing the data that you can get, but consumend cpu, io and mem is there.

    stats := stats.Stat{Pid: os.Getpid()}
    err := stats.Update()
    if err != nil {
    	panic(err)
    }

    stats.Utime // User time consumed by pid
    stats.Vsize // Virtual memory
    stats.Rss // Memory allocated right now in ram.
    stats.Rsslim // Maximum allowd memory
    stats.DelayacctBlkioTicks // I think is delays because of IO in centiseconds

NOTES: If the comm have a space in the middle, this program will fail to read all the arguments.
Look in man proc for more info.
*/
package procstat

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Update s with current values, usign the pid stored in the Stat
func (s *Stat) Update() error {
	if s.Pid == 0 {
		return errors.New("can't check for pid 0")
	}

	path := filepath.Join("/proc", strconv.FormatInt(int64(s.Pid), 10), "stat")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fscanf(file, "%d %s %c %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
		&s.Pid, &s.Comm, &s.State,
		&s.PPid,
		&s.PGrp,
		&s.Session,
		&s.TtyNr,
		&s.Tpgid,
		&s.Flags,
		&s.Minflt,
		&s.Cminflt,
		&s.Majflt,
		&s.Cmajflt,
		&s.Utime,
		&s.Stime,
		&s.Cutime,
		&s.Cstime,
		&s.Priority,
		&s.Nice,
		&s.NumThreads,
		&s.Itrealvalue,
		&s.Starttime,
		&s.Vsize,
		&s.Rss,
		&s.Rsslim,
		&s.Startcode,
		&s.Endcode,
		&s.Startstack,
		&s.Kstkesp,
		&s.Kstkeip,
		&s.Signal,
		&s.Blocked,
		&s.Sigignore,
		&s.Sigcatch,
		&s.Wchan,
		&s.Nswap,
		&s.Cnswap,
		&s.ExitSignal,
		&s.Processor,
		&s.RtPriority,
		&s.Policy,
		&s.DelayacctBlkioTicks,
		&s.GuestTime,
		&s.CguestTime)

	return err
}

/*
Stat holds all the information available in /proc/:pid/stat with information about the process.  This is used by ps(1).
It is defined in /usr/src/linux/fs/proc/array.c.

The fields, in order, with their proper scanf(3) format
specifiers, are:
*/
type Stat struct {
	Pid int // (1) The process ID.

	Comm string // (2) The filename of the executable, in parentheses.  This is visible whether or not the executable is swapped out.

	State byte // (3) One character from the string "RSDZTW" where R is running, S is sleeping in an interruptible wait, D is waiting in uninterruptible disk sleep, Z is zombie, T is traced or stopped (on a signal), and W is paging.

	PPid int // (4) The PID of the parent.

	PGrp int // (5) The process group ID of the process.

	Session int // (6) The session ID of the process.

	TtyNr int // (7) The controlling terminal of the process.  (The minor device number is contained in the combination of bits 31 to 20 and 7 to 0; the major device number is in bits 15 to 8.)

	Tpgid int // (8) The ID of the foreground process group of the controlling terminal of the process.

	Flags int64 // (9) The kernel flags word of the process.  For bit meanings, see the PF_* defines in the Linux kernel source file include/linux/sched.h.  Details depend on the kernel version.

	Minflt int64 // (10) The number of minor faults the process has made which have not required loading a memory page from disk.

	Cminflt uint64 // (11) The number of minor faults that the process's waited-for children have made.

	Majflt uint64 // (12) The number of major faults the process has made which have required loading a memory page from disk.

	Cmajflt uint64 // (13) The number of major faults that the process's waited-for children have made.

	Utime uint64 //  (14) Amount of time that this process has been scheduled in user mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).  This includes guest time, guest_time (time spent running a virtual CPU, see below), so that applications that are not aware of the guest time field do not lose that time from their calculations.

	Stime uint64 //  (15) Amount of time that this process has been scheduled in kernel mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).

	Cutime int64 //  (16) Amount of time that this process's waited-for children have been scheduled in user mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).  (See also times(2).)  This includes guest time, cguest_time (time spent running a virtual CPU, see below).

	Cstime int64 // (17) Amount of time that this process's waited-for children have been scheduled in kernel mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).

	Priority int64 // (18) (Explanation for Linux 2.6) For processes running a real-time scheduling policy (policy below; see sched_setscheduler(2)), this is the negated scheduling priority, minus one; that is, a number in the range -2 to -100, corresponding to real-time priorities 1 to 99.  For processes running under a non-real-time scheduling policy, this is the raw nice value (setpriority(2)) as represented in the kernel.  The kernel stores nice values as numbers in the range 0 (high) to 39 (low), corresponding to the user-visible nice range of -20 to 19.  Before Linux 2.6, this was a scaled value based on the scheduler weighting given to this process.
	Nice     int64 //   (19) The nice value (see setpriority(2)), a value in the range 19 (low priority) to -20 (high priority).

	NumThreads int64 // (20) Number of threads in this process (since Linux 2.6).  Before kernel 2.6, this field was hard coded to 0 as a placeholder for an earlier removed field.

	Itrealvalue int64 // (21) The time in jiffies before the next SIGALRM is sent to the process due to an interval timer.  Since kernel 2.6.17, this field is no longer maintained, and is hard coded as 0.

	Starttime uint64 //(22) The time the process started after system boot.  In kernels before Linux 2.6, this value was expressed in jiffies.  Since Linux 2.6, the value is expressed in clock ticks (divide by sysconf(_SC_CLK_TCK)).

	Vsize uint64 //   (23) Virtual memory size in bytes.

	Rss uint64 //     (24) Resident Set Size: number of pages the process has in real memory.  This is just the pages which count toward text, data, or stack space.  This does not include pages which have not been demand-loaded in, or which are swapped out.

	Rsslim uint64 // (25) Current soft limit in bytes on the rss of the process; see the description of RLIMIT_RSS in getrlimit(2).

	Startcode uint64 // (26) The address above which program text can run.

	Endcode uint64 // (27) The address below which program text can run.

	Startstack uint64 // (28) The address of the start (i.e., bottom) of the stack.

	Kstkesp uint64 // (29) The current value of ESP (stack pointer), as found in the kernel stack page for the process.

	Kstkeip uint64 // (30) The current EIP (instruction pointer).

	Signal uint64 //  (31) The bitmap of pending signals, displayed as a decimal number.  Obsolete, because it does not provide information on real-time signals; use /proc/[pid]/status instead.

	Blocked uint64 // (32) The bitmap of blocked signals, displayed as a decimal number.  Obsolete, because it does not provide information on real-time signals; use /proc/[pid]/status instead.

	Sigignore uint64 // (33) The bitmap of ignored signals, displayed as a decimal number.  Obsolete, because it does not provide information on real-time signals; use /proc/[pid]/status instead.

	Sigcatch uint64 // // (34) The bitmap of caught signals, displayed as a decimal number.  Obsolete, because it does not provide information on real-time signals; use /proc/[pid]/status instead.

	Wchan uint64 // (35) This is the "channel" in which the process is waiting.  It is the address of a system call, and can be looked up in a namelist if you need a textual name.  (If you have an up-to-date /etc/psdatabase, then try ps -l to see the WCHAN field in action.)

	Nswap uint64 // (36) Number of pages swapped (not maintained).

	Cnswap uint64 //  (37) Cumulative nswap for child processes (not maintained).

	ExitSignal int // (38) Signal to be sent to parent when we die.

	Processor int // (39) CPU number last executed on.

	RtPriority uint // (40) Real-time scheduling priority, a number in the range 1 to 99 for processes scheduled under a real-time policy, or 0, for non-real-time processes (see sched_setscheduler(2)).

	Policy uint // (41) Scheduling policy (see sched_setscheduler(2)).  Decode using the SCHED_* constants in linux/sched.h.

	DelayacctBlkioTicks uint64 // (42) Aggregated block I/O delays, measured in clock ticks (centiseconds).

	GuestTime uint64 // (43) Guest time of the process (time spent running a virtual CPU for a guest operating system), measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).

	CguestTime uint // (44) Guest time of the process's children, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).

}
