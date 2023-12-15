import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ?
	`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
`: readFileSync(__dirname + "/data.txt", "utf-8")).trim();

const hash = (s: string) =>
	s.split("")
		.reduce<number>((cv, v) => ((cv + v.charCodeAt(0)) * 17) % 256, 0);

const out = dataTxt.split(",").reduce<number>((cv, v) => cv + hash(v), 0);

out;
