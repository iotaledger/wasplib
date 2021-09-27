import { Base58, ED25519 } from '../crypto';
import { blake2b } from 'blakejs';
import { Buffer } from '../buffer';
import { SimpleBufferCursor } from '../simple_buffer_cursor';
import type { IKeyPair } from '../models';
import type { IOffLedger } from './IOffLedger';

export class OffLedger {
  public static ToStruct(buffer: Buffer): IOffLedger {
    const publicKeySize = 32;
    const colorLength = 32;
    const signatureSize = 64;

    let reader = new SimpleBufferCursor(buffer);

    const requestType = reader.readIntBE(1);
    const contract = reader.readUInt32LE();
    const entrypoint = reader.readUInt32LE();
    const numArguments = reader.readUInt32LE();

    const args = [];

    for (let i = 0; i < numArguments; i++) {
      let sz16 = reader.readUInt16LE();
      let key = reader.readBytes(sz16);
      let sz32 = reader.readUInt32LE();
      let value = reader.readBytes(sz32);

      args.push({ key: key, value: value });
    }

    const publicKey = reader.readBytes(publicKeySize);
    const noonce = reader.readUInt64LE();
    const numBalances = reader.readUInt32LE();

    const balances = [];

    for (let i = 0; i < numBalances; i++) {
      const colorBytes = reader.readBytes(colorLength);
      const balance = reader.readUInt64LE();

      balances.push({ color: colorBytes, balance: balance });
    }

    const signature = reader.readBytes(signatureSize);

    let offLedgerStruct: IOffLedger = {
      requestType: requestType,
      contract: contract,
      entrypoint: entrypoint,
      arguments: args,
      publicKey: Buffer.from(publicKey),
      noonce: noonce,
      balances: balances,
      signature: Buffer.from(signature),
    };

    return offLedgerStruct;
  }

  public static ToBuffer(req: IOffLedger): Buffer {
    const buffer = new SimpleBufferCursor(Buffer.alloc(0));

    if ([0, 1].includes(req.requestType)) {
      buffer.writeIntBE(req.requestType, 1);
    }

    buffer.writeUInt32LE(req.contract);
    buffer.writeUInt32LE(req.entrypoint);


    buffer.writeUInt32LE(req.arguments.length || 0);

    if (req.arguments) {
      for (let arg of req.arguments) {
        //  debugger;

        const keyBuffer = Buffer.from(arg.key);

        buffer.writeUInt16LE(keyBuffer.length);
        buffer.writeBytes(keyBuffer);

        const valueBuffer = Buffer.alloc(8);
        valueBuffer.writeInt32LE(arg.value, 0);

        buffer.writeUInt32LE(valueBuffer.length);
        buffer.writeBytes(valueBuffer);
      }
    }

    buffer.writeBytes(req.publicKey);
    buffer.writeUInt64LE(req.noonce);

    buffer.writeUInt32LE(req.balances.length || 0);

    if (req.balances) {
      for (let balance of req.balances) {
        buffer.writeBytes(balance.color);
        buffer.writeUInt64LE(balance.balance);
      }
    }

    if (req.signature && req.signature.length > 0) {
      buffer.writeBytes(req.signature);
    }

    return buffer.buffer;
  }

  public static Sign(request: IOffLedger, keyPair: IKeyPair): IOffLedger {

    // Create a copy without requestType and signature
    // adding the requestType and|or an empty signature would result in an invalid signature in the next step.
    const cleanCopyOfRequest: IOffLedger = {
      arguments: request.arguments,
      balances: request.balances,
      contract: request.contract,
      entrypoint: request.entrypoint,
      noonce: request.noonce,
      publicKey: keyPair.publicKey,

      requestType: null,
      signature: null
    };

    const requestBuffer = this.ToBuffer(cleanCopyOfRequest);

    request.publicKey = keyPair.publicKey;
    request.signature = ED25519.privateSign(keyPair, requestBuffer);

    return request;
  }

  public static GetRequestId(request: IOffLedger): string {
    const bufferedRequest = OffLedger.ToBuffer(request);
    const hash = blake2b(bufferedRequest, undefined, 32);
    const extendedHash = Buffer.concat([hash, Buffer.alloc(2)]);
    const id = Base58.encode(extendedHash);

    return id;
  }
}
