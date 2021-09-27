import type { IUnlockBlock } from './IUnlockBlock';
import type { Buffer } from '../buffer';
import type { BuiltOutputResult } from '../basic_wallet';
export interface ITransaction {
    /**
    * The transaction's version.
    */
    version: number;

    /**
     * The transaction's timestamp.
     */
    timestamp: bigint;

    /**
     * The nodeID to pledge access mana.
     */
    aManaPledge: string;

    /**
    * The nodeID to pledge consensus mana.
    */
    cManaPledge: string;

    /**
     * The inputs to send.
     */
    inputs: string[];

    payload: Buffer;

    chainId: string;
    /**
     * The outputs to send.
     */
    outputs: BuiltOutputResult;

    /**
     * The signatures to send.
     */
    unlockBlocks: IUnlockBlock[];
}
