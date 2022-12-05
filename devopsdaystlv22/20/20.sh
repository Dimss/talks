docker run -it \
  --device=/dev/nvidia0:/dev/nvidia0 \
  --device=/dev/nvidiactl:/dev/nvidiactl \
  --volume=/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu \
  --volume=/usr/bin/nvidia-smi:/usr/bin/nvidia-smi \
  tensorflow/tensorflow:latest-gpu bash 
