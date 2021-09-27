<script lang="ts">
  export const name = 'app';

  import { Base58 } from './wasp_client/crypto/base58';
  import {
    BasicClient,
    ColorCollection,
    Colors,
    HName,
    IUnspentOutput,
    IUnspentOutputAddress,
    PoWWorkerManager,
    SimpleBufferCursor,
  } from './wasp_client';
  import { Buffer } from './wasp_client/buffer';
  import { FairRouletteService } from './fairroulette_client';
  import type { Bet } from './fairroulette_client';
  import { onMount } from 'svelte';
  import { Seed } from './wasp_client/crypto/seed';
  import config, { chainId } from '../config.dev';
  import Roulette from './Roulette.svelte';

  import { seed, addressIndex, keyPair, address } from './store';
  import { WalletService } from './wasp_client';
  import type { IOnLedger } from './wasp_client/binary_models/IOnLedger';
  import { OnLedger } from './wasp_client/binary_models/on_ledger';
  import { BasicWallet } from './wasp_client/basic_wallet';

  let fundsUpdaterHandle;

  let client: BasicClient;
  let walletService: WalletService;
  let fairRouletteService: FairRouletteService;

  const powManager: PoWWorkerManager = new PoWWorkerManager();

  const view = {
    seedString: '',
    balance: 0n,
    timestamp: 0,

    showController: true,
    isWorking: false,

    round: {
      active: false,
      number: 0n,
      eventList: [] as string[],
      betSelection: 1,
      winningNumber: 0n,
      startedAt: 0,
    },
  };

  // Entrypoint
  async function initialize() {
    log('[PAGE] loading');

    if (config.seed) {
      $seed = Base58.decode(config.seed);
    } else {
      $seed = Seed.generate();
    }

    //TODO: Remove this at some point.
    if (!config.chainId && config.chainResolverUrl) {
      try {
        const response = await fetch(config.chainResolverUrl);
        const content = await response.json();

        config.chainId = content.chainId;
      } catch (ex) {
        log(ex.message);
      }
    }

    view.seedString = Base58.encode($seed);

    client = new BasicClient({
      GoShimmerAPIUrl: config.goshimmerApiUrl,
      WaspAPIUrl: config.waspApiUrl,
      SeedUnsafe: $seed,
    });

    fairRouletteService = new FairRouletteService(client, config.chainId);
    walletService = new WalletService(client);

    powManager.load('/build/pow.worker.js');

    subscribeToRouletteEvents();
    setAddress($addressIndex);
    updateFunds();

    startFundsUpdater();

    // The best solution would be to call all of them in parallel. This is currently not possible.
    // As those requests can fail in certain cases, we need to wrap them in exception handlers,
    // to make sure that the other requests are being sent and that the page properly loads.
    const requests = [
      () => fairRouletteService.getRoundStatus().then((x) => (view.round.active = x == 1)),
      () => fairRouletteService.getRoundNumber().then((x) => (view.round.number = x)),
      () => fairRouletteService.getLastWinningNumber().then((x) => (view.round.winningNumber = x)),
      () => fairRouletteService.getRoundStartedAt().then((x) => (view.round.startedAt = x)),
    ];

    for (let request of requests) {
      await request().catch((e) => log(`[ERROR] ${e.message}`));
    }

    log('[PAGE] loaded');
  }

  onMount(initialize);
  // /Entrypoint

  function log(text: string) {
    view.round.eventList.push(`${new Date().toLocaleTimeString()} | ${text}`);
  }

  function setAddress(index: number) {
    $addressIndex = index;

    $address = Seed.generateAddress($seed, $addressIndex);
    $keyPair = Seed.generateKeyPair($seed, $addressIndex);
  }

  function createNewAddress() {
    $addressIndex++;
    setAddress($addressIndex);
  }

  async function updateFunds() {
    try {
      view.timestamp = Date.now() / 1000;
      view.balance = await walletService.getFunds($address, Colors.IOTA_COLOR_STRING);
    } catch (ex) {
      view.balance = 0n;
    }
  }

  function startFundsUpdater() {
    if (fundsUpdaterHandle) {
      fundsUpdaterHandle = clearInterval(fundsUpdaterHandle);
    }

    fundsUpdaterHandle = setInterval(updateFunds, 1000);
  }

  async function placeBet() {
    view.isWorking = true;
    try {
      await fairRouletteService.placeBetOnLedger(
        $keyPair,
        $address,
        view.round.betSelection,
        1234n
      );
    } catch (ex) {
      log(ex.message);

      throw ex;
    }
    view.isWorking = false;
  }

  async function sendFaucetRequest() {
    view.isWorking = true;

    const faucetRequestResult = await walletService.getFaucetRequest($address);

    // In this example a difficulty of 20 is enough, might need a retune for prod to 21 or 22
    faucetRequestResult.faucetRequest.nonce = await powManager.requestProofOfWork(
      20,
      faucetRequestResult.poWBuffer
    );

    try {
      await client.sendFaucetRequest(faucetRequestResult.faucetRequest);
    } catch (ex) {
      log(ex.message);
    }

    view.isWorking = false;
  }

  function calculateRoundLengthLeft(timestamp: number) {
    const roundStarted = view.round.startedAt;

    if (roundStarted == 0) {
      return 0;
    }

    const diff = timestamp - roundStarted;

    // TODO: Explain.
    const executionCompensation = 5;
    const roundTimeLeft = Math.round(
      fairRouletteService.roundLength + executionCompensation - diff
    );

    if (roundTimeLeft <= 0) {
      return 0;
    }
    return roundTimeLeft;
  }

  function subscribeToRouletteEvents() {
    fairRouletteService.on('roundStarted', (timestamp) => {
      view.round.active = true;
      view.round.startedAt = timestamp;
      log('[ROUND] started');
    });

    fairRouletteService.on('roundStopped', () => {
      view.round.active = false;
      log('[ROUND] ended');
    });

    fairRouletteService.on('roundNumber', (roundNumber: bigint) => {
      view.round.number = roundNumber;
      log(`[ROUND] Current round number: ${roundNumber}`);
    });

    fairRouletteService.on('winningNumber', (winningNumber: bigint) => {
      view.round.winningNumber = winningNumber;
      log(`[ROUND] The winning number was: ${winningNumber}`);
    });

    fairRouletteService.on('betPlaced', (bet: Bet) => {
      log(`[BET] Bet placed from ${bet.better} on ${bet.betNumber} with ${bet.amount}`);
    });

    fairRouletteService.on('payout', (bet: Bet) => {
      log(`[WIN] Payout for ${bet.better} with ${bet.amount}`);
    });
  }

  function isWealthy(balance: bigint) {
    return balance >= 200;
  }
