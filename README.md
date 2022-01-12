# blackbox
an Android Flight Data Recorder

The first release is being built using the [Simulink Support Package for Android Devices](https://uk.mathworks.com/help/supportpkg/android/examples/getting-started-with-android-devices.html). MATLAB/Simulink R2021a or later is required to open any Simulink files in this repository. Android Studio is recommended for opening the files resulting from the build process.

## Current Capabilities
The App can currently send all supported sensor data to a host server at a fixed, predefined IP address and socket (currently set to the local IP of my laptop, socket 20001), recording data at a fixed data rate of 10 Hz.

A basic UDP server has been used to verify functionality - the python script is also available.

Tested working on a OnePlus 7T and Pixel 2 - both running Android 11.

## Simulink Android Setup Fix
There's a bug in the setup process of the Simulink Android package, caused by some renaming in the 31.0.0 release of the Android SDK Build Tools. 

The fix is to rename two files - follow the instructions [in this link](https://stackoverflow.com/a/68430992) for now, I'll create a bat file that automates it at some point.
