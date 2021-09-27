import { IFaucetRequest, IOffLedger, OffLedger } from './binary_models';
import type { Buffer } from './buffer';

import type {
  IAllowedManaPledgeResponse,
  IOffLedgerRequest,
  IFaucetResponse,
  IResponse,
  IUnspentOutputsRequest,
  IUnspentOutputsResponse,
  ISendTransactionRequest,
  ISendTransactionResponse,
  IKeyPair,
  ISingleUnspentOutputResponse,
} from './models';


export interface IExtendedResponse<U> {
  body: U;
  response: Response;
}

export interface BasicClientConfiguration {
  WaspAPIUrl: string;
  GoShimmerAPIUrl: string;
  SeedUnsafe: Buffer;
}

export interface CallViewResponse extends IResponse {
  Items: [{ Key: string, Value: string; }];
}

export class BasicClient {

  private configuration: BasicClientConfiguration;

  constructor(configuration: BasicClientConfiguration) {
    this.configuration = configuration;
  }

  public async getAllowedManaPledge(): Promise<IAllowedManaPledgeResponse> {
    return this.sendRequest<null, IAllowedManaPledgeResponse>(this.configuration.GoShimmerAPIUrl,
      'get', 'mana/allowedManaPledge');
  }

  public async postRequest(offLedger: IOffLedger, address: string, keyPair: IKeyPair) {
    return Promise.resolve();
  }

  public async sendFaucetRequest(faucetRequest: IFaucetRequest): Promise<IFaucetResponse> {
    const response = await this.sendRequest<IFaucetRequest, IFaucetResponse>(this.configuration.GoShimmerAPIUrl, 'post', 'faucet', faucetRequest);

    return response;
  }

  public async sendOffLedgerRequest(chainId: string, offLedgerRequest: IOffLedger): Promise<void> {
    const request = { Request: OffLedger.ToBuffer(offLedgerRequest).toString('base64') };

    await this.sendRequestExt<IOffLedgerRequest, null>(this.configuration.WaspAPIUrl, 'post', `request/${chainId}`, request);
  }

  public async sendExecutionRequest(chainId: string, offLedgerRequestId: string): Promise<void> {
    await this.sendRequestExt<IOffLedgerRequest, null>(this.configuration.WaspAPIUrl, 'get', `chain/${chainId}/request/${offLedgerRequestId}/wait`);
  }

  public async callView(chainId: string, contractHName: string, entryPoint: string): Promise<CallViewResponse> {
    const url = `chain/${chainId}/contract/${contractHName}/callview/${entryPoint}`;

    const result = await this.sendRequestExt<any, CallViewResponse>(this.configuration.WaspAPIUrl, 'get', url);

    return result.body;
  }

  public async getAddressUnspentOutputs(address: string) {
    return this.sendRequest<IUnspentOutputsRequest, ISingleUnspentOutputResponse>(this.configuration.GoShimmerAPIUrl,
      'get', `ledgerstate/addresses/${address}/unspentOutputs`);
  }

  public async unspentOutputs(request: IUnspentOutputsRequest): Promise<IUnspentOutputsResponse> {
    return this.sendRequest<IUnspentOutputsRequest, IUnspentOutputsResponse>(this.configuration.GoShimmerAPIUrl,
      'post', 'ledgerstate/addresses/unspentOutputs', request);
  }

  public async sendTransaction(request: ISendTransactionRequest): Promise<ISendTransactionResponse> {
    return this.sendRequest<ISendTransactionRequest, ISendTransactionResponse>(this.configuration.GoShimmerAPIUrl,
      'post', 'ledgerstate/transactions', request);
  }

  private async sendRequest<T, U extends IResponse>(
    url: string,
    verb: 'put' | 'post' | 'get' | 'delete',
    path: string,
    request?: T | undefined): Promise<U> {

    const response = await this.sendRequestExt<T, U>(url, verb, path, request);

    return response.body;
  }

  private async sendRequestExt<T, U extends IResponse>(
    url: string,
    verb: 'put' | 'post' | 'get' | 'delete',
    path: string,
    request?: T | undefined): Promise<IExtendedResponse<U>> {

    let response: U;
    let fetchResponse: Response;

    try {
      const headers: { [id: string]: string; } = {
        'Content-Type': 'application/json'
      };

      if (verb == 'get' || verb == 'delete') {
        fetchResponse = await fetch(
          `${url}/${path}`,
          {
            method: verb,
            headers,
          }
        );
      } else if (verb == 'post' || verb == 'put') {
        fetchResponse = await fetch(
          `${url}/${path}`,
          {
            method: verb,
            headers,
            body: JSON.stringify(request)
          }
        );
      }

      if (!fetchResponse) {
        throw new Error('No data was returned from the API');
      }

      try {
        response = await fetchResponse.json();
      } catch (err) {

        if (!fetchResponse.ok) {
          const text = await fetchResponse.text();
          throw new Error(err.message + '   ---   ' + text);
        }

      }

    } catch (err) {
      throw new Error(`The application is not able to complete the request, due to the following error:\n\n${err.message}`);
    }

    return { body: response, response: fetchResponse };
  }
}
