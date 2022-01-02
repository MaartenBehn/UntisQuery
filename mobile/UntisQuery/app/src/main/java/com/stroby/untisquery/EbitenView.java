package com.stroby.untisquery;

import android.content.Context;
import android.view.ViewGroup;

public class EbitenView extends ViewGroup {
    public EbitenView(Context context) {
        super(context);
    }

    // onErrorOnGameUpdate is called on the main thread when an error happens when updating a game.
    // You can define your own error handler, e.g., using Crashlytics, by overwriting this method.
    protected void onErrorOnGameUpdate(Exception e) {
        // Default error handling implementation.
    }

    // suspendGame suspends the game.
    // It is recommended to call this when the application is being suspended e.g.,
    // Activity's onPause is called.
    public void suspendGame() {
        // ...
    }

    // resumeGame resumes the game.
    // It is recommended to call this when the application is being resumed e.g.,
    // Activity's onResume is called.
    public void resumeGame() {
        // ...
    }

    @Override
    protected void onLayout(boolean b, int i, int i1, int i2, int i3) {

    }
}