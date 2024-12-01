import * as fs from "fs";
import * as math from "mathjs";

/*const dataTxt = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`;*/

const dataTxt = fs.readFileSync(__dirname + "/data.txt", "utf-8");

type Node = string;

const path = dataTxt.split("\n")[0].split("");
const nodesData = dataTxt.split("\n\n")[1].trim().split("\n");
const nodes = new Map<Node, [Node, Node]>();

for (let node of nodesData) {
	const target = node.split("=")[1].trim().slice(1, 9).split(", ");
	if (target.length < 2)
		throw new Error("not enough paths on this node");

	nodes.set(node.split(" ")[0], [target[0], target[1]]);
}

const nodesArray = [...nodes];

let iterations = 0;
let currentNodes = nodesArray.filter(v => v[0].endsWith("A"));
let lengths: number[] = [];
while (currentNodes.length > 0) {
	const direction = path[iterations % path.length] === "L" ? 0 : 1;
	let temp: typeof currentNodes = [];
	currentNodes.forEach(node => {
		const res = nodesArray.find(v => v[0] === node[1][direction])!;
		if (node[0].endsWith("Z")) {
			lengths.push(iterations);
			return;
		}
		temp.push(res);
	});
	currentNodes = temp;
	++iterations;
}

console.log(math.lcm(...lengths));
