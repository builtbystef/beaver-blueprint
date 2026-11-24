import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite-plus";

export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    // Proxy the API to a running `beaver-blueprint serve`.
    proxy: {
      "/api": "http://127.0.0.1:4243",
    },
  },
  build: {
    outDir: "dist",
    emptyOutDir: true,
  },
  fmt: {
    ignorePatterns: ["dist/**", "src/api/generated/**", "openapi.json"],
  },
  lint: {
    plugins: ["eslint", "typescript", "unicorn", "oxc", "react"],
    options: {
      typeAware: true,
      typeCheck: true,
    },
    ignorePatterns: ["dist/**", "src/api/generated/**"],
  },
});
