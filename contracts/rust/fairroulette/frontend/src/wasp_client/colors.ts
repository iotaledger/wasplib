import { Buffer } from './buffer';

export class Colors {
  public static readonly IOTA_COLOR_STRING: string = '11111111111111111111111111111111';
  public static readonly IOTA_COLOR_BYTES: Buffer = Buffer.alloc(32);
}
export type ColorCollection = { [key: string]: bigint; };
