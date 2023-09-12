/*
  Warnings:

  - You are about to drop the column `additional_study_data` on the `Subject` table. All the data in the column will be lost.
  - You are about to drop the column `study_data` on the `Subject` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Subject" DROP COLUMN "additional_study_data",
DROP COLUMN "study_data",
ADD COLUMN     "additionalStudyData" JSONB NOT NULL DEFAULT '{}',
ADD COLUMN     "studyData" JSONB NOT NULL DEFAULT '[]';
