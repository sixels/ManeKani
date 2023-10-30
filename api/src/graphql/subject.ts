import {
  CreateSubjectDto,
  Resource as DResource,
  StudyData as DStudyData,
  StudyDataItem as DStudyDataItem,
  Subject as DSubject,
  UpdateSubjectDto,
} from 'manekani-core';
import {
  Field,
  InputType,
  Int,
  ObjectType,
  OmitType,
  PartialType,
  PickType,
} from '@nestjs/graphql';

import { GraphQLJSONObject } from 'graphql-type-json';

@ObjectType()
export class Subject implements DSubject {
  @Field()
  readonly id: string;

  @Field()
  readonly createdAt: Date;
  @Field()
  readonly updatedAt: Date;

  @Field()
  readonly category: string;
  @Field(() => Int)
  readonly level: number;
  @Field()
  readonly name: string;
  @Field()
  readonly slug: string;
  @Field(() => Int)
  readonly priority: number;

  @Field({ nullable: true })
  readonly value?: string;
  @Field({ nullable: true })
  readonly valueImage?: string;

  @Field(() => [StudyData], { defaultValue: [] })
  readonly studyData: StudyData[];
  @Field(() => [Resource], { defaultValue: [] })
  readonly resources: Resource[];

  @Field(() => [String], { defaultValue: [] })
  readonly dependencies: string[];
  @Field(() => [String], { defaultValue: [] })
  readonly dependents: string[];
  @Field(() => [String], { defaultValue: [] })
  readonly similar: string[];

  @Field()
  readonly deckId: string;
  @Field()
  readonly ownerId: string;
}

@ObjectType()
@InputType('StudyDataItemInput')
export class StudyDataItem implements DStudyDataItem {
  @Field()
  readonly value: string;

  @Field()
  readonly isPrimary: boolean;

  @Field()
  readonly isValidAnswer: boolean;

  @Field()
  readonly isHidden: boolean;

  @Field({ nullable: true })
  readonly category?: string;

  @Field({ nullable: true })
  readonly resourceId?: string;
}

@ObjectType()
@InputType('StudyDataInput')
export class StudyData implements DStudyData {
  @Field()
  readonly type: string;

  @Field({ nullable: true })
  readonly mnemonic?: string;

  @Field(() => [StudyDataItem], { defaultValue: [] })
  readonly items: StudyDataItem[];
}

@ObjectType()
@InputType('ResourceInput')
export class Resource implements DResource {
  @Field()
  readonly name: string;

  @Field(() => GraphQLJSONObject)
  readonly metadata: Record<string, any>;
}

@InputType()
export class CreateSubjectInput
  extends PickType(
    Subject,
    [
      'category',
      'level',
      'name',
      'value',
      'valueImage',
      'slug',
      'priority',
      'dependencies',
      'dependents',
      'similar',
      'studyData',
      'resources',
    ] as const,
    InputType,
  )
  implements CreateSubjectDto {}

@InputType()
export class UpdateSubjectInput
  extends PartialType(CreateSubjectInput)
  implements UpdateSubjectDto {}
