elkeid driver lkm  解码器
============

## 运行
```bash
$ go get github.com/sestack/elkeid-decoder
$ cd $GOPATH/github.com/sestack/elkeid-decoder/cmd
$ go run main.go
```

## 用法:
~~~
Usage: elkeid-decoder [options]:
  -filters string
    	事件过滤器,多个事件使用‘，’分割
  -list
    	打印支持的事件
  -version
    	打印版本信息
~~~

### 支持的事件列表
~~~
$ go run main.go -list
ID	Event
165	mount
703	interrupt_table_hook
356	memfd_create 
42	connect
59	execve
603	load_module
702	hidden_kernel_module
601	dns_query
602	create_file
604	update_cred
607	call_user_mode_helper_exec
610	usb_device_event
611	privilege_escalation
49	bind
82	rename
700	proc_file
701	syscall_table_hook
$
~~~

### 查看事件
~~~
$ go run main.go -filters 601,603
{"data_type":"dns_query","uid":0,"exe":"/usr/bin/ping","pid":63124,"ppid":61536,"pgid":63124,"tgid":63124,"sid":61536,"comm":"ping","node_name":"pi.example.com","sessionid":107,"pns":4026531836,"root_pns":4026531836,"query":"github.com","sa_family":2,"dip":"114.114.114.114","dport":53,"sip":"192.168.64.243","sport":54265,"opcode":0,"rcode":0}
{"data_type":"dns_query","uid":0,"exe":"/usr/bin/ping","pid":63124,"ppid":61536,"pgid":63124,"tgid":63124,"sid":61536,"comm":"ping","node_name":"pi.example.com","sessionid":107,"pns":4026531836,"root_pns":4026531836,"query":"github.com","sa_family":2,"dip":"114.114.114.114","dport":53,"sip":"192.168.64.243","sport":54265,"opcode":0,"rcode":0}
{"data_type":"load_module","uid":0,"exe":"/usr/bin/kmod","pid":63128,"ppid":61536,"pgid":63128,"tgid":63128,"sid":61536,"comm":"modprobe","node_name":"pi.example.com","sessionid":107,"pns":4026531836,"root_pns":4026531836,"ko_file":"ip_vs","pid_tree":"63128.modprobe\u003c61536.bash\u003c61535.sshd\u003c61533.sshd\u003c282.sshd\u003c1.systemd","run_path":"/root"}
$
~~~

## License

* Apache License Version 2.0


## Reference

* Elkeid.(https://github.com/bytedance/Elkeid/blob/main/driver/README-zh_CN.md)
