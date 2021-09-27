export interface IWalletOutputInclusionState {
    solid?: boolean;
    confirmed?: boolean;
    rejected?: boolean;
    liked?: boolean;
    conflicting?: boolean;
    finalized?: boolean;
    preferred?: boolean;
}
