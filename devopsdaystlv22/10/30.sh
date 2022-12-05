######## 30.sh ########
# mount as read only
mkdir -p /mnt/{test1,test2,test3}
DEVICE_ID=$(losetup -a  | grep "${LOOPBACK_FILE}" | awk '{print $1}' | tr -d :)
mount -o loop,ro "${DEVICE_ID}" /mnt/test1
mount -o loop,ro "${DEVICE_ID}" /mnt/test2