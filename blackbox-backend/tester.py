import socket
import random
import struct
import time
import requests

IP = "143.47.241.215"
URL = "live.imperialrocketry.com"

UDP_PORT = 2052
WEBSITE_PORT = 5000

print(requests.post(url=f"http://{URL}/end"))
print(requests.post(url=f"http://{URL}/start"))

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

try:
    while True:
        data = struct.pack(">ifffffffffffff", time.monotonic_ns()//1_000_000, random.random(), random.random(), random.random(),random.random(), random.random(), random.random(),random.random(), random.random(), random.random(), (random.random()-0.5) * 10, random.random(), random.random(), random.random())
        sock.sendto(data, (IP, UDP_PORT))
        time.sleep(0.1)
except KeyboardInterrupt:
    pass

print(requests.post(url=f"http://{URL}/end"))