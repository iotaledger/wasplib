import { Writable, writable } from 'svelte/store'
import type { Buffer, IKeyPair } from './wasp_client';

export const seed: Writable<Buffer> = writable();
export const keyPair: Writable<IKeyPair> = writable();
export const address: Writable<string> = writable();
export const addressIndex: Writable<number> = writable(0);
