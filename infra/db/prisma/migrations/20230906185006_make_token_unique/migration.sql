/*
  Warnings:

  - A unique constraint covering the columns `[token]` on the table `ApiToken` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "ApiToken_token_key" ON "ApiToken"("token");
