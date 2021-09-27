import type { Buffer } from "../buffer";

export interface IKeyPair {
  publicKey: Buffer;
  secretKey: Buffer;
}
