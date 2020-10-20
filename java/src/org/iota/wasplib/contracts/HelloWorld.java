package org.iota.wasplib.contracts;

import org.iota.wasplib.client.context.ScContext;
import org.iota.wasplib.client.context.ScExports;

public class HelloWorld {
	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.Add("helloWorld");
	}

	//export helloWorld
	public static void helloWorld() {
		ScContext sc = new ScContext();
		sc.Log("Hello, world!");
	}
}
