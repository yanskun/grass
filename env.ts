import { config } from "https://deno.land/x/dotenv@v2.0.0/mod.ts";

const {
  READ_USER_TOKEN,
} = config({ safe: true });

export {
  READ_USER_TOKEN,
};
