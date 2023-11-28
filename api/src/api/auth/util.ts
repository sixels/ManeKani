import { UserSession } from "@manekani/core";
import { ExecutionContext } from "@nestjs/common";
import { GqlExecutionContext } from "@nestjs/graphql";
import { Request } from "express";

type RequestWithUser = Request & { user?: UserSession };

export function getRequest(context: ExecutionContext): RequestWithUser {
	if ((context.getType() as string) == "graphql") {
		const ctx = GqlExecutionContext.create(context);
		return ctx.getContext().req;
	}
	return context.switchToHttp().getRequest();
}
