## Pprof性能分析实践

### 1. CPU性能分析

安装工具

```
yum install -y firefox
yum install -y sensible-utils
yum groupinstall "Gnome" -y
```

示例代码

```
[root@demo pp]# cat httpProfiling.go
package main

import (
    "fmt"
    "io"
    "net/http"
    _ "net/http/pprof"
    "time"
)


func sleep(sleepTime int) {
    time.Sleep(time.Duration(sleepTime) * time.Millisecond)
    fmt.Println("Slept for ", sleepTime, " Milliseconds")
}

func main() {
    Handler := func(w http.ResponseWriter, req *http.Request) {
        sleep(5)
        sleep(10)
        io.WriteString(w, "Memory Management Test")
    }
    http.HandleFunc("/", Handler)
    http.ListenAndServe(":1234", nil)
}

[root@demo pp]# go run httpProfiling.go
```

获取性能分析数据

```
[root@demo pp]# curl -s "localhost:1234/debug/pprof/profile?seconds=10" > out.dump
```

压测

```
[root@demo pp]# yum -y install httpd-tools
[root@demo pp]# ab -n 5000 -c 10 http://localhost:1234/
```

执行性能分析

```
[root@demo pp]# go tool pprof out.dump
File: httpProfiling
Type: cpu
Time: Apr 28, 2024 at 9:48pm (GMT)
Duration: 10.03s, Total samples = 780ms ( 7.78%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top10
Showing nodes accounting for 600ms, 76.92% of 780ms total
Showing top 10 nodes out of 178
      flat  flat%   sum%        cum   cum%
     390ms 50.00% 50.00%      390ms 50.00%  runtime/internal/syscall.Syscall6
      60ms  7.69% 57.69%       60ms  7.69%  runtime.futex
      40ms  5.13% 62.82%       40ms  5.13%  runtime.memmove
      30ms  3.85% 66.67%       30ms  3.85%  runtime.usleep
      20ms  2.56% 69.23%       20ms  2.56%  runtime.nanotime
      20ms  2.56% 71.79%       20ms  2.56%  time.Now
      10ms  1.28% 73.08%       20ms  2.56%  bufio.(*Reader).fill
      10ms  1.28% 74.36%       10ms  1.28%  context.(*cancelCtx).Done
      10ms  1.28% 75.64%       10ms  1.28%  indexbytebody
      10ms  1.28% 76.92%       10ms  1.28%  net/http.(*connReader).lock
(pprof) top10 -cum
Showing nodes accounting for 400ms, 51.28% of 780ms total
Showing top 10 nodes out of 178
      flat  flat%   sum%        cum   cum%
         0     0%     0%      490ms 62.82%  net/http.(*conn).serve
     390ms 50.00% 50.00%      390ms 50.00%  runtime/internal/syscall.Syscall6
         0     0% 50.00%      330ms 42.31%  syscall.RawSyscall6
         0     0% 50.00%      290ms 37.18%  syscall.Syscall
         0     0% 50.00%      190ms 24.36%  internal/poll.ignoringEINTRIO (inline)
         0     0% 50.00%      190ms 24.36%  runtime.mcall
      10ms  1.28% 51.28%      190ms 24.36%  runtime.schedule
         0     0% 51.28%      180ms 23.08%  internal/poll.(*FD).Write
         0     0% 51.28%      170ms 21.79%  runtime.findRunnable
         0     0% 51.28%      170ms 21.79%  syscall.Write (inline)
(pprof) svg
Generating report in profile001.svg
(pprof) svg sleep
Generating report in profile002.svg
(pprof) list sleep
Total: 780ms
ROUTINE ======================== main.sleep in /root/pp/httpProfiling.go
         0      120ms (flat, cum) 15.38% of Total
         .          .     12:func sleep(sleepTime int) {
         .       20ms     13:    time.Sleep(time.Duration(sleepTime) * time.Millisecond)
         .      100ms     14:    fmt.Println("Slept for ", sleepTime, " Milliseconds")
         .          .     15:}
         .          .     16:
         .          .     17:func main() {
         .          .     18:    Handler := func(w http.ResponseWriter, req *http.Request) {
         .          .     19:        sleep(5)
ROUTINE ======================== runtime.futexsleep in /usr/local/go/src/runtime/os_linux.go
      10ms       30ms (flat, cum)  3.85% of Total
         .          .     62:func futexsleep(addr *uint32, val uint32, ns int64) {
         .          .     63:   // Some Linux kernels have a bug where futex of
         .          .     64:   // FUTEX_WAIT returns an internal error code
         .          .     65:   // as an errno. Libpthread ignores the return value
         .          .     66:   // here, and so can we: as it says a few lines up,
         .          .     67:   // spurious wakeups are allowed.
         .          .     68:   if ns < 0 {
      10ms       30ms     69:           futex(unsafe.Pointer(addr), _FUTEX_WAIT_PRIVATE, val, nil, nil, 0)
         .          .     70:           return
         .          .     71:   }
         .          .     72:
         .          .     73:   var ts timespec
         .          .     74:   ts.setNsec(ns)
ROUTINE ======================== runtime.notesleep in /usr/local/go/src/runtime/lock_futex.go
      10ms       40ms (flat, cum)  5.13% of Total
         .          .    148:func notesleep(n *note) {
         .          .    149:   gp := getg()
         .          .    150:   if gp != gp.m.g0 {
         .          .    151:           throw("notesleep not on g0")
         .          .    152:   }
         .          .    153:   ns := int64(-1)
         .          .    154:   if *cgo_yield != nil {
         .          .    155:           // Sleep for an arbitrary-but-moderate interval to poll libc interceptors.
         .          .    156:           ns = 10e6
         .          .    157:   }
         .          .    158:   for atomic.Load(key32(&n.key)) == 0 {
         .          .    159:           gp.m.blocked = true
         .       30ms    160:           futexsleep(key32(&n.key), 0, ns)
         .          .    161:           if *cgo_yield != nil {
         .          .    162:                   asmcgocall(*cgo_yield, nil)
         .          .    163:           }
      10ms       10ms    164:           gp.m.blocked = false
         .          .    165:   }
         .          .    166:}
         .          .    167:
         .          .    168:// May run with m.p==nil if called from notetsleep, so write barriers
         .          .    169:// are not allowed.
ROUTINE ======================== runtime.usleep in /usr/local/go/src/runtime/sys_linux_amd64.s
      30ms       30ms (flat, cum)  3.85% of Total
         .          .    120:TEXT runtime·usleep(SB),NOSPLIT,$16
         .          .    121:   MOVL    $0, DX
         .          .    122:   MOVL    usec+0(FP), AX
         .          .    123:   MOVL    $1000000, CX
         .          .    124:   DIVL    CX
         .          .    125:   MOVQ    AX, 0(SP)
         .          .    126:   MOVL    $1000, AX       // usec to nsec
         .          .    127:   MULL    DX
         .          .    128:   MOVQ    AX, 8(SP)
         .          .    129:
         .          .    130:   // nanosleep(&ts, 0)
         .          .    131:   MOVQ    SP, DI
         .          .    132:   MOVL    $0, SI
         .          .    133:   MOVL    $SYS_nanosleep, AX
         .          .    134:   SYSCALL
      30ms       30ms    135:   RET
         .          .    136:
         .          .    137:TEXT runtime·gettid(SB),NOSPLIT,$0-4
         .          .    138:   MOVL    $SYS_gettid, AX
         .          .    139:   SYSCALL
         .          .    140:   MOVL    AX, ret+0(FP)

```

