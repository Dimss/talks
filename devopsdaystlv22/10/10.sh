######## 10.sh ########
mkdir -p "${LOOPBACK_DIR}"
dd if=/dev/zero of="${LOOPBACK_DIR}/${LOOPBACK_FILE}" bs=1M count=10
ls -allh "${LOOPBACK_DIR}"