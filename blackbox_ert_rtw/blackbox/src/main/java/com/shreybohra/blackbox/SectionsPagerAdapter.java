package com.shreybohra.blackbox;

import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentPagerAdapter;


class SectionsPagerAdapter extends FragmentPagerAdapter {

  private static final int NUMTABS = 1;
   public static int getNumTabs() { return NUMTABS; }

    public SectionsPagerAdapter(FragmentManager fm) {
        super(fm);
    }

   @Override
    public Fragment getItem(int position) {
        switch (position) {
			case 0:  return InfoFragment.newInstance();

			default: return null;
		}
	}

    @Override
    public int getCount() { return NUMTABS; }

    @Override
    public CharSequence getPageTitle(int position) {
        switch (position) {

			case 0:  return "Info";

			default: return null;
		}
	}

}