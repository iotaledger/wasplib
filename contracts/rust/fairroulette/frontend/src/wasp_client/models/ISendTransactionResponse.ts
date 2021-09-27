import type { IResponse } from "./IResponse";

export interface ISendTransactionResponse extends IResponse {
    transaction_id?: string;
}