</script>

<main>
  <div class="header">
    <ul>
      <li>
        <span class="header_balance_title">Balance</span>
        <span class="header_balance_text">{view.balance}i</span>
      </li>
      <li>
        <span class="header_seed_title">Seed</span>
        <span class="header_seed_text">{view.seedString}</span>
      </li>
      <li>
        <span class="header_address_title">Address</span>
        <span class="header_address_text">{$address}</span>
      </li>
    </ul>
  </div>

  {#if view.isWorking}
    <div class="loading_dim">
      <div class="loading_wrapper">
        <div class="loading_logo" />
        <div class="loading_text">Loading and working</div>
      </div>
    </div>
  {/if}

  {#if view.showController}
    <div class="welcome_screen">
      <div class="request_funds" class:disabled={isWealthy(view.balance)}>
        <div class="request_funds_text">1 | Request funds</div>
        <div class="request_funds_image" />
        <button
          class="request_funds_button"
          disabled={isWealthy(view.balance)}
          on:click={() => sendFaucetRequest()}>Request funds</button
        >
      </div>

      <div class="place_bet" class:disabled={!isWealthy(view.balance)}>
        <div class="place_bet_text">2 | Place a bet</div>
        <div class="place_bet_image" />

        <div class="bet_selection">
          <input
            type="radio"
            bind:group={view.round.betSelection}
            name="bet_number"
            value="1"
            checked
          />
          <label for="1">1</label>

          <input type="radio" bind:group={view.round.betSelection} name="bet_number" value="2" />
          <label for="2">2</label>

          <input type="radio" bind:group={view.round.betSelection} name="bet_number" value="3" />
          <label for="3">3</label>

          <input type="radio" bind:group={view.round.betSelection} name="bet_number" value="4" />
          <label for="4">4</label>

          <input type="radio" bind:group={view.round.betSelection} name="bet_number" value="5" />
          <label for="5">5</label>
        </div>

        <button
          class="submit_bet_button"
          on:click={() => placeBet()}
          disabled={!isWealthy(view.balance)}>Submit bet</button
        >
      </div>
    </div>

    <div class="welcome_screen" />
  {/if}

  <div class="content">
    <div class="wheel_container">
      <Roulette width="300" height="300" spin={view.round.active} />

      <div class="wheel_status">
        <div>Round {view.round.active ? 'started' : 'stopped'}</div>
        <div>Round Nr. {view.round.number}</div>
        <div>
          Winning Nr. {view.round.active ? 'Undecided' : view.round.winningNumber}
        </div>
        <div>
          Round TS: {view.round.startedAt}
        </div>
        <div>
          Remaining seconds: {calculateRoundLengthLeft(view.timestamp)}
        </div>
      </div>
    </div>

    <div class="log">
      <textarea rows="4" cols="50">{view.round.eventList.join('\n')}</textarea>
    </div>
  </div>
</main>

<style>
  textarea {
    width: 100%;
    height: 100%;
  }

  .log {
    width: 75%;
    margin-left: 35px;
  }

  .wheel_container {
    border: 1px solid white;
    border-radius: 4px;
    padding: 5px;
  }

  .wheel_status {
    margin-top: 15px;
    width: 100%;
    height: 100px;
    font-size: 26px;
    color: gray;
  }

  .content {
    display: flex;
    width: 75%;
    margin: 15px auto;
  }

  main {
    width: 100%;
    height: 100%;
  }

  div.header {
    width: 100%;
    background-color: rgba(72, 87, 118, 0.2);
    height: 45px;
  }

  div.header ul {
    margin: 0;
    padding: 0;
    display: flex;
    list-style: none;
    height: 100%;
  }

  div.header li {
    height: 100%;
    margin: auto;
    display: flex;
    padding: 0;
    align-items: center;
    justify-content: center;
    color: rgba(255, 255, 255, 0.9);

    text-shadow: 0px 2px 2px rgba(13, 15, 4, 0.42);
  }

  .header_balance_title {
    font-size: 24px;
  }

  .header_balance_text {
    padding-left: 5px;
    font-size: 26px;
  }

  .header_seed_title,
  .header_address_title {
    font-size: 20px;
  }

  .header_seed_text,
  .header_address_text {
    font-size: 16px;
    padding-left: 5px;
  }

  .loading_dim {
    background: rgba(0, 0, 0, 0.5);
    width: 100%;
    height: 100%;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    z-index: 999;
  }

  .loading_wrapper {
    position: absolute;

    top: 0;
    left: 0;
    right: 0;
    bottom: 0;

    width: 260px;
    height: 260px;

    margin: auto;
  }

  .loading_logo {
    width: 160px;
    height: 160px;
    margin: auto;
  }

  /**
    * Source images for this gif are from: https://www.reddit.com/r/Iota/comments/8kp9uv/created_an_iota_loading_animation_free_for_anyone/
    * They were scaled down and the transparency was replaced with the dim background of the page or kept transparent in case the browser supports webp.
  */
  .loading_logo {
    background-image: image-set(
      url('/iota.webp') type('image/webp'),
      url('/iota.gif') type('image/gif')
    );

    background-repeat: no-repeat;
  }

  .loading_text {
    text-shadow: 0px 2px 2px rgba(13, 15, 4, 0.42);
    font-size: 24px;
    color: white;
    text-align: center;
  }

  .welcome_screen {
    display: flex;
    color: white;
    margin: 25px auto;
    width: 25%;
  }

  .welcome_screen > .disabled {
    color: dimgrey !important;
  }

  .request_funds {
    width: 200px;
  }

  .place_bet {
    width: 200px;
  }

  .request_funds_text,
  .place_bet_text {
    text-shadow: 0px 2px 2px rgba(13, 15, 4, 0.42);
    font-size: 24px;
    text-align: center;
  }

  .request_funds_image {
    border: none;
    background-color: transparent;
    background-image: url(/money_w_k.png);
    width: 96px;
    height: 96px;
    margin: 15px auto;
  }

  .place_bet_image {
    border: none;
    background-color: transparent;
    background-image: url(/dice_w.png);
    width: 96px;
    height: 108px;
    margin: 15px auto;
  }

  .request_funds_button,
  .submit_bet_button {
    margin: 15px auto;
    display: block;
  }

  .bet_selection {
    margin: 15px auto;
    color: white;
    display: flex;
  }

  .bet_selection input {
    flex: 1;
  }
</style>
