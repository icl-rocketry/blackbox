import socket
import struct
from dataclasses import dataclass
import dash
from dash import dcc
from dash import html
import plotly
from dash.dependencies import Input, Output
from collections import deque


@dataclass
class SensorData:

    timestamp: float = 0.0
    or_x: float = 0.0
    or_y: float = 0.0
    or_z: float = 0.0
    gyro_x: float = 0.0
    gyro_y: float = 0.0
    gyro_z: float = 0.0
    acc_x: float = 0.0
    acc_y: float = 0.0
    acc_z: float = 0.0
    mag_x: float = 0.0
    mag_y: float = 0.0
    mag_z: float = 0.0
    lat: float = 0.0
    long: float = 0.0
    alt: float = 0.0
    temp: float = 0.0
    lux: float = 0.0
    pressure: float = 0.0

data = SensorData()

hostname = socket.gethostname()
localIP = socket.gethostbyname(hostname)
localPort   = 20001
bufferSize  = 1024

# Create a datagram socket
UDPServerSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip

UDPServerSocket.bind((localIP, localPort))
print("UDP server up and listening on ", localIP) 

external_stylesheets = ['https://codepen.io/chriddyp/pen/bWLwgP.css']

app = dash.Dash(__name__, external_stylesheets=external_stylesheets)
app.layout = html.Div(
    html.Div([
        html.H4('Android Sensor Data'),
        html.Div(id='live-update-text'),
        dcc.Graph(id='live-update-graph'),
        dcc.Interval(
            id='interval-component',
            interval = 50, # in milliseconds
            n_intervals=0
        )
    ])
)


while(True):

    bytesAddressPair = UDPServerSocket.recvfrom(bufferSize)
    message = bytesAddressPair[0]
    # address = bytesAddressPair[1]

    # clientMsg = f"Message from Client: {message.hex()}"
    # print(clientMsg, len(clientMsg)//2)
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
 
    # Simulink generated app sends an empty packet after data packet - need to discard
    UDPServerSocket.recvfrom(bufferSize)