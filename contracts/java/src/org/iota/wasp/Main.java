// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp;

import org.iota.wasp.contracts.dividend.lib.DividendThunk;
import org.iota.wasp.contracts.donatewithfeedback.lib.DonateWithFeedbackThunk;
import org.iota.wasp.contracts.erc20.lib.Erc20Thunk;
import org.iota.wasp.contracts.fairauction.lib.FairAuctionThunk;
import org.iota.wasp.contracts.fairroulette.lib.FairRouletteThunk;
import org.iota.wasp.contracts.helloworld.lib.HelloWorldThunk;
import org.iota.wasp.contracts.inccounter.lib.IncCounterThunk;
import org.iota.wasp.contracts.testcore.lib.TestCoreThunk;
import org.iota.wasp.contracts.tokenregistry.lib.TokenRegistryThunk;

public class Main {
	public static void main(String[] args) {
		DividendThunk.onLoad();
		DonateWithFeedbackThunk.onLoad();
		Erc20Thunk.onLoad();
		FairAuctionThunk.onLoad();
		FairRouletteThunk.onLoad();
		HelloWorldThunk.onLoad();
		IncCounterThunk.onLoad();
		TestCoreThunk.onLoad();
		TokenRegistryThunk.onLoad();
	}
}
