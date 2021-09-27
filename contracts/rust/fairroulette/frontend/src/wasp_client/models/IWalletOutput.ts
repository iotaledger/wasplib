import type { ColorCollection } from '../colors';
import type { IWalletOutputBalance } from "./IWalletOutputBalance";
import type { IWalletOutputInclusionState } from "./IWalletOutputInclusionState";

export interface IWalletOutput {
    /**
     * The id.
     */
    id: string;

    /**
     * The balances.
     */
    balances: ColorCollection;

    /**
     * Inclusion state.
     */
    inclusionState: IWalletOutputInclusionState;
}
