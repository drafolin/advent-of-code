import * as fs from "fs";

fs.readFile("./data.txt", "utf8", (err, data) => {
	const numbers = [
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine"
	];

	let sum = 0;
	for (let v of data.split("\n")) {
		if (v === "")
			continue;

		console.log("Original:", v);

		let iterator = v.matchAll(/(?=(one|two|three|four|five|six|seven|eight|nine|zero|\d))/g);
		let matches = [...iterator].map(v => v[1]).map(v => isNaN(parseInt(v)) ? numbers.indexOf(v) : v);

		console.log("matches", matches);

		let string = matches.join("");

		for (let i in numbers) {
			string = string.replace(numbers[i], i.toString());
		}

		console.log("Parsed: ", string);


		let firstDigit: string | undefined;
		let lastDigit: string | undefined;


		for (let w of string.split("")) {
			if (isNaN(parseInt(w)))
				continue;

			if (firstDigit === undefined)
				firstDigit = w;

			lastDigit = w;
		}

		if (!firstDigit || !lastDigit)
			throw new Error("no lol");

		let result = `${firstDigit}${lastDigit}`;
		console.log("result:", result);
		sum += parseInt(result);
	}

	console.log(sum);
});
