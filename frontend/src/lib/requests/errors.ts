export type ErrorKind = "FETCH_ERROR" | "INTERNAL_SERVER_ERROR";

export type BaseError = { kind: ErrorKind };
export type FetchEror = BaseError & { kind: "FETCH_ERROR"; details: string };
export type InternalServerError = BaseError & { kind: "INTERNAL_SERVER_ERROR"; details: string };

export type AnyError = FetchEror | InternalServerError;
