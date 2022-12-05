import tensorflow as tf
import time
from datetime import datetime as dt
while True:
   print(f"[{dt.now()}] {tf.config.list_physical_devices('GPU')}", flush=True)
   time.sleep(1)
