docker run \
 --gpus=0  \
 -v $(pwd)/gpu-test.py:/tmp/gpu-test.py \
 tensorflow/tensorflow:latest-gpu \
 python /tmp/gpu-test.py
