import { DecksAdapter, TokensAdapter, UsersAdapter } from "@manekani/core";
import { DbClient, DecksDatabase, TokensDatabase, UsersDatabase } from "@manekani/infra-db";
const database = new DbClient();

const tokens = new TokensAdapter(new TokensDatabase(database));
const users = new UsersAdapter(new UsersDatabase(database));
const decks = new DecksAdapter(new DecksDatabase(database));

export { database, tokens, users, decks };
