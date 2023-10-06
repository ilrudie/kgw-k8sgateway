/* eslint-disable */
// package: azure.options.gloo.solo.io
// file: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/azure/azure.proto

import * as jspb from "google-protobuf";
import * as extproto_ext_pb from "../../../../../../../../../extproto/ext_pb";
import * as github_com_solo_io_solo_kit_api_v1_ref_pb from "../../../../../../../../../github.com/solo-io/solo-kit/api/v1/ref_pb";

export class UpstreamSpec extends jspb.Message {
  getFunctionAppName(): string;
  setFunctionAppName(value: string): void;

  hasSecretRef(): boolean;
  clearSecretRef(): void;
  getSecretRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setSecretRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  clearFunctionsList(): void;
  getFunctionsList(): Array<UpstreamSpec.FunctionSpec>;
  setFunctionsList(value: Array<UpstreamSpec.FunctionSpec>): void;
  addFunctions(value?: UpstreamSpec.FunctionSpec, index?: number): UpstreamSpec.FunctionSpec;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpstreamSpec.AsObject;
  static toObject(includeInstance: boolean, msg: UpstreamSpec): UpstreamSpec.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpstreamSpec, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpstreamSpec;
  static deserializeBinaryFromReader(message: UpstreamSpec, reader: jspb.BinaryReader): UpstreamSpec;
}

export namespace UpstreamSpec {
  export type AsObject = {
    functionAppName: string,
    secretRef?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    functionsList: Array<UpstreamSpec.FunctionSpec.AsObject>,
  }

  export class FunctionSpec extends jspb.Message {
    getFunctionName(): string;
    setFunctionName(value: string): void;

    getAuthLevel(): UpstreamSpec.FunctionSpec.AuthLevelMap[keyof UpstreamSpec.FunctionSpec.AuthLevelMap];
    setAuthLevel(value: UpstreamSpec.FunctionSpec.AuthLevelMap[keyof UpstreamSpec.FunctionSpec.AuthLevelMap]): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FunctionSpec.AsObject;
    static toObject(includeInstance: boolean, msg: FunctionSpec): FunctionSpec.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FunctionSpec, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FunctionSpec;
    static deserializeBinaryFromReader(message: FunctionSpec, reader: jspb.BinaryReader): FunctionSpec;
  }

  export namespace FunctionSpec {
    export type AsObject = {
      functionName: string,
      authLevel: UpstreamSpec.FunctionSpec.AuthLevelMap[keyof UpstreamSpec.FunctionSpec.AuthLevelMap],
    }

    export interface AuthLevelMap {
      ANONYMOUS: 0;
      FUNCTION: 1;
      ADMIN: 2;
    }

    export const AuthLevel: AuthLevelMap;
  }
}

export class DestinationSpec extends jspb.Message {
  getFunctionName(): string;
  setFunctionName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestinationSpec.AsObject;
  static toObject(includeInstance: boolean, msg: DestinationSpec): DestinationSpec.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DestinationSpec, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestinationSpec;
  static deserializeBinaryFromReader(message: DestinationSpec, reader: jspb.BinaryReader): DestinationSpec;
}

export namespace DestinationSpec {
  export type AsObject = {
    functionName: string,
  }
}