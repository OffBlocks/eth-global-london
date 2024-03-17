/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type { IPermit2, IPermit2Interface } from "../../contracts/IPermit2";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "token",
        type: "address",
      },
      {
        internalType: "address",
        name: "spender",
        type: "address",
      },
      {
        internalType: "uint160",
        name: "amount",
        type: "uint160",
      },
      {
        internalType: "uint48",
        name: "expiration",
        type: "uint48",
      },
    ],
    name: "approve",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

export class IPermit2__factory {
  static readonly abi = _abi;
  static createInterface(): IPermit2Interface {
    return new utils.Interface(_abi) as IPermit2Interface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IPermit2 {
    return new Contract(address, _abi, signerOrProvider) as IPermit2;
  }
}