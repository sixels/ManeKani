import { DbClient } from "@manekani/infra-db";
import { Injectable, OnModuleInit } from "@nestjs/common";

@Injectable()
export class DatabaseService extends DbClient implements OnModuleInit {
	onModuleInit() {
		this.$connect();
	}
}
