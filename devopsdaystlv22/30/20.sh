docker run \
 --gpus=0  \
  tensorflow/tensorflow:latest-gpu \
  python -c "import tensorflow as tf; import time; while True: tf.config.list_logical_devices('GPU');"