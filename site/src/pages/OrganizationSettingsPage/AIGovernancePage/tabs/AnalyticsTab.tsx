import type { Organization } from "api/typesGenerated";
import { Button } from "components/Button/Button";
import {
	Popover,
	PopoverContent,
	PopoverTrigger,
} from "components/Popover/Popover";
import {
	Table,
	TableBody,
	TableCell,
	TableHead,
	TableHeader,
	TableRow,
} from "components/Table/Table";
import dayjs from "dayjs";
import { Check, ChevronDown, DownloadIcon, FilterIcon } from "lucide-react";
import {
	DateRange,
	type DateRangeValue,
} from "pages/TemplatePage/TemplateInsightsPage/DateRange";
import { type FC, useState } from "react";
import {
	Cell,
	Legend,
	Pie,
	PieChart,
	ResponsiveContainer,
	Tooltip,
} from "recharts";
import { aiToolsUsageData, mockTopTasks, mockTopUsers } from "../data/mockData";

interface AnalyticsTabProps {
	organization: Organization;
}

export const AnalyticsTab: FC<AnalyticsTabProps> = ({ organization }) => {
	const [isLoading, setIsLoading] = useState(true);
	const [dateRange, setDateRange] = useState<DateRangeValue>({
		startDate: dayjs().subtract(30, "day").toDate(),
		endDate: dayjs().toDate(),
	});
	const [activeFilters, setActiveFilters] = useState<string[]>([]);

	// Filter options
	const filterOptions = [
		{ value: "claude", label: "Claude Code" },
		{ value: "github", label: "GitHub Copilot" },
		{ value: "amazon", label: "Amazon Q" },
		{ value: "roo", label: "Roo Code" },
	];

	// Simulate loading
	setTimeout(() => {
		setIsLoading(false);
	}, 1000);

	// Get top users and tasks from mock data
	const topUsers = mockTopUsers;
	const topTasks = mockTopTasks;

	// Calculate total stats
	const totalRequests = topUsers.reduce((acc, user) => acc + user.requests, 0);
	const totalCost = topUsers.reduce((acc, user) => acc + user.cost, 0);
	const totalTokens = topUsers.reduce((acc, user) => acc + user.tokens, 0);

	const handleDateRangeChange = (range: DateRangeValue) => {
		setDateRange(range);
	};

	const toggleFilter = (value: string) => {
		setActiveFilters((prev) =>
			prev.includes(value)
				? prev.filter((item) => item !== value)
				: [...prev, value],
		);
	};

	return (
		<div className="space-y-8">
			{/* Date range selector */}
			<div className="flex justify-between items-center">
				<div className="text-sm font-medium">Analytics Overview</div>
				<div className="flex items-center gap-2">
					<Popover>
						<PopoverTrigger asChild>
							<Button
								variant="outline"
								size="sm"
								className="flex items-center gap-1"
							>
								<FilterIcon className="size-icon-xs" />
								Filter
								{activeFilters.length > 0 && (
									<span className="bg-highlight-primary/30 px-1.5 py-0.5 rounded-full text-xs">
										{activeFilters.length}
									</span>
								)}
								<ChevronDown className="size-icon-xs ml-1" />
							</Button>
						</PopoverTrigger>
						<PopoverContent className="w-56 p-2" align="end">
							<div className="text-sm font-medium mb-2 p-1">
								Filter by Provider
							</div>
							<div className="space-y-1">
								{filterOptions.map((option) => (
									<div
										key={option.value}
										className="flex items-center gap-2 rounded p-1.5 hover:bg-surface-active-hover cursor-pointer"
										onClick={() => toggleFilter(option.value)}
									>
										<div className="w-4 h-4 rounded border border-solid flex items-center justify-center">
											{activeFilters.includes(option.value) && (
												<Check className="size-3" />
											)}
										</div>
										<span className="text-sm">{option.label}</span>
									</div>
								))}
							</div>
							{activeFilters.length > 0 && (
								<div className="border-t border-solid pt-2 mt-2">
									<Button
										variant="outline"
										size="sm"
										className="w-full text-xs"
										onClick={() => setActiveFilters([])}
									>
										Clear filters
									</Button>
								</div>
							)}
						</PopoverContent>
					</Popover>
					<DateRange value={dateRange} onChange={handleDateRangeChange} />
				</div>
			</div>

			{/* AI Tool Usage Pie Chart */}
			<div className="grid grid-cols-3 gap-4">
				<div className="bg-surface-secondary rounded-lg p-4 border border-solid">
					<div className="text-sm text-content-secondary mb-1">
						Total Requests
					</div>
					<div className="text-2xl font-medium">
						{totalRequests.toLocaleString()}
					</div>
				</div>
				<div className="bg-surface-secondary rounded-lg p-4 border border-solid">
					<div className="text-sm text-content-secondary mb-1">Total Cost</div>
					<div className="text-2xl font-medium">${totalCost.toFixed(2)}</div>
				</div>
				<div className="bg-surface-secondary rounded-lg p-4 border border-solid">
					<div className="text-sm text-content-secondary mb-1">
						Total Tokens
					</div>
					<div className="text-2xl font-medium">
						{totalTokens.toLocaleString()}
					</div>
				</div>
			</div>

			{/* AI Tool Usage Distribution */}
			<div className="bg-surface-secondary rounded-lg p-4 border border-solid">
				<h3 className="text-base font-medium m-0 mb-4">
					AI Tool Usage Distribution
				</h3>
				<div className="grid grid-cols-[2fr_1fr] gap-6">
					<div className="h-64 flex items-center justify-center">
						<ResponsiveContainer width="100%" height="100%">
							<PieChart>
								<Pie
									data={aiToolsUsageData}
									cx="50%"
									cy="50%"
									innerRadius={70}
									outerRadius={100}
									paddingAngle={2}
									dataKey="value"
									isAnimationActive={false}
								>
									{aiToolsUsageData.map((entry, index) => (
										<Cell
											key={`cell-${index}`}
											fill={entry.color}
											stroke="var(--border)"
											strokeWidth={1}
										/>
									))}
								</Pie>
								<Tooltip
									formatter={(value) => [`${value}%`, "Usage"]}
									contentStyle={{
										backgroundColor: "var(--surface-primary)",
										border: "1px solid var(--border)",
										borderRadius: "4px",
										padding: "8px 12px",
										boxShadow: "0 2px 8px rgba(0, 0, 0, 0.15)",
									}}
								/>
							</PieChart>
						</ResponsiveContainer>
					</div>
					<div className="py-2 pl-4 border-l border-solid flex flex-col justify-center">
						<div className="space-y-5">
							<div className="text-sm font-medium text-content-secondary">
								Usage Breakdown
							</div>
							{aiToolsUsageData.map((item) => (
								<div key={item.name} className="flex items-center gap-3">
									<div
										className="w-4 h-4 rounded-sm flex-shrink-0"
										style={{ backgroundColor: item.color }}
									></div>
									<div className="flex-1 flex items-center justify-between">
										<span className="text-sm font-medium">{item.name}</span>
										<span className="text-sm text-content-secondary">
											{item.value}%
										</span>
									</div>
								</div>
							))}
							<div className="pt-2 mt-2 border-t border-solid">
								<div className="flex items-center justify-between text-sm">
									<span className="font-medium">Total Usage</span>
									<span className="font-medium">100%</span>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>

			{/* Top users table */}
			<div>
				<div className="flex justify-between items-center mb-4">
					<h3 className="text-base font-medium m-0">Top Users by Cost</h3>
					<Button variant="outline" size="sm">
						<DownloadIcon className="size-icon-xs" />
						Export Data
					</Button>
				</div>

				<div className="border border-solid rounded-lg overflow-hidden">
					<Table>
						<TableHeader>
							<TableRow>
								<TableHead>User</TableHead>
								<TableHead>Requests</TableHead>
								<TableHead>Tokens</TableHead>
								<TableHead>Cost</TableHead>
							</TableRow>
						</TableHeader>
						<TableBody>
							{topUsers.map((user, index) => (
								<TableRow key={index}>
									<TableCell>{user.username}</TableCell>
									<TableCell>{user.requests}</TableCell>
									<TableCell>{user.tokens.toLocaleString()}</TableCell>
									<TableCell>${user.cost.toFixed(2)}</TableCell>
								</TableRow>
							))}
						</TableBody>
					</Table>
				</div>
			</div>

			{/* Top tasks table */}
			<div>
				<div className="flex justify-between items-center mb-4">
					<h3 className="text-base font-medium m-0">Top Tasks by Cost</h3>
				</div>

				<div className="border border-solid rounded-lg overflow-hidden">
					<Table>
						<TableHeader>
							<TableRow>
								<TableHead>Task</TableHead>
								<TableHead>User</TableHead>
								<TableHead>Workspace</TableHead>
								<TableHead>Tokens</TableHead>
								<TableHead>Cost</TableHead>
							</TableRow>
						</TableHeader>
						<TableBody>
							{topTasks.map((task, index) => (
								<TableRow key={index}>
									<TableCell className="max-w-sm truncate">
										{task.prompt}
									</TableCell>
									<TableCell>{task.user}</TableCell>
									<TableCell>
										<span className="text-content-link">{task.workspace}</span>
									</TableCell>
									<TableCell>{task.tokens.toLocaleString()}</TableCell>
									<TableCell>${task.cost.toFixed(2)}</TableCell>
								</TableRow>
							))}
						</TableBody>
					</Table>
				</div>
			</div>
		</div>
	);
};
