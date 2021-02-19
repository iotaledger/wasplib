// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp;

import org.iota.wasp.contracts.dividend.lib.*;
import org.iota.wasp.contracts.donatewithfeedback.lib.*;
import org.iota.wasp.contracts.erc20.lib.*;
import org.iota.wasp.contracts.fairauction.lib.*;
import org.iota.wasp.contracts.fairroulette.lib.*;
import org.iota.wasp.contracts.helloworld.lib.*;
import org.iota.wasp.contracts.inccounter.lib.*;
import org.iota.wasp.contracts.testcore.lib.*;
import org.iota.wasp.contracts.tokenregistry.lib.*;

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
