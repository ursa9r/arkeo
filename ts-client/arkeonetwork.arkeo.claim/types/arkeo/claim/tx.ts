/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Chain, chainFromJSON, chainToJSON } from "./claim_record";

export const protobufPackage = "arkeonetwork.arkeo.claim";

export interface MsgClaimEth {
  creator: string;
  /** the adress the claim is for */
  ethAddress: string;
  /** EIP712 signature that has to be signed by ethAddress */
  signature: string;
}

export interface MsgClaimEthResponse {
}

export interface MsgClaimArkeo {
  creator: string;
}

export interface MsgClaimArkeoResponse {
}

export interface MsgTransferClaim {
  creator: string;
  toAddress: string;
}

export interface MsgTransferClaimResponse {
}

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgAddClaim {
  creator: string;
  chain: Chain;
  address: string;
  amount: number;
}

export interface MsgAddClaimResponse {
}

function createBaseMsgClaimEth(): MsgClaimEth {
  return { creator: "", ethAddress: "", signature: "" };
}

export const MsgClaimEth = {
  encode(message: MsgClaimEth, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.ethAddress !== "") {
      writer.uint32(18).string(message.ethAddress);
    }
    if (message.signature !== "") {
      writer.uint32(26).string(message.signature);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimEth {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimEth();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.ethAddress = reader.string();
          break;
        case 3:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimEth {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      ethAddress: isSet(object.ethAddress) ? String(object.ethAddress) : "",
      signature: isSet(object.signature) ? String(object.signature) : "",
    };
  },

  toJSON(message: MsgClaimEth): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.ethAddress !== undefined && (obj.ethAddress = message.ethAddress);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimEth>, I>>(object: I): MsgClaimEth {
    const message = createBaseMsgClaimEth();
    message.creator = object.creator ?? "";
    message.ethAddress = object.ethAddress ?? "";
    message.signature = object.signature ?? "";
    return message;
  },
};

function createBaseMsgClaimEthResponse(): MsgClaimEthResponse {
  return {};
}

export const MsgClaimEthResponse = {
  encode(_: MsgClaimEthResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimEthResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimEthResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgClaimEthResponse {
    return {};
  },

  toJSON(_: MsgClaimEthResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimEthResponse>, I>>(_: I): MsgClaimEthResponse {
    const message = createBaseMsgClaimEthResponse();
    return message;
  },
};

function createBaseMsgClaimArkeo(): MsgClaimArkeo {
  return { creator: "" };
}

export const MsgClaimArkeo = {
  encode(message: MsgClaimArkeo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimArkeo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimArkeo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimArkeo {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgClaimArkeo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimArkeo>, I>>(object: I): MsgClaimArkeo {
    const message = createBaseMsgClaimArkeo();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgClaimArkeoResponse(): MsgClaimArkeoResponse {
  return {};
}

export const MsgClaimArkeoResponse = {
  encode(_: MsgClaimArkeoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimArkeoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimArkeoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgClaimArkeoResponse {
    return {};
  },

  toJSON(_: MsgClaimArkeoResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimArkeoResponse>, I>>(_: I): MsgClaimArkeoResponse {
    const message = createBaseMsgClaimArkeoResponse();
    return message;
  },
};

function createBaseMsgTransferClaim(): MsgTransferClaim {
  return { creator: "", toAddress: "" };
}

export const MsgTransferClaim = {
  encode(message: MsgTransferClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferClaim();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.toAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferClaim {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      toAddress: isSet(object.toAddress) ? String(object.toAddress) : "",
    };
  },

  toJSON(message: MsgTransferClaim): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferClaim>, I>>(object: I): MsgTransferClaim {
    const message = createBaseMsgTransferClaim();
    message.creator = object.creator ?? "";
    message.toAddress = object.toAddress ?? "";
    return message;
  },
};

function createBaseMsgTransferClaimResponse(): MsgTransferClaimResponse {
  return {};
}

export const MsgTransferClaimResponse = {
  encode(_: MsgTransferClaimResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferClaimResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferClaimResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgTransferClaimResponse {
    return {};
  },

  toJSON(_: MsgTransferClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferClaimResponse>, I>>(_: I): MsgTransferClaimResponse {
    const message = createBaseMsgTransferClaimResponse();
    return message;
  },
};

function createBaseMsgAddClaim(): MsgAddClaim {
  return { creator: "", chain: 0, address: "", amount: 0 };
}

export const MsgAddClaim = {
  encode(message: MsgAddClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chain !== 0) {
      writer.uint32(16).int32(message.chain);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.amount !== 0) {
      writer.uint32(32).int32(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddClaim();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chain = reader.int32() as any;
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.amount = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddClaim {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      chain: isSet(object.chain) ? chainFromJSON(object.chain) : 0,
      address: isSet(object.address) ? String(object.address) : "",
      amount: isSet(object.amount) ? Number(object.amount) : 0,
    };
  },

  toJSON(message: MsgAddClaim): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chain !== undefined && (obj.chain = chainToJSON(message.chain));
    message.address !== undefined && (obj.address = message.address);
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddClaim>, I>>(object: I): MsgAddClaim {
    const message = createBaseMsgAddClaim();
    message.creator = object.creator ?? "";
    message.chain = object.chain ?? 0;
    message.address = object.address ?? "";
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseMsgAddClaimResponse(): MsgAddClaimResponse {
  return {};
}

export const MsgAddClaimResponse = {
  encode(_: MsgAddClaimResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddClaimResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddClaimResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddClaimResponse {
    return {};
  },

  toJSON(_: MsgAddClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddClaimResponse>, I>>(_: I): MsgAddClaimResponse {
    const message = createBaseMsgAddClaimResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  ClaimEth(request: MsgClaimEth): Promise<MsgClaimEthResponse>;
  ClaimArkeo(request: MsgClaimArkeo): Promise<MsgClaimArkeoResponse>;
  TransferClaim(request: MsgTransferClaim): Promise<MsgTransferClaimResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddClaim(request: MsgAddClaim): Promise<MsgAddClaimResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ClaimEth = this.ClaimEth.bind(this);
    this.ClaimArkeo = this.ClaimArkeo.bind(this);
    this.TransferClaim = this.TransferClaim.bind(this);
    this.AddClaim = this.AddClaim.bind(this);
  }
  ClaimEth(request: MsgClaimEth): Promise<MsgClaimEthResponse> {
    const data = MsgClaimEth.encode(request).finish();
    const promise = this.rpc.request("arkeonetwork.arkeo.claim.Msg", "ClaimEth", data);
    return promise.then((data) => MsgClaimEthResponse.decode(new _m0.Reader(data)));
  }

  ClaimArkeo(request: MsgClaimArkeo): Promise<MsgClaimArkeoResponse> {
    const data = MsgClaimArkeo.encode(request).finish();
    const promise = this.rpc.request("arkeonetwork.arkeo.claim.Msg", "ClaimArkeo", data);
    return promise.then((data) => MsgClaimArkeoResponse.decode(new _m0.Reader(data)));
  }

  TransferClaim(request: MsgTransferClaim): Promise<MsgTransferClaimResponse> {
    const data = MsgTransferClaim.encode(request).finish();
    const promise = this.rpc.request("arkeonetwork.arkeo.claim.Msg", "TransferClaim", data);
    return promise.then((data) => MsgTransferClaimResponse.decode(new _m0.Reader(data)));
  }

  AddClaim(request: MsgAddClaim): Promise<MsgAddClaimResponse> {
    const data = MsgAddClaim.encode(request).finish();
    const promise = this.rpc.request("arkeonetwork.arkeo.claim.Msg", "AddClaim", data);
    return promise.then((data) => MsgAddClaimResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
