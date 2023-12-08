import * as fs from "fs";

/*const dataTxt = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
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

let iterations = 0;
let currentNode = "AAA";
while (currentNode !== "ZZZ") {
	const [left, right] = nodes.get(currentNode)!;
	currentNode = path[iterations % path.length] === "L" ? left : right;
	++iterations;
}

iterations;
