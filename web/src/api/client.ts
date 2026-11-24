import { client } from "./generated/client.gen";

// The SPA is served by `beaver-blueprint serve`, or proxied by the Vite dev
// server, so API paths are same-origin.
client.setConfig({ baseUrl: "" });

export * from "./generated";
