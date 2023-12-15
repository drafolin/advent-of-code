import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ?
	`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
`: readFileSync(__dirname + "/data.txt", "utf-8")).trim();

type lens = { label: string, focalLength: number; };

const hash = (s: string) =>
	s.split("")
		.reduce<number>((cv, v) => ((cv + v.charCodeAt(0)) * 17) % 256, 0);

const boxes = new Map<number, lens[]>();

dataTxt.split(",").forEach((i) => {
	if (i.includes("-")) {
		const [label] = i.split("-");
		let box = boxes.get(hash(label));
		if (box) {
			const index = box.findIndex((v) => v.label === label);
			if (index !== -1) {
				box.splice(index, 1);
				boxes.set(hash(label), box);
			}
		}
	} else {
		const [label, focalLength] = i.split("=");
		let box = boxes.get(hash(label)) ?? [];
		let index = box.findIndex(v => v.label === label);
		if (index === -1)
			box.push({ label, focalLength: parseInt(focalLength) });
		else
			box[index] = { label, focalLength: parseInt(focalLength) };
		boxes.set(hash(label), box);
	}
});

let power = 0;

for (let [i, v] of boxes) {
	for (let [j, lens] of v.entries()) {
		power += (i + 1) *
			(j + 1) *
			lens.focalLength;
	}
}

power;
