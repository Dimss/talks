basedir="/mnt"
for disk in disk1 disk2 disk3; do
  mkdir -p $basedir/$disk
  mount -t ramfs -o size=10m ramfs /mnt/$disk
done
ls -allh $basedir

