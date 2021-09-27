import type { Buffer } from "../buffer";

export interface OffLedgerArgument {
  key: string,
  value: number;
}

export interface Balance {
  color: Buffer;
  balance: bigint;
}

export interface IOffLedger {
  requestType?: number,
  contract: number,
  entrypoint: number,
  arguments: OffLedgerArgument[],
  noonce: bigint,
  balances: Balance[],

  // Public Key and Signature will get set in the Sign function, so no inital set is required
  publicKey?: Buffer,
  signature?: Buffer,
}
