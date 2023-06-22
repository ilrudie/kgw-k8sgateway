/* eslint-disable */
// package: core.solo.io
// file: github.com/solo-io/solo-kit/api/v1/status.proto

import * as jspb from "google-protobuf";
import * as extproto_ext_pb from "../../../../../extproto/ext_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

export class NamespacedStatuses extends jspb.Message {
  getStatusesMap(): jspb.Map<string, Status>;
  clearStatusesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NamespacedStatuses.AsObject;
  static toObject(includeInstance: boolean, msg: NamespacedStatuses): NamespacedStatuses.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NamespacedStatuses, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NamespacedStatuses;
  static deserializeBinaryFromReader(message: NamespacedStatuses, reader: jspb.BinaryReader): NamespacedStatuses;
}

export namespace NamespacedStatuses {
  export type AsObject = {
    statusesMap: Array<[string, Status.AsObject]>,
  }
}

export class Status extends jspb.Message {
  getState(): Status.StateMap[keyof Status.StateMap];
  setState(value: Status.StateMap[keyof Status.StateMap]): void;

  getReason(): string;
  setReason(value: string): void;

  getReportedBy(): string;
  setReportedBy(value: string): void;

  getSubresourceStatusesMap(): jspb.Map<string, Status>;
  clearSubresourceStatusesMap(): void;
  hasDetails(): boolean;
  clearDetails(): void;
  getDetails(): google_protobuf_struct_pb.Struct | undefined;
  setDetails(value?: google_protobuf_struct_pb.Struct): void;

  clearMessagesList(): void;
  getMessagesList(): Array<string>;
  setMessagesList(value: Array<string>): void;
  addMessages(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Status.AsObject;
  static toObject(includeInstance: boolean, msg: Status): Status.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Status, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Status;
  static deserializeBinaryFromReader(message: Status, reader: jspb.BinaryReader): Status;
}

export namespace Status {
  export type AsObject = {
    state: Status.StateMap[keyof Status.StateMap],
    reason: string,
    reportedBy: string,
    subresourceStatusesMap: Array<[string, Status.AsObject]>,
    details?: google_protobuf_struct_pb.Struct.AsObject,
    messagesList: Array<string>,
  }

  export interface StateMap {
    PENDING: 0;
    ACCEPTED: 1;
    REJECTED: 2;
    WARNING: 3;
  }

  export const State: StateMap;
}