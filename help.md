##  [Go] 解决missing go.sum entry for module providing package <package_name>
https://cloud.tencent.com/developer/article/1822013

## I got error "Failed setting memory limit: Permission denied - /tmp/warden/cgroup/memory/instance-18ensb4lmgh/memory.memsw.limit_in_bytes"
https://github.com/yudai/cf_nise_installer/issues/124

sudo echo "GRUB_CMDLINE_LINUX=\"cgroup_enable=memory swapaccount=1\"" >> /etc/default/grub 
sudo /usr/sbin/update-grub

## cobra 使用
https://www.cnblogs.com/chenqionghe/p/12661871.html
https://xcbeyond.cn/blog/golang/cobra-quick-start/

mkdir mydocker && cd mydocker
go mod init mydocker
cobra init

cobra add

打包成可执行文件
go build -o mydocker

## umount: device is busy. Why?

ubuntu@VM-8-17-ubuntu:~/vessel$ sudo ./vessel exec 085e81a6e2af echo "hello world"
Error: unable to mount proc to proc: device or resource busy

https://unix.stackexchange.com/questions/15024/umount-device-is-busy-why

## Git
git push origin master


## Go mod
go mod tidy // 删除未使用的依赖，并重新生成 go.mod 和 go.sum，下载缺失的依赖

## To do
写一个README.md

写删除镜像和删除容器的命令     

## Git - fatal: Unable to create '/path/my_project/.git/index.lock': File exists

https://stackoverflow.com/questions/7860751/git-fatal-unable-to-create-path-my-project-git-index-lock-file-exists