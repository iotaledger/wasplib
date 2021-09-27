import { blake2b } from 'blakejs';
import { Buffer } from './buffer';

export default class ProofOfWork {
  public static numberToUInt64LE(n: bigint): Buffer {
    const buffer = Buffer.alloc(8);
    buffer.writeBigUInt64LE(n, 0);

    return buffer;
  }

  public static calculateProofOfWork(target: number, message: Buffer): number {
    for (let nonce = 0; ; nonce++) {
      const nonceLE = this.numberToUInt64LE(BigInt(nonce));
      const data = Buffer.concat([message, nonceLE]);

      const digest = blake2b(data);
      const b = Buffer.alloc(4);

      for (let i = 0; i < 4; i++) {
        b[i] = digest[i];
      }

      const leadingZeros = Math.clz32(b.readUInt32BE(0));

      if (leadingZeros >= target) {
        console.log("PoW Single Thread done");
        return nonce;
      }
    }
  }
}

