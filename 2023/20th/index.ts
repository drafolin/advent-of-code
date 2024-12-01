import { readFileSync } from "fs";
import { lcm } from "mathjs";

const testingNr: number = 0;
const dataTxt = (testingNr === 1 ?
	`broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`: testingNr === 2 ? `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
` : readFileSync(__dirname + "/data.txt", "utf-8")
).trim().split("\n");


const HIGH = true;
const LOW = false;

type Signal = typeof HIGH | typeof LOW;

let queue: { from: string, to: string, signal: Signal; }[] = [];

class Module {
	outputs: string[];
	name: string;
	private lastOut: boolean;
	constructor(outputs: string[], name: string) {
		this.name = name;
		this.outputs = outputs;
		this.lastOut = false;
	}

	input = (signal: Signal, input: string) => {
		this.outputs.forEach(v => {
			queue.push({ to: v, from: this.name, signal: signal });
		});
		this.lastOut = signal;
	};

	state = () => this.lastOut;

	connect = (input: string) => { };
}

class DebugModule extends Module {
	constructor(name: string) {
		super([], name);
	}

	input = (signal: Signal, input: string) => {
		console.log("recieved signal: ", { signal, input });
	};
}

class FlipFlop extends Module {
	memory: Signal;
	constructor(outputs: string[], name: string) {
		super(outputs, name);
		this.memory = false;
	}

	input = (signal: Signal) => {
		if (signal === LOW) {
			this.memory = !this.memory;
			this.outputs.forEach(v => {
				queue.push({ to: v, from: this.name, signal: this.memory });
			});
		}
	};

	state = () => this.memory;

	connect = (input: string) => { };
}

let finalGate = "";
let counts: number[] = [];
let count = 0;

class Conjunction extends Module {
	memory: Map<string, boolean>;
	constructor(outputs: string[], name: string) {
		super(outputs, name);
		this.memory = new Map<string, boolean>();
	}

	input = (signal: boolean, input: string) => {
		if (this.name === finalGate && signal === HIGH)
			counts.push(count);
		this.memory.set(input, signal);
		if ([...this.memory.values()].every(v => v === HIGH)) {
			this.outputs.forEach(v => {
				queue.push({ to: v, from: this.name, signal: LOW });
			});
		}
		else {
			this.outputs.forEach(v => {
				queue.push({ to: v, from: this.name, signal: HIGH });
			});
		}
	};

	state = () => [...this.memory.values()].every(v => v === HIGH) ? LOW : HIGH;

	connect = (input: string) => {
		this.memory.set(input, LOW);
	};
}

const nodes = new Map<string, Module>();

for (let node of dataTxt) {
	const [name, outputs] = node.split(" -> ");
	switch (name[0]) {
		case "%": nodes.set(name.slice(1), new FlipFlop(outputs.split(", "), name.slice(1))); break;
		case "&": nodes.set(name.slice(1), new Conjunction(outputs.split(", "), name.slice(1))); break;
		default: nodes.set(name, new Module(outputs.split(", "), name)); break;
	}
}

for (let node of nodes.keys()) {
	nodes.get(node)!.outputs.forEach(v => {
		if (!nodes.has(v)) {
			if (v === "rx")
				finalGate = node;
			nodes.set(v, new DebugModule(v));
		}

		nodes.get(v)!.connect(node);
	});
}

const pushButton = () => {
	queue.push({ from: "", to: "broadcaster", signal: LOW });

	while (queue.length > 0) {
		const instruction = queue.shift()!;
		nodes.get(instruction.to)!.input(instruction.signal, instruction.from);
	}
};

while (counts.length < 4) {
	count++;
	pushButton();
}
console.log(lcm(lcm(counts[0], counts[1]), lcm(counts[2], counts[3])));
