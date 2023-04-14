import tensorflow as tf
import time
from datetime import datetime as dt
gpus = tf.config.list_physical_devices('GPU')
tf.config.set_logical_device_configuration(gpus[0], [tf.config.LogicalDeviceConfiguration(memory_limit=2048)])
while True:
   print(f"[{dt.now()}] {tf.constant([[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]])}", flush=True)
   time.sleep(1)
