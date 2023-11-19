import { TokensAdapter, UsersAdapter } from 'manekani-core';
import { DbClient, TokensDatabase, UsersDatabase } from 'manekani-infra-db';
const database = new DbClient();

const tokens = new TokensAdapter(new TokensDatabase(database));
const users = new UsersAdapter(new UsersDatabase(database));

export { database, tokens, users };
