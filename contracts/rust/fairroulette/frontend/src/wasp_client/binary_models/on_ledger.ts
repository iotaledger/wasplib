import { Base58, ED25519 } from '../crypto';
import { blake2b } from 'blakejs';
import { Buffer } from '../buffer';
import { SimpleBufferCursor } from '../simple_buffer_cursor';
import type { IKeyPair } from '../models';
import type { IOnLedger } from './IOnLedger';

export class OnLedger {
  public static ToStruct(buffer: Buffer): IOnLedger {
    const publicKeySize = 32;
    const colorLength = 32;
    const signatureSize = 64;

    let reader = new SimpleBufferCursor(buffer);

    const contract = reader.readUInt32LE();
    const entrypoint = reader.readUInt32LE();
    const noonce = reader.readBytes(1);

    const numArguments = reader.readUInt32LE();

    const args = [];

    for (let i = 0; i < numArguments; i++) {
      let sz16 = reader.readUInt16LE();
      let key = reader.readBytes(sz16);
      let sz32 = reader.readUInt32LE();
      let value = reader.readBytes(sz32);

      args.push({ key: key, value: value });
    }

    let offLedgerStruct: IOnLedger = {
      contract: contract,
      entrypoint: entrypoint,
      arguments: args,
      noonce: 0,
    };

    return offLedgerStruct;
  }

  public static ToBuffer(req: IOnLedger): Buffer {
    const buffer = new SimpleBufferCursor(Buffer.alloc(0));

    buffer.writeUInt32LE(0);
    buffer.writeUInt32LE(req.contract);
    buffer.writeUInt32LE(req.entrypoint);
    buffer.writeBytes(Buffer.alloc(1, 0));

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


    return buffer.buffer;
  }

  public static Sign(request: IOnLedger, keyPair: IKeyPair): IOnLedger {

    // Create a copy without requestType and signature
    // adding the requestType and|or an empty signature would result in an invalid signature in the next step.
    const cleanCopyOfRequest: IOnLedger = {
      arguments: request.arguments,
      contract: request.contract,
      entrypoint: request.entrypoint,
      noonce: request.noonce,
    };

    const requestBuffer = this.ToBuffer(cleanCopyOfRequest);

    return request;
  }

  public static GetRequestId(request: IOnLedger): string {
    const bufferedRequest = OnLedger.ToBuffer(request);
    const hash = blake2b(bufferedRequest, undefined, 32);
    const extendedHash = Buffer.concat([hash, Buffer.alloc(2)]);
    const id = Base58.encode(extendedHash);

    return id;
  }
}
