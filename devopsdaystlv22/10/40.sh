# mount as read write
LOOPBACK_FILE=loopback.file
DEVICE_ID=$(losetup -a  | grep ${LOOPBACK_FILE} | awk '{print $1}' | tr -d :)
mount -o loop,rw "${DEVICE_ID}" /mnt/test3