import { readFileSync } from "fs";

const testing = true;
const dataTxt = (testing ?
	`19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
`: readFileSync(`${__dirname}/data.txt`, "utf-8"))
	.trim()
	.split("\n")
	.map(v =>
		v
			.split("@")
			.map(w =>
				w
					.trim()
					.split(", ")
			)
	);
