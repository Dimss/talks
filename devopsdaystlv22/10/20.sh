######## 20.sh ########
# create virtual block device
losetup -fP "${LOOPBACK_DIR}/${LOOPBACK_FILE}"
losetup -a

# format virtual block device
mkfs.ext4 "${LOOPBACK_DIR}/${LOOPBACK_FILE}"