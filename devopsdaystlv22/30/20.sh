for c in gpu-1 gpu-2; do
  docker run \
   --gpus=0  \
   --name $c \
   -d \
   -v $(pwd)/gpu-test.py:/tmp/gpu-test.py \
   tensorflow/tensorflow:latest-gpu \
   bash -c "sleep inf"
done