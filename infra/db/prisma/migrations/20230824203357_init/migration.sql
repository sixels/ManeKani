-- CreateTable
CREATE TABLE "Subject" (
    "id" UUID NOT NULL,
    "category" TEXT NOT NULL,
    "level" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "value" TEXT,
    "valueImage" TEXT,
    "slug" TEXT NOT NULL,
    "priority" INTEGER NOT NULL,
    "resources" JSONB NOT NULL DEFAULT '[]',
    "study_data" JSONB NOT NULL DEFAULT '[]',
    "additional_study_data" JSONB NOT NULL DEFAULT '{}',
    "deckId" UUID NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "Subject_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Deck" (
    "id" UUID NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "ownerId" TEXT NOT NULL,

    CONSTRAINT "Deck_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "User" (
    "id" TEXT NOT NULL,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "_SubjectDependency" (
    "A" UUID NOT NULL,
    "B" UUID NOT NULL
);

-- CreateTable
CREATE TABLE "_SubjectSimilarity" (
    "A" UUID NOT NULL,
    "B" UUID NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "Subject_category_slug_deckId_key" ON "Subject"("category", "slug", "deckId");

-- CreateIndex
CREATE UNIQUE INDEX "Deck_name_ownerId_key" ON "Deck"("name", "ownerId");

-- CreateIndex
CREATE UNIQUE INDEX "_SubjectDependency_AB_unique" ON "_SubjectDependency"("A", "B");

-- CreateIndex
CREATE INDEX "_SubjectDependency_B_index" ON "_SubjectDependency"("B");

-- CreateIndex
CREATE UNIQUE INDEX "_SubjectSimilarity_AB_unique" ON "_SubjectSimilarity"("A", "B");

-- CreateIndex
CREATE INDEX "_SubjectSimilarity_B_index" ON "_SubjectSimilarity"("B");

-- AddForeignKey
ALTER TABLE "Subject" ADD CONSTRAINT "Subject_deckId_fkey" FOREIGN KEY ("deckId") REFERENCES "Deck"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Deck" ADD CONSTRAINT "Deck_ownerId_fkey" FOREIGN KEY ("ownerId") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_SubjectDependency" ADD CONSTRAINT "_SubjectDependency_A_fkey" FOREIGN KEY ("A") REFERENCES "Subject"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_SubjectDependency" ADD CONSTRAINT "_SubjectDependency_B_fkey" FOREIGN KEY ("B") REFERENCES "Subject"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_SubjectSimilarity" ADD CONSTRAINT "_SubjectSimilarity_A_fkey" FOREIGN KEY ("A") REFERENCES "Subject"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_SubjectSimilarity" ADD CONSTRAINT "_SubjectSimilarity_B_fkey" FOREIGN KEY ("B") REFERENCES "Subject"("id") ON DELETE CASCADE ON UPDATE CASCADE;
