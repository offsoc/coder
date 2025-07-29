// Mock data generation functions
// We're using random data but with fixed seed for consistent results
const randomInt = (min: number, max: number) => {
	return Math.floor(Math.random() * (max - min + 1)) + min;
};

const randomChoice = <T>(arr: T[]): T => {
	return arr[Math.floor(Math.random() * arr.length)];
};

const randomFullName = (): string => {
	const firstNames = [
		"John",
		"Jane",
		"Michael",
		"Emily",
		"David",
		"Sarah",
		"Chris",
		"Emma",
		"Alex",
		"Olivia",
	];
	const lastNames = [
		"Smith",
		"Johnson",
		"Williams",
		"Brown",
		"Jones",
		"Miller",
		"Davis",
		"Garcia",
		"Rodriguez",
		"Wilson",
	];

	return `${randomChoice(firstNames)} ${randomChoice(lastNames)}`;
};

const uuid = (): string => {
	return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, (c) => {
		const r = (Math.random() * 16) | 0;
		const v = c === "x" ? r : (r & 0x3) | 0x8;
		return v.toString(16);
	});
};

// Mock data for request logs
interface RequestLog {
	id: string;
	timestamp: string;
	user: string;
	prompt: string;
	tokens: number;
	toolCalls: number;
	cost: number;
	status: "completed" | "error" | "throttled";
}

// Generate a random set of request logs
const generateRequestLogs = (count: number): RequestLog[] => {
	return Array.from({ length: count }, () => {
		const tokens = randomInt(50, 8000);
		const toolCalls = randomInt(0, 15);

		// Calculate a realistic cost based on tokens (approximately $10 per million tokens)
		const cost = (tokens / 1000000) * 10;

		// Generate random status with weighted distribution
		const statusRandom = Math.random();
		let status: RequestLog["status"];
		if (statusRandom < 0.9) {
			status = "completed";
		} else if (statusRandom < 0.95) {
			status = "error";
		} else {
			status = "throttled";
		}

		// Generate a random date within the last 14 days
		const now = new Date();
		const daysAgo = randomInt(0, 14);
		const date = new Date(now);
		date.setDate(date.getDate() - daysAgo);
		const dateStr = date.toISOString().split("T")[0];
		const timeStr = date.toTimeString().split(" ")[0];

		return {
			id: uuid(),
			timestamp: `${dateStr} ${timeStr}`,
			user: randomFullName(),
			prompt: randomChoice([
				"Help me debug this code",
				"Generate a SQL query for user analytics",
				"Analyze this log file for errors",
				"Write a unit test for this function",
				"Explain how this algorithm works",
				"Refactor this component for better performance",
				"Create a regex pattern to match email addresses",
				"Convert this function to async/await",
				"Optimize this database query",
				"Find security vulnerabilities in this code",
			]),
			tokens,
			toolCalls,
			cost,
			status,
		};
	});
};

// Mock data for top users
interface TopUser {
	username: string;
	requests: number;
	tokens: number;
	cost: number;
}

// AI tool usage data
interface AIToolUsage {
	name: string;
	value: number;
	color: string;
}

export const aiToolsUsageData: AIToolUsage[] = [
	{ name: "Claude Code", value: 65, color: "#7953C3" },
	{ name: "GitHub Copilot", value: 20, color: "#3FB950" },
	{ name: "Amazon Q", value: 10, color: "#FF9900" },
	{ name: "Roo Code", value: 5, color: "#FF5630" },
];

// Generate mock top users
const generateTopUsers = (count: number): TopUser[] => {
	return Array.from({ length: count }, () => {
		const requests = randomInt(10, 500);
		const tokens = requests * randomInt(500, 5000);
		const cost = (tokens / 1000000) * 10;

		return {
			username: randomFullName(),
			requests,
			tokens,
			cost,
		};
	}).sort((a, b) => b.cost - a.cost);
};

// Mock data for top tasks
interface TopTask {
	prompt: string;
	user: string;
	workspace: string;
	tokens: number;
	cost: number;
}

// Generate mock top tasks
const generateTopTasks = (count: number): TopTask[] => {
	return Array.from({ length: count }, () => {
		const tokens = randomInt(5000, 50000);
		const cost = (tokens / 1000000) * 10;

		return {
			prompt: randomChoice([
				"Analyze codebase structure and suggest improvements",
				"Generate comprehensive test suite for authentication service",
				"Design database schema for new feature",
				"Code review and suggest security improvements",
				"Optimize frontend performance and suggest changes",
				"Debug complex race condition in concurrent code",
				"Create documentation for API endpoints",
				"Refactor legacy module to modern standards",
				"Implement error handling strategy across application",
				"Create deployment pipeline configuration",
			]),
			user: randomFullName(),
			workspace: randomChoice([
				"dev-environment",
				"staging-env",
				"prod-support",
				"data-science",
				"backend-api",
				"frontend-app",
				"test-automation",
				"devops-tools",
			]),
			tokens,
			cost,
		};
	}).sort((a, b) => b.cost - a.cost);
};

// Create the mock data
export const mockRequestLogs = generateRequestLogs(20);
export const mockTopUsers = generateTopUsers(10);
export const mockTopTasks = generateTopTasks(10);
