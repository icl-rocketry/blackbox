package com.shreybohra.blackbox;

import android.Manifest;
import android.content.pm.PackageManager;
import android.os.Bundle;
import android.widget.Toast;
import android.util.Log;
import java.util.ArrayList;
import android.content.res.Configuration;
import androidx.core.app.ActivityCompat;
import androidx.fragment.app.Fragment;
import androidx.core.content.ContextCompat;
import com.google.android.material.tabs.TabLayout;
import androidx.viewpager.widget.ViewPager;
import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import android.content.Context;
import android.hardware.Sensor;
import android.hardware.SensorEvent;
import android.hardware.SensorEventListener;
import android.hardware.SensorManager;

public class blackbox extends AppCompatActivity implements SensorEventListener, OnFragmentInteractionListener {
    private SectionsPagerAdapter mSectionsPagerAdapter;

    private ViewPager mViewPager;
    private InfoFragment infoFragment = null;
     private float mLightData = 0.0f;
     private float[] mGyroscopeData = { 0.0f, 0.0f, 0.0f };
     private float[] mAccelerometerData = { 0.0f, 0.0f, 0.0f };
     private float mTemperatureData = 0.0f;
     private float mPressureData = 0.0f;
     private float[] mMagnetometerData = { 0.0f, 0.0f, 0.0f };
   private final float[] mRotationMatrix = new float[9];
private final float[] mOrientationAngles = new float[3];
     private SensorManager mSensorManager;
     private GPSHandler mGPSHandler;
    private static final int MY_PERMISSIONS_REQUEST_FINE_LOCATION = 300;
    private boolean isFineLocationPermissionGranted = false;
    private boolean isFineLocationPermissionRequested = false;

