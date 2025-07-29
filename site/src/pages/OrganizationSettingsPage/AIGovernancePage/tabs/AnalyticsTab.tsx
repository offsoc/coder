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
import { AIToolsUsageChart } from "../components/AIToolsUsageChart";
import { mockTopTasks, mockTopUsers } from "../data/mockData";

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
			<AIToolsUsageChart className="mb-0" />

			{/* Top users table */}
			<div>
				<div className="flex justify-between items-center mb-4">
					<h3 className="text-base font-medium m-0">Top Users by Cost</h3>
					<Button variant="outline" size="sm">
						<DownloadIcon className="size-icon-xs" />
						Export Data
					</Button>
				</div>

				<div className="overflow-hidden rounded-md border border-solid border-border">
					<Table>
						<TableHeader>
							<TableRow className="bg-surface-secondary hover:bg-surface-secondary">
								<TableHead className="font-medium">User</TableHead>
								<TableHead className="font-medium">Requests</TableHead>
								<TableHead className="font-medium">Tokens</TableHead>
								<TableHead className="font-medium">Cost</TableHead>
							</TableRow>
						</TableHeader>
						<TableBody>
							{topUsers.map((user, index) => (
								<TableRow
									key={index}
									className="border-t border-solid border-border"
								>
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

				<div className="overflow-hidden rounded-md border border-solid border-border">
					<Table>
						<TableHeader>
							<TableRow className="bg-surface-secondary hover:bg-surface-secondary">
								<TableHead className="font-medium">Task</TableHead>
								<TableHead className="font-medium">User</TableHead>
								<TableHead className="font-medium">Workspace</TableHead>
								<TableHead className="font-medium">Tokens</TableHead>
								<TableHead className="font-medium">Cost</TableHead>
							</TableRow>
						</TableHeader>
						<TableBody>
							{topTasks.map((task, index) => (
								<TableRow
									key={index}
									className="border-t border-solid border-border"
								>
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
