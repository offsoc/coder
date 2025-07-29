import type { Organization } from "api/typesGenerated";
import { Button } from "components/Button/Button";
import {
	Table,
	TableBody,
	TableCell,
	TableHead,
	TableHeader,
	TableRow,
} from "components/Table/Table";
import dayjs from "dayjs";
import { DownloadIcon } from "lucide-react";
import {
	DateRange,
	type DateRangeValue,
} from "pages/TemplatePage/TemplateInsightsPage/DateRange";
import { type FC, useState } from "react";
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

	return (
		<div className="space-y-8">
			{/* Date range selector */}
			<div className="flex justify-between items-center">
				<div className="text-sm font-medium">Analytics Overview</div>
				<DateRange value={dateRange} onChange={handleDateRangeChange} />
			</div>

			{/* Top metrics section */}
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
