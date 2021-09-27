import { blake2b } from 'blakejs'
import { Buffer } from '../buffer'

export class HName {
  public static HashAsString(textToHash: string): string {
    const hashNumber = this.HashAsNumber(textToHash);
    const hashString = hashNumber.toString(16);

    return hashString;
  }

  public static HashAsBuffer(textToHash: string): Buffer {
    const buffer = Buffer.from(textToHash);
    const result = blake2b(buffer, undefined, 32);

    return Buffer.from(result);
  }

  public static HashAsNumber(textToHash: string): number {
    const bufferedHash = this.HashAsBuffer(textToHash);
    const resultHash = Buffer.from(bufferedHash).readUInt32LE(0);

    return resultHash;
  }
}
