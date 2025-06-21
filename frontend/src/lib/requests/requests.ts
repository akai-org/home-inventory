import type { AnyError } from "./errors";

type Result<Value, Err> = [Value, null] | [null, Err];
export type PResult<V, E> = Promise<Result<V, E>>;

const API_URL = "localhost:8080/api";

async function request<V, E extends AnyError>(method: string, uri: string, init?: RequestInit, headers?: Headers): PResult<V, E> {
	let resp: Response;
	try {
		resp = await fetch(`${API_URL}${uri}`, {
			...init,
			method: method,
			credentials: "same-origin",
			headers: headers,
		});
	} catch (err) {
		console.error("cannot make fetch request:", err);
		return [null, { kind: "FETCH_ERROR", details: "cannot make request" } as E];
	}

	if (!resp.headers.get("content-type")?.includes("application/json")) {
		// unsupported content-type
		console.error("response from server returned with unexpected content-type:", resp.headers.get("contant-type"));
		return [null, { kind: "INTERNAL_SERVER_ERROR", details: "response returend unknown content type" } as E];
	}

	const json = await resp.json();
	if (!resp.ok) return [null, json as E];

	return [json as V, null];
}

async function json_request<V, E extends AnyError>(method: string, uri: string, body?: object): PResult<V, E> {
	const init = { body: JSON.stringify(body ?? {}) };
	const headers = new Headers({ "Content-Type": "application/json" });
	return await request<V, E>(method, uri, init, headers);
}

type OnlyString = { [key: string]: string };
async function uri_body<V, E extends AnyError>(method: string, uri: string, body?: OnlyString): PResult<V, E> {
	let query = "";
	if (body) query = `?${new URLSearchParams(body).toLocaleString()}`;
	return await request<V, E>(method, `${uri}${query}`);
}

export async function get<V, E extends AnyError = AnyError>(uri: string, body?: OnlyString): PResult<V, E> {
	return uri_body("GET", uri, body);
}

export async function post<V, E extends AnyError = AnyError>(uri: string, body?: object): PResult<V, E> {
	return json_request("POST", uri, body);
}
