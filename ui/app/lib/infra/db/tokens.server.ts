import { TokensDatabase } from 'manekani-infra-db';
import { database } from './db.server';
import { TokensAdapter } from 'manekani-core';

const tokensDb = new TokensAdapter(new TokensDatabase(database));

export { tokensDb };
