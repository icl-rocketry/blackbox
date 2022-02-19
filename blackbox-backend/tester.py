import socket
import random
import struct
import time
import requests

IP = "127.0.0.1"

UDP_PORT = 1053
WEBSITE_PORT = 5000

print(requests.post(url=f"http://{IP}:{WEBSITE_PORT}/end"))
print(requests.post(url=f"http://{IP}:{WEBSITE_PORT}/start"))

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

try:
    while True:
        data = struct.pack(">ifffffffffffff", time.monotonic_ns()//1_000_000, random.random(), random.random(), random.random(),random.random(), random.random(), random.random(),random.random(), random.random(), random.random(), (random.random()-0.5) * 10, random.random(), random.random(), random.random())
        sock.sendto(data, (IP, UDP_PORT))
        time.sleep(0.1)
except KeyboardInterrupt:
    pass

print(requests.post(url=f"http://{IP}:{WEBSITE_PORT}/end"))