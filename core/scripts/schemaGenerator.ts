import { writeFile } from "fs";
import path from "path";
import { DeckSchema, SubjectSchema, UserSchema } from "../lib/domain";

function generateSchemas() {
	const schemas = {
		user: UserSchema,
		subject: SubjectSchema,
		deck: DeckSchema,
	};

	for (const [objectName, schema] of Object.entries(schemas)) {
		const outputDirectory = path.join(process.cwd(), "schemas/");
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
	}
}

generateSchemas();
