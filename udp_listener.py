import socket
import struct

hostname = socket.gethostname()
localIP = socket.gethostbyname(hostname)
localPort   = 20001

bufferSize  = 1024

# msgFromServer       = "Hello UDP Client"
# bytesToSend         = str.encode(msgFromServer) 

# Create a datagram socket

UDPServerSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip

UDPServerSocket.bind((localIP, localPort))
print("UDP server up and listening on ", localIP) 

# Listen for incoming datagrams

while(True):
    bytesAddressPair = UDPServerSocket.recvfrom(bufferSize)

    message = bytesAddressPair[0]
    address = bytesAddressPair[1]

    clientMsg = f"Message from Client: {message.hex()}"
    print(clientMsg, len(clientMsg)//2)
    (timestamp, or_x, or_y, or_z, gyro_x, gyro_y, gyro_z, acc_x, acc_y, acc_z, mag_x, mag_y, mag_z, lat, long, alt, temp, lux, pressure) = struct.unpack("fffffffffffffffffff", message)
    
    print(f"""
    timestamp = {timestamp}
    orientation = {or_x}, {or_y}, {or_z}
    gyro = {gyro_x}, {gyro_y}, {gyro_z}
    acc = {acc_x}, {acc_y}, {acc_z}
    mag = {mag_x}, {mag_y}, {mag_z}
    location = {lat}, {long}, {alt}
    temp = {temp}
    lux = {lux}
    pressure = {pressure}
    """)
    # clientIP  = "Client IP Address:{}".format(address)
    # print(clientIP) 

    # Sending a reply to client

    # UDPServerSocket.sendto(bytesToSend, address)
    UDPServerSocket.recvfrom(bufferSize)