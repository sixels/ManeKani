import path from 'path';
import { DeckSchema, SubjectSchema, UserSchema } from '../lib/domain';
import { writeFile } from 'fs';

function generateSchemas() {
  const schemas = {
    user: UserSchema,
    subject: SubjectSchema,
    deck: DeckSchema,
  };

  Object.entries(schemas).forEach(([objectName, schema]) => {
    const outputDirectory = path.join(process.cwd(), 'schemas/');
    const outputFileName = `${objectName}.schema.json`;

    writeFile(
      path.join(outputDirectory, outputFileName),
      JSON.stringify(schema, null, 2),
      (err) => {
        if (err) {
          console.error(err);
          return;
        }
        console.log(`Schema ${objectName} generated`);
      },
    );
  });
}

generateSchemas();
