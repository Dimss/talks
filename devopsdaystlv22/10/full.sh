# create empty file
LOOPBACK_FILE=loopback.file
LOOPBACK_DIR=/tmp/loopback-devices
mkdir -p ${LOOPBACK_FILE}
dd if=/dev/zero of="${LOOPBACK_DIR}/${LOOPBACK_FILE}" bs=1M count=10
ls -allh ${LOOPBACK_DIR}

# create virtual block device
losetup -fP "${LOOPBACK_DIR}/${LOOPBACK_FILE}"
losetup -a


# format virtual block device
mkfs.ext4 "${LOOPBACK_DIR}/${LOOPBACK_FILE}"

# mount as read only
mkdir -p /mnt/{test1,test2,test3}
DEVICE_ID=$(losetup -a  | grep ${LOOPBACK_FILE} | awk '{print $1}' | tr -d :)
mount -o loop,ro "${DEVICE_ID}" /mnt/test1
mount -o loop,ro "${DEVICE_ID}" /mnt/test2

# mount as read write
mount -o loop,rw "${DEVICE_ID}" /mnt/test3