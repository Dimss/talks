for c in gpu-1 gpu-2; do
  docker run \
   --gpus=0  \
   -eLD_LIBRARY_PATH=/lib/hostlibs \
   -v /lib/x86_64-linux-gnu/:/lib/hostlibs \
   --name $c \
   -d \
   -v $(pwd)/gpu-test.py:/tmp/gpu-test.py \
   tensorflow/tensorflow:latest-gpu \
   python /tmp/gpu-test.py
done