     private void registerSensorManager() {
        mSensorManager.registerListener(this,
            mSensorManager.getDefaultSensor(Sensor.TYPE_ACCELEROMETER),
            SensorManager.SENSOR_DELAY_FASTEST);
        mSensorManager.registerListener(this,
            mSensorManager.getDefaultSensor(Sensor.TYPE_MAGNETIC_FIELD),
            SensorManager.SENSOR_DELAY_FASTEST);
        mSensorManager.registerListener(this,
            mSensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE),
            SensorManager.SENSOR_DELAY_FASTEST);
        mSensorManager.registerListener(this,
            mSensorManager.getDefaultSensor(Sensor.TYPE_LIGHT),
            SensorManager.SENSOR_DELAY_FASTEST);
        mSensorManager.registerListener(this,
            mSensorManager.getDefaultSensor(Sensor.TYPE_AMBIENT_TEMPERATURE),
            SensorManager.SENSOR_DELAY_FASTEST);
        mSensorManager.registerListener(this,
            mSensorManager.getDefaultSensor(Sensor.TYPE_PRESSURE),
            SensorManager.SENSOR_DELAY_FASTEST);
     }

    private boolean checkIfAllPermissionsGranted()
    {
        return true && isFineLocationPermissionGranted;
    }
    private void requestPermission() {
        String permissionRationale = "";
        // Here, thisClass is the current activity
        //request for fine location
        if (ContextCompat.checkSelfPermission(thisClass,
                Manifest.permission.ACCESS_FINE_LOCATION)
                != PackageManager.PERMISSION_GRANTED) {

            // Permission is not granted. Should we show an explanation?
            if (ActivityCompat.shouldShowRequestPermissionRationale(thisClass,
                    Manifest.permission.ACCESS_FINE_LOCATION)) {
                permissionRationale += "Access fine location, ";
            } else {
                // No explanation needed; request the permission
                if (!isFineLocationPermissionRequested) {
                    isFineLocationPermissionRequested = true;
                    ActivityCompat.requestPermissions(thisClass,
                            new String[]{Manifest.permission.ACCESS_FINE_LOCATION},
                            MY_PERMISSIONS_REQUEST_FINE_LOCATION);
                    return;
                }
            }
        } else {
            // Permission has already been granted
            isFineLocationPermissionGranted = true;
        }
        if (!permissionRationale.isEmpty())
            if (infoFragment != null) {
                infoFragment.updateModelInfo(permissionRationale + "permission not granted. Model cannot start.");
            }
    }
    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        //Uncomment the following line to specify a custom App Title
        //toolbar.setTitle("My custom Title");
        setSupportActionBar(toolbar);

        // Create a FragmentPagerAdapter that returns individual fragments
        mSectionsPagerAdapter = new SectionsPagerAdapter(getSupportFragmentManager());

        // Set up the ViewPager with the sections adapter.
        mViewPager = (ViewPager) findViewById(R.id.container);
        mViewPager.setAdapter(mSectionsPagerAdapter);
        mViewPager.setOffscreenPageLimit(SectionsPagerAdapter.getNumTabs()-1);

        TabLayout tabLayout = (TabLayout) findViewById(R.id.tabs);
        tabLayout.setupWithViewPager(mViewPager);

        // Initiate the SensorManager
        mSensorManager = (SensorManager) getSystemService(Context.SENSOR_SERVICE);
         mGPSHandler = new GPSHandler(this);
        thisClass = this;
     }

    private blackbox thisClass;
    private final Thread BgThread = new Thread() {
    @Override
    public void run() {
            String argv[] = new String[] {"MainActivity","blackbox"};
            naMain(argv, thisClass);
        }
    };

    public void flashMessage(final String inMessage) {
        runOnUiThread(new Runnable() {
              public void run() {
                    Toast.makeText(getBaseContext(), inMessage, Toast.LENGTH_SHORT).show();
              }
        });
    }

    protected void onDestroy() {
         if (BgThread.isAlive())
             naOnAppStateChange(6);
         super.onDestroy();
         System.exit(0); //to kill all our threads.
    }

	@Override
    public void onAttachFragment(Fragment fragment) {
        super.onAttachFragment(fragment);
        if (fragment instanceof InfoFragment) {
            this.infoFragment = (InfoFragment) fragment;
            infoFragment.setFragmentInteractionListener(this);
            requestPermission();
        }
    }

	@Override
    public void onFragmentCreate(String name) {

    }

    @Override
    public void onFragmentStart(String name) {
        switch (name) {
            case "Info":
                break;
            default:
                break;
    }
    }

    @Override
    public void onFragmentResume(String name) {
        switch (name) {
            case "Info":
                if (checkIfAllPermissionsGranted()){
                    if (!BgThread.isAlive()) {
                        BgThread.start();
                    }
                }
                break;
            default:
                break;
        }
    }

    @Override
    public void onFragmentPause(String name) {
    }
    @Override
    protected void onResume() {
         super.onResume();
         if (BgThread.isAlive())
             naOnAppStateChange(3);
         registerSensorManager();
    }

    @Override
    protected void onPause() {
        if (BgThread.isAlive())
            naOnAppStateChange(4);
         mSensorManager.unregisterListener(this);
        super.onPause();
    }

    @Override
    public void onConfigurationChanged(Configuration newConfig) {
        super.onConfigurationChanged(newConfig);
    }

    @Override
    public void onRequestPermissionsResult(int requestCode,
                                           String permissions[], int[] grantResults) {
        switch (requestCode) {
            case MY_PERMISSIONS_REQUEST_FINE_LOCATION:
                // If request is cancelled, the result arrays are empty.
                if (grantResults.length > 0
                        && grantResults[0] == PackageManager.PERMISSION_GRANTED) {
                    // permission was granted, yay! Do related task you need to do.
                    isFineLocationPermissionGranted = true;
                    mGPSHandler = new GPSHandler(thisClass);
                } else {
                    // permission denied, boo!
                    flashMessage("Access location Permission not granted");
                }
                isFineLocationPermissionRequested = false;
                break;

            // other case lines to check for other
            // permissions this app might request.
        }
        if (!checkIfAllPermissionsGranted()) {
            requestPermission();
        }
    }

    @Override
    public void onAccuracyChanged(Sensor sensor, int accuracy) {
    }

    @Override
    public void onSensorChanged(SensorEvent event) {
        float [] values = event.values;
        //Comment out if you want to log the data in logcat
        //String logMessage = String.format("%d: 0'%g'", event.sensor.getType(), values[0]);
        //Log.d("Sensor Data IN:", logMessage);
        switch(event.sensor.getType()) {
            case Sensor.TYPE_LIGHT:
                mLightData = values[0];
                break;
            case Sensor.TYPE_GYROSCOPE:
                mGyroscopeData[0] = values[0];
                mGyroscopeData[1] = values[1];
                mGyroscopeData[2] = values[2];
                break;
            case Sensor.TYPE_ACCELEROMETER:
                mAccelerometerData[0] = values[0];
                mAccelerometerData[1] = values[1];
                mAccelerometerData[2] = values[2];
                break;
            case Sensor.TYPE_MAGNETIC_FIELD:
                mMagnetometerData[0] = values[0];
                mMagnetometerData[1] = values[1];
                mMagnetometerData[2] = values[2];
                break;
            case Sensor.TYPE_AMBIENT_TEMPERATURE:
                mTemperatureData = values[0];
                break;
            case Sensor.TYPE_PRESSURE:
                mPressureData = values[0];
                break;
        }
    }

    // Get SensorEvent Data throws exception if the data is null
    public float getLightData() {
        return mLightData;
    }

    public float[] getGyroscopeData() {
        return mGyroscopeData;
    }

    public float[] getAccelerometerData() {
        return mAccelerometerData;
    }

    public float[] getMagnetometerData() {
        return mMagnetometerData;
    }

    public float[] getOrientationData() {
        mSensorManager.getRotationMatrix(mRotationMatrix, null,mAccelerometerData, mMagnetometerData);
        mSensorManager.getOrientation(mRotationMatrix, mOrientationAngles);
        mOrientationAngles[0] = (float)Math.toDegrees(mOrientationAngles[0]);
        mOrientationAngles[1] = (float)Math.toDegrees(mOrientationAngles[1]);
        mOrientationAngles[2] = (float)Math.toDegrees(mOrientationAngles[2]);
        return mOrientationAngles;
    }

    public float getTemperatureData() {
        return mTemperatureData;
    }

    public float getPressureData() {
        return mPressureData;
    }

    // Get GPS Data if GPS is enabled. Otherwise return 0,0
    public double[] getGPSData() {
        return mGPSHandler.getLocationData();
    }

    private native int naMain(String[] argv, blackbox pThis);
    private native void naOnAppStateChange(int state);
    static {
        System.loadLibrary("blackbox");
    }

}
