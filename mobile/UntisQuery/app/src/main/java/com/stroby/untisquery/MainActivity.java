package com.stroby.untisquery;

import android.os.Bundle;

import androidx.appcompat.app.AppCompatActivity;

import go.Seq;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        Seq.setContext(getApplicationContext());
    }
}