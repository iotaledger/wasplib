import { Buffer } from './buffer';

export class SimpleBufferCursor {

  private _buffer: Buffer;
  private _traverse: number;

  get buffer() {
    return this._buffer;
  }

  constructor(buffer: Buffer = Buffer.alloc(0)) {
    this._buffer = buffer;
    this._traverse = 0;
  }


  readIntBE(length: number): number {
    const value = this._buffer.readIntBE(this._traverse, length);
    this._traverse += length;

    return value;
  }

  readUInt32LE(): number {
    const value = this._buffer.readUInt32LE(this._traverse);
    this._traverse += 4;

    return value;
  }

  readUInt64LE(): bigint {
    const value = this._buffer.readBigUInt64LE(this._traverse);
    this._traverse += 8;

    return value;
  }

  readUInt16LE(): number {
    const value = this._buffer.readUInt16LE(this._traverse);
    this._traverse += 2;

    return value;
  }

  readBytes(length: number): Uint8Array {
    let subBuffer = this._buffer.subarray(this._traverse, this._traverse + length);
    this._traverse += length;

    return subBuffer;
  }

  writeIntBE(value: number, length: number) {
    const nBuffer = Buffer.alloc(length);
    nBuffer.writeIntBE(value, 0, length);

    this._buffer = Buffer.concat([this._buffer, nBuffer]);
  }

  writeInt8(value: number) {
    const nBuffer = Buffer.alloc(1);
    nBuffer.writeInt8(value, 0);

    this._buffer = Buffer.concat([this._buffer, nBuffer]);
  }

  writeUInt32LE(value: number) {
    const nBuffer = Buffer.alloc(4);
    nBuffer.writeUInt32LE(value, 0);

    this._buffer = Buffer.concat([this._buffer, nBuffer]);
  }

  writeUInt64LE(value: bigint) {
    const nBuffer = Buffer.alloc(8);
    nBuffer.writeBigUInt64LE(value, 0);

    this._buffer = Buffer.concat([this._buffer, nBuffer]);
  }

  writeUInt16LE(value: number) {
    const nBuffer = Buffer.alloc(2);
    nBuffer.writeUInt16LE(value, 0);

    this._buffer = Buffer.concat([this._buffer, nBuffer]);
  }

  writeBytes(bytes: Buffer) {
    this._buffer = Buffer.concat([this._buffer, bytes]);
  }
}
