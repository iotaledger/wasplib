import { Base58 } from '../crypto'
import { Buffer } from '../buffer'
import type { IFaucetRequest } from './IFaucetRequest';

export class Faucet {
  public static ToStruct(buffer: Buffer): IFaucetRequest {
    return null;
  }

  public static ToBuffer(faucetRequest: IFaucetRequest): Buffer {
    const buffers = [];

    const payloadLen = Buffer.alloc(4);
    payloadLen.writeUInt32LE(109, 0);
    buffers.push(payloadLen);

    const faucetRequestType = Buffer.alloc(4);
    faucetRequestType.writeUInt32LE(2, 0);
    buffers.push(faucetRequestType);

    const addressBytes = Base58.decode(faucetRequest.address);
    buffers.push(addressBytes);

    const aManaPledgeBytes = Base58.decode(faucetRequest.accessManaPledgeID);
    buffers.push(aManaPledgeBytes);

    const cManaPledgeBytes = Base58.decode(faucetRequest.consensusManaPledgeID);
    buffers.push(cManaPledgeBytes);

    const data = Buffer.concat(buffers);

    return data;
  }
}
