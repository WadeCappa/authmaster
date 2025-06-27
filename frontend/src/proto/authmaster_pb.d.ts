import * as jspb from 'google-protobuf'



export class TestAuthRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestAuthRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TestAuthRequest): TestAuthRequest.AsObject;
  static serializeBinaryToWriter(message: TestAuthRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestAuthRequest;
  static deserializeBinaryFromReader(message: TestAuthRequest, reader: jspb.BinaryReader): TestAuthRequest;
}

export namespace TestAuthRequest {
  export type AsObject = {
  }
}

export class TestAuthResponse extends jspb.Message {
  getUserid(): number;
  setUserid(value: number): TestAuthResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestAuthResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TestAuthResponse): TestAuthResponse.AsObject;
  static serializeBinaryToWriter(message: TestAuthResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestAuthResponse;
  static deserializeBinaryFromReader(message: TestAuthResponse, reader: jspb.BinaryReader): TestAuthResponse;
}

export namespace TestAuthResponse {
  export type AsObject = {
    userid: number,
  }
}

export class LoginRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): LoginRequest;

  getPassword(): string;
  setPassword(value: string): LoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class LoginResponse extends jspb.Message {
  getToken(): string;
  setToken(value: string): LoginResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginResponse): LoginResponse.AsObject;
  static serializeBinaryToWriter(message: LoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginResponse;
  static deserializeBinaryFromReader(message: LoginResponse, reader: jspb.BinaryReader): LoginResponse;
}

export namespace LoginResponse {
  export type AsObject = {
    token: string,
  }
}

export class CreateUserRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): CreateUserRequest;

  getPassword(): string;
  setPassword(value: string): CreateUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
  static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
  static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class CreateUserResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserResponse): CreateUserResponse.AsObject;
  static serializeBinaryToWriter(message: CreateUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserResponse;
  static deserializeBinaryFromReader(message: CreateUserResponse, reader: jspb.BinaryReader): CreateUserResponse;
}

export namespace CreateUserResponse {
  export type AsObject = {
  }
}