在性能分析中，"flat" 和 "cumulative" (简称 "cum") 是两个常用的概念，通常用于分析函数或方法的性能。

Flat Profile（平坦分析）：这个指标表示每个函数或方法本身的性能情况，不考虑其调用其他函数或方法所花费的时间。在 flat profile 中，你可以看到每个函数或方法自身所消耗的 CPU 时间、内存等资源。
Cumulative Profile（累积分析）：与 flat 相反，这个指标考虑了一个函数或方法以及它调用的所有其他函数或方法所消耗的总时间。换句话说，cumulative profile 显示了函数调用树中每个函数的总时间，包括它的子函数调用。

总的来说，flat profile 主要用于查看每个函数或方法的个体性能，而 cumulative profile 则用于查看函数调用链中的性能分布情况，以及了解整体调用树的性能瓶颈。

### 2. MEM性能分析

Go 中最常见的内存泄漏之一是 goroutine 的无限制创建。当您超载无缓冲的通道或您的抽象有大量并发产生未完成的新 goroutine 时，这种情况经常发生。 Goroutine 的占用空间非常小，系统通常可以产生大量的 Goroutine，但它们最终会存在一个上限，当尝试在生产环境中对程序进行故障排除时，找到该上限会变得很费力。

```
[root@demo pp]# cat memoryLeak.go
package main

import (
 "fmt"
 "net/http"

 _ "net/http/pprof"
 "runtime"
 "time"
)

func main() {
 http.HandleFunc("/leak", leakyAbstraction)
 http.ListenAndServe("localhost:6060", nil)
}

func leakyAbstraction(w http.ResponseWriter, r *http.Request) {
 ch := make(chan string)

 for {
   fmt.Fprintln(w, "Number of Goroutines: ", runtime.NumGoroutine())
   go func() { ch <- wait() }()
 }
}

func wait() string {
 time.Sleep(5 * time.Microsecond)
 return "Hello Gophers!"
}

[root@demo pp]# go run memoryLeak.go
```

