LOOPBACK_FILE=loopback.file
LOOPBACK_DIR=/tmp/loopback-devices
# create virtual block device
losetup -fP "${LOOPBACK_DIR}/${LOOPBACK_FILE}"
losetup -a
# format virtual block device
mkfs.ext4 "${LOOPBACK_DIR}/${LOOPBACK_FILE}"