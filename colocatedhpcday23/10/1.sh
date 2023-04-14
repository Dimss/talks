docker run \
 --gpus=0  \
  tensorflow/tensorflow:latest-gpu \
  python -c \
   "import tensorflow as tf; gpus = tf.config.list_physical_devices('GPU'); tf.config.set_logical_device_configuration(gpus[0], [tf.config.LogicalDeviceConfiguration(memory_limit=2048)]); tf.constant([[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]])"


docker run \
 --gpus=0  \
 -eLD_LIBRARY_PATH=/lib/hostlibs \
 -v /lib/x86_64-linux-gnu/:/lib/hostlibs \
 tensorflow/tensorflow:latest-gpu \
 python -c \
"import tensorflow as tf; gpus = tf.config.list_physical_devices('GPU');
tf.config.set_logical_device_configuration(gpus[0], [tf.config.LogicalDeviceConfiguration(memory_limit=2048)]);
while True: tf.constant([[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]]); print(1);"
