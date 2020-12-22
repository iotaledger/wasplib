// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib;

import org.iota.wasplib.contracts.dividend.Dividend;
import org.iota.wasplib.contracts.donatewithfeedback.DonateWithFeedback;
import org.iota.wasplib.contracts.erc20.Erc20;
import org.iota.wasplib.contracts.fairauction.FairAuction;
import org.iota.wasplib.contracts.fairroulette.FairRoulette;
import org.iota.wasplib.contracts.helloworld.HelloWorld;
import org.iota.wasplib.contracts.inccounter.IncCounter;
import org.iota.wasplib.contracts.tokenregistry.TokenRegistry;

public class Main {
	public static void main(String[] args) {
		Dividend.onLoad();
		DonateWithFeedback.onLoad();
		Erc20.onLoad();
		FairAuction.onLoad();
		FairRoulette.onLoad();
		HelloWorld.onLoad();
		IncCounter.onLoad();
		TokenRegistry.onLoad();
	}
}
