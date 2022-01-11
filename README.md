# blackbox
an Android Flight Data Recorder

The first release is being built using the [Simulink Support Package for Android Devices](https://uk.mathworks.com/help/supportpkg/android/examples/getting-started-with-android-devices.html). MATLAB/Simulink R2021a or later is required to open any files in this repository. 

## Simulink Android Setup Fix
There's a bug in the setup process of the Simulink Android package, caused by some renaming in the 31.0.0 release of the Android SDK Build Tools. 

The fix is to rename two files - follow the instructions [in this link](https://stackoverflow.com/a/68430992) for now, I'll create a bat file that automates it at some point.
