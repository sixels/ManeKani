import { CreateDeckDto, Deck as DDeck, UpdateDeckDto } from "@manekani/core";
import { Field, InputType, ObjectType, PartialType } from "@nestjs/graphql";

@ObjectType()
export class Deck implements DDeck {
	@Field()
	readonly id: string;

	@Field()
	readonly createdAt: Date;
	@Field()
	readonly updatedAt: Date;

	@Field()
	readonly name: string;

	@Field()
	readonly description: string;
	@Field(() => [String], { defaultValue: [] })
	readonly subjectIds: string[];
	
	@Field()
	readonly ownerId: string;

	@Field(() => [String], { defaultValue: [] })
	readonly subscribedUserIds: string[];
}

@InputType()
export class CreateDeckInput implements CreateDeckDto {
	@Field()
	name: string;
	@Field()
	description: string;
}

@InputType()
export class UpdateDeckInput
	extends PartialType(CreateDeckInput)
	implements UpdateDeckDto {}
