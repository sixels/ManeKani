-- AlterTable
ALTER TABLE "User" ADD COLUMN     "displayName" TEXT,
ADD COLUMN     "isComplete" BOOLEAN NOT NULL DEFAULT false;
