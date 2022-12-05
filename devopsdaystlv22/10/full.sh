# create empty file
dd if=/dev/zero of=/tmp/loopback.file bs=1M count=10

# create virtual block device
losetup -fP /tmp/loopback.file
losetup -a

# format virtual block device
mkfs.ext4 /tmp/loopback.file

# mount as read only
mkdir -p /mnt/{test1,test2}
DEVICE_ID=$(losetup -a  | grep loopback.file | awk '{print $1}' | tr -d :)
mount -o loop,ro "${DEVICE_ID}" /mnt/test1
mount -o loop,ro "${DEVICE_ID}" /mnt/test2

# mount as read write
mount -o loop,rw "${DEVICE_ID}" /mnt/test3