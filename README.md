# golang-firewalld
用golang编写开启防火墙端口，利用port -p 端口，直接就可以打开端口，不需要输入很长的firewalld-cmd
解放双手，不再手动输入firewall-cmd命令，go build程序后，放在/usr/local/bin里， 比如说想打开65535端口 则:
<br>
 [root@localhost ~]port -p 65535
 <br>
 [root@localhost ~]# firewall-cmd --list-port
 <br>
80/tcp 3000/tcp 3478/tcp 3999/tcp 8080/tcp 9940/tcp 65535/tcp
