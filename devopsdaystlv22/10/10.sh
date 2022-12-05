LOOPBACK_FILE=loopback.file
LOOPBACK_DIR=/tmp/loopback-devices
mkdir -p ${LOOPBACK_FILE}
dd if=/dev/zero of="${LOOPBACK_DIR}/${LOOPBACK_FILE}" bs=1M count=10
ls -allh ${LOOPBACK_DIR}