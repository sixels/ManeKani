import { PrismaClient } from "@prisma/client";

export class DbClient extends PrismaClient {}

export * from "./decks";
export * from "./subjects";
export * from "./tokens";
export * from "./users";
