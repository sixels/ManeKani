import { TokensAdapter } from "@manekani/core";
import { TokensDatabase } from "@manekani/infra-db";
import { database } from "./db.server";

const tokensDb = new TokensAdapter(new TokensDatabase(database));

export { tokensDb };
