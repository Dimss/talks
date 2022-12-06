import tensorflow as tf
import time
from datetime import datetime as dt
while True:
   print(f"[{dt.now()}] {tf.config.list_physical_devices('GPU')}", flush=True)
   print(tf.reduce_sum(tf.random.normal([1000, 1000])))
   time.sleep(1)
