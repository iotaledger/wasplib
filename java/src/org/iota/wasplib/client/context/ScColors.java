package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableStringArray;

public class ScColors {
	ScImmutableStringArray colors;

	ScColors(ScImmutableStringArray colors) {
		this.colors = colors;
	}

	public int Length() {
		return colors.Length();
	}

	public ScColor GetColor(int index) {
		return new ScColor(colors.GetString(index).Value());
	}
}
