import { Base58 } from './crypto';
import { Buffer } from './buffer';
import { ColorCollection, Colors } from './colors';
import { Transaction } from './transaction';
import type { BasicClient } from './basic_client';
import type { IKeyPair, ITransaction, IUnlockBlock, IWalletAddressOutput, IWalletOutput } from './models';

export type BuiltOutputResult = {
  [address: string]: {
    /**
     * The color.
     */
    color: string;
    /**
     * The value.
     */
    value: bigint;
  }[];
};

export type ConsumedOutputs = {
  [address: string]: { [outputID: string]: IWalletOutput; };
};


export class BasicWallet {

  private client: BasicClient;
  constructor(client: BasicClient) {
    this.client = client;
  }

  private fakeBigBalance(balances: ColorCollection) {
    const colorCollection: ColorCollection = {};

    for (let color in balances) {
      colorCollection[color] = BigInt(balances[color]);
    }

    return colorCollection;
  }

  public async getUnspentOutputs(address: string) {
    const unspents = await this.client.unspentOutputs({ addresses: [address] });

    const usedAddresses = unspents.unspentOutputs.filter(u => u.outputs.length > 0);

    const unspentOutputs = usedAddresses.map(uo => ({
      address: uo.address.base58,
      outputs: uo.outputs.map(uid => ({
        id: uid.output.outputID.base58,
        balances: this.fakeBigBalance(uid.output.output.balances),
        inclusionState: uid.inclusionState
      }))
    }));

    return unspentOutputs;
  }

  public determineOutputsToConsume(unspentOutputs: IWalletAddressOutput[], iotas: bigint): ConsumedOutputs {
    const outputsToConsume: { [address: string]: { [outputID: string]: IWalletOutput; }; } = {};

    let iotasLeft = iotas;

    for (const unspentOutput of unspentOutputs) {
      let outputsFromAddressSpent = false;

      for (const output of unspentOutput.outputs) {
        let requiredColorFoundInOutput = false;

        if (!output.balances[Colors.IOTA_COLOR_STRING]) {
          continue;
        }

        const balance = output.balances[Colors.IOTA_COLOR_STRING];

        if (iotasLeft > 0n) {
          if (iotasLeft > balance) {
            iotasLeft -= balance;
          } else {
            iotasLeft = 0n;
          }

          requiredColorFoundInOutput = true;
        }

        // if we found required tokens in this output
        if (requiredColorFoundInOutput) {
          // store the output in the outputs to use for the transfer
          outputsToConsume[unspentOutput.address] = {};
          outputsToConsume[unspentOutput.address][output.id] = output;

          // mark address as spent
          outputsFromAddressSpent = true;
        }

      }

      if (outputsFromAddressSpent) {
        for (const output of unspentOutput.outputs) {
          outputsToConsume[unspentOutput.address][output.id] = output;
        }
      }
    }

    return outputsToConsume;
  }

  public buildOutputs(remainderAddress: string, destinationAddress: string, iotas: bigint, consumedFunds: ColorCollection): BuiltOutputResult {
    const outputsByColor: { [address: string]: ColorCollection; } = {};

    // build outputs for destinations

    if (!outputsByColor[destinationAddress]) {
      outputsByColor[destinationAddress] = {};
    }


    if (!outputsByColor[destinationAddress][Colors.IOTA_COLOR_STRING]) {
      outputsByColor[destinationAddress][Colors.IOTA_COLOR_STRING] = 0n;
    }
    const t = outputsByColor[destinationAddress][Colors.IOTA_COLOR_STRING];
    outputsByColor[destinationAddress][Colors.IOTA_COLOR_STRING] += iotas;

    consumedFunds[Colors.IOTA_COLOR_STRING] -= iotas;
    if (consumedFunds[Colors.IOTA_COLOR_STRING] === 0n) {
      delete consumedFunds[Colors.IOTA_COLOR_STRING];
    }



    // build outputs for remainder
    if (Object.keys(consumedFunds).length > 0) {
      if (!remainderAddress) {
        throw new Error("No remainder address available");
      }
      if (!outputsByColor[remainderAddress]) {
        outputsByColor[remainderAddress] = {};
      }
      for (const consumed in consumedFunds) {
        if (!outputsByColor[remainderAddress][consumed]) {
          outputsByColor[remainderAddress][consumed] = 0n;
        }
        outputsByColor[remainderAddress][consumed] += consumedFunds[consumed];
      }
    }

    // construct result
    const outputsBySlice: BuiltOutputResult = {};

    for (const address in outputsByColor) {
      outputsBySlice[address] = [];
      for (const color in outputsByColor[address]) {
        outputsBySlice[address].push({
          color,
          value: outputsByColor[address][color]
        });
      }
    }

    return outputsBySlice;
  }

  public buildInputs(outputsToUseAsInputs: { [address: string]: { [outputID: string]: IWalletOutput; }; }): {
    /**
     * The inputs to send.
     */
    inputs: string[];
    /**
     * The fund that were consumed.
     */
    consumedFunds: ColorCollection;
  } {
    const inputs: string[] = [];
    const consumedFunds: ColorCollection = {};

    for (const address in outputsToUseAsInputs) {
      for (const outputID in outputsToUseAsInputs[address]) {
        inputs.push(outputID);

        for (const color in outputsToUseAsInputs[address][outputID].balances) {
          const balance = outputsToUseAsInputs[address][outputID].balances[color];

          if (!consumedFunds[color]) {
            consumedFunds[color] = balance;
          } else {
            consumedFunds[color] += balance;
          }
        }
      }
    }

    inputs.sort((a, b) => Base58.decode(a).compare(Base58.decode(b)));

    return { inputs, consumedFunds };
  }

  public unlockBlocks(tx: ITransaction, keyPair: IKeyPair, address: string, consumedOutputs: ConsumedOutputs, builtInputs: string[]) {
    const unlockBlocks: IUnlockBlock[] = [];
    const txEssence = Transaction.essence(tx, Buffer.alloc(0));

    const addressByOutputID: { [outputID: string]: string; } = {};
    for (const address in consumedOutputs) {
      for (const outputID in consumedOutputs[address]) {
        addressByOutputID[outputID] = address;
      }
    }

    const existingUnlockBlocks: { [address: string]: number; } = {};
    for (const index in builtInputs) {
      const addr = address == addressByOutputID[builtInputs[index]];
      if (addr) {
        if (existingUnlockBlocks[address] !== undefined) {
          unlockBlocks.push({ type: 1, referenceIndex: existingUnlockBlocks[address], publicKey: Buffer.alloc(0), signature: Buffer.alloc(0) });
          continue;
        }

        const signatureUnlockBlock = { type: 0, referenceIndex: 0, publicKey: keyPair.publicKey, signature: Transaction.sign(keyPair, txEssence) };
        existingUnlockBlocks[address] = unlockBlocks.length;
        unlockBlocks.push(signatureUnlockBlock);
      }
    }

    return unlockBlocks;
  }
}
