/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  PayableOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type { FunctionFragment, Result } from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../common";

export interface IUniversalRouterInterface extends utils.Interface {
  functions: {
    "execute(bytes,bytes[],uint256)": FunctionFragment;
  };

  getFunction(nameOrSignatureOrTopic: "execute"): FunctionFragment;

  encodeFunctionData(
    functionFragment: "execute",
    values: [
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>[],
      PromiseOrValue<BigNumberish>
    ]
  ): string;

  decodeFunctionResult(functionFragment: "execute", data: BytesLike): Result;

  events: {};
}

export interface IUniversalRouter extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: IUniversalRouterInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    execute(
      commands: PromiseOrValue<BytesLike>,
      inputs: PromiseOrValue<BytesLike>[],
      deadline: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  execute(
    commands: PromiseOrValue<BytesLike>,
    inputs: PromiseOrValue<BytesLike>[],
    deadline: PromiseOrValue<BigNumberish>,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    execute(
      commands: PromiseOrValue<BytesLike>,
      inputs: PromiseOrValue<BytesLike>[],
      deadline: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {};

  estimateGas: {
    execute(
      commands: PromiseOrValue<BytesLike>,
      inputs: PromiseOrValue<BytesLike>[],
      deadline: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    execute(
      commands: PromiseOrValue<BytesLike>,
      inputs: PromiseOrValue<BytesLike>[],
      deadline: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
