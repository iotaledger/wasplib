import type { Buffer } from '../buffer';
export interface IUnlockBlock {
    type: number;
    referenceIndex: number;
    publicKey: Buffer;
    signature: Buffer;
}
