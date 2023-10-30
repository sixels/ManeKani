import { Injectable, OnModuleInit } from '@nestjs/common';
import { DbClient } from 'manekani-infra-db';

@Injectable()
export class DatabaseService extends DbClient {
  onModuleInit() {
    this.$connect();
  }
}
