import socket
import struct
from dataclasses import dataclass
import dash
from dash import dcc
from dash import html
import plotly
from dash.dependencies import Input, Output
from collections import deque
import plotly.graph_objs as go

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
t = deque(maxlen=50)
y1 = deque(maxlen=50)
y2 = deque(maxlen=50)
y3 = deque(maxlen=50)

hostname = socket.gethostname()
localIP = socket.gethostbyname(hostname)
localPort   = 20001
bufferSize  = 1024

# Create a datagram socket
UDPServerSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)

# Bind to address and ip
UDPServerSocket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
UDPServerSocket.bind((localIP, localPort))
print("UDP server up and listening on ", localIP) 

external_stylesheets = ['https://codepen.io/chriddyp/pen/bWLwgP.css']
app = dash.Dash(__name__, external_stylesheets=external_stylesheets)

app.layout = html.Div(
    html.Div([
        html.H1('Android Sensor Data'),
        dcc.Graph(id='live-update-graph', animate = False),
        dcc.Interval(
            id='interval-component',
            interval = 500, # in milliseconds
            n_intervals=0
        )
    ])
)

@app.callback(Output('live-update-graph', 'figure'),
              Input('interval-component', 'n_intervals'))
def update_graph_live(n):

    # get data
    bytesAddressPair = UDPServerSocket.recvfrom(bufferSize)
    message = bytesAddressPair[0]
    (timestamp, or_x, or_y, or_z, gyro_x, gyro_y, gyro_z, acc_x, acc_y, acc_z, mag_x, mag_y, mag_z, lat, long, alt, temp, lux, pressure) = struct.unpack("fffffffffffffffffff", message)
    UDPServerSocket.recvfrom(bufferSize)

    # append to deque
    t.append(timestamp)
    y1.append(acc_x)
    y2.append(acc_y)
    y3.append(acc_z)
    # print(t)
    # print(y1)

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

    

    data = plotly.graph_objs.Scatter(
            x=list(t),
            y=list(y1),
            name='Scatter',
            mode= 'lines+markers'
    )

    return {'data': [data],
            'layout' : go.Layout(xaxis=dict(
                    range=[min(t),max(t)]),yaxis = 
                    dict(range = [min(y1),max(y1)]),
                    )}

if __name__ == '__main__':
    app.run_server(debug=True)
