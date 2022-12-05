dd if=/dev/zero of=/tmp/loopback.file bs=1M count=10
losetup -fP /tmp/loopback.file
losetup -a
mkfs.ext4 /tmp/loopback.file
mkdir /mnt/{test1,test2}

# mount as read only
mount -o loop,ro /dev/loop8 /mnt/test1
mount -o loop,ro /dev/loop8 /mnt/test2

# mount as read write
mount -o loop,rw /dev/loop8 /mnt/test3
