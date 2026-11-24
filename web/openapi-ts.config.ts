import { defineConfig } from "@hey-api/openapi-ts";

// Generated from openapi.json (written by the Go server's openapigen) by
// `just openapi`. Both are committed so the frontend builds without Go.
export default defineConfig({
  input: "./openapi.json",
  output: "src/api/generated",
  plugins: ["@hey-api/client-fetch"],
});