```
[root@demo pp]# curl -s "localhost:6060/debug/pprof/heap?seconds=30" > out.dump
```

```
[root@demo pp]# curl localhost:6060/leak
```

```
[root@demo pp]# go tool pprof -alloc_space  out.dump
File: memoryLeak
Type: alloc_space
Time: Apr 28, 2024 at 10:45pm (GMT)
Duration: 30.74s, Total samples = 528.83MB
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 投
(pprof) top10
Showing nodes accounting for 526.24MB, 99.51% of 528.83MB total
Dropped 19 nodes (cum <= 2.64MB)
Showing top 10 nodes out of 13
      flat  flat%   sum%        cum   cum%
  335.13MB 63.37% 63.37%   335.13MB 63.37%  runtime.malg
   69.01MB 13.05% 76.42%   132.51MB 25.06%  main.leakyAbstraction.func1
   63.50MB 12.01% 88.43%    63.50MB 12.01%  time.Sleep
   41.10MB  7.77% 96.20%    41.10MB  7.77%  runtime.allgadd
   17.50MB  3.31% 99.51%    17.50MB  3.31%  main.leakyAbstraction
         0     0% 99.51%    63.50MB 12.01%  main.wait (inline)
         0     0% 99.51%    20.09MB  3.80%  net/http.(*ServeMux).ServeHTTP
         0     0% 99.51%    20.09MB  3.80%  net/http.(*conn).serve
         0     0% 99.51%    20.09MB  3.80%  net/http.HandlerFunc.ServeHTTP
         0     0% 99.51%    20.09MB  3.80%  net/http.serverHandler.ServeHTTP
(pprof) top10 -cum
Showing nodes accounting for 508.74MB, 96.20% of 528.83MB total
Dropped 19 nodes (cum <= 2.64MB)
Showing top 10 nodes out of 13
      flat  flat%   sum%        cum   cum%
         0     0%     0%   376.23MB 71.14%  runtime.newproc.func1
         0     0%     0%   376.23MB 71.14%  runtime.newproc1
         0     0%     0%   376.23MB 71.14%  runtime.systemstack
  335.13MB 63.37% 63.37%   335.13MB 63.37%  runtime.malg
   69.01MB 13.05% 76.42%   132.51MB 25.06%  main.leakyAbstraction.func1
         0     0% 76.42%    63.50MB 12.01%  main.wait (inline)
   63.50MB 12.01% 88.43%    63.50MB 12.01%  time.Sleep
   41.10MB  7.77% 96.20%    41.10MB  7.77%  runtime.allgadd
         0     0% 96.20%    20.09MB  3.80%  net/http.(*ServeMux).ServeHTTP
         0     0% 96.20%    20.09MB  3.80%  net/http.(*conn).serve
(pprof) svg
Generating report in profile001.svg
(pprof) svg leakyAbstraction
Generating report in profile002.svg
(pprof) list leakyAbstraction
Total: 528.83MB
ROUTINE ======================== main.leakyAbstraction in /root/pp/memoryLeak.go
   17.50MB    17.50MB (flat, cum)  3.31% of Total
         .          .     17:func leakyAbstraction(w http.ResponseWriter, r *http.Request) {
         .          .     18: ch := make(chan string)
         .          .     19:
         .          .     20: for {
    6.50MB     6.50MB     21:   fmt.Fprintln(w, "Number of Goroutines: ", runtime.NumGoroutine())
      11MB       11MB     22:   go func() { ch <- wait() }()
         .          .     23: }
         .          .     24:}
         .          .     25:
         .          .     26:func wait() string {
         .          .     27: time.Sleep(5 * time.Microsecond)
ROUTINE ======================== main.leakyAbstraction.func1 in /root/pp/memoryLeak.go
   69.01MB   132.51MB (flat, cum) 25.06% of Total
   69.01MB   132.51MB     22:   go func() { ch <- wait() }()
         .          .     23: }
         .          .     24:}
         .          .     25:
         .          .     26:func wait() string {
         .          .     27: time.Sleep(5 * time.Microsecond)
```
