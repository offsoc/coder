import type { Organization } from "api/typesGenerated";
import { Button } from "components/Button/Button";
import { EmptyState } from "components/EmptyState/EmptyState";
import { SearchField } from "components/SearchField/SearchField";
import {
	Table,
	TableBody,
	TableCell,
	TableHead,
	TableHeader,
	TableRow,
} from "components/Table/Table";
import { TableLoader } from "components/TableLoader/TableLoader";
import {
	ChevronDown,
	ChevronRight,
	DownloadIcon,
	FilterIcon,
} from "lucide-react";
import { type FC, useState } from "react";
import { mockRequestLogs } from "../data/mockData";

interface RequestLogsTabProps {
	organization: Organization;
}

export const RequestLogsTab: FC<RequestLogsTabProps> = ({ organization }) => {
	const [isLoading, setIsLoading] = useState(true);
	const [searchQuery, setSearchQuery] = useState("");
	const [expandedRows, setExpandedRows] = useState<Record<string, boolean>>({});

	// Simulate loading
	setTimeout(() => {
		setIsLoading(false);
	}, 1000);

	const logs = mockRequestLogs;
	const filteredLogs = logs.filter(
		(log) =>
			log.prompt.toLowerCase().includes(searchQuery.toLowerCase()) ||
			log.user.toLowerCase().includes(searchQuery.toLowerCase()) ||
			log.status.toLowerCase().includes(searchQuery.toLowerCase()),
	);

	const toggleRowExpand = (id: string) => {
		setExpandedRows((prev) => ({
			...prev,
			[id]: !prev[id],
		}));
	};

	return (
		<div>
			<div className="flex gap-2 items-center mb-4">
				<SearchField
					className="w-full"
					placeholder="Search by user, prompt, or status..."
					value={searchQuery}
					onChange={setSearchQuery}
				/>
				<Button variant="outline">
					<FilterIcon className="size-icon-xs" />
					Filter
				</Button>
				<Button variant="outline">
					<DownloadIcon className="size-icon-xs" />
					Export
				</Button>
			</div>

			<div className="border border-solid rounded-lg overflow-hidden">
				<Table>
					<TableHeader>
						<TableRow>
							<TableHead className="w-10"></TableHead>
							<TableHead>Timestamp</TableHead>
							<TableHead>User</TableHead>
							<TableHead>Prompt</TableHead>
							<TableHead>Tokens</TableHead>
							<TableHead>Tool Calls</TableHead>
							<TableHead>Cost</TableHead>
							<TableHead>Status</TableHead>
						</TableRow>
					</TableHeader>
					<TableBody>
						{isLoading ? (
							<TableLoader />
						) : filteredLogs.length === 0 ? (
							<TableRow>
								<TableCell colSpan={8}>
									<EmptyState
										message="No logs found"
										description="Try adjusting your search or filters."
									/>
								</TableCell>
							</TableRow>
						) : (
							filteredLogs.map((log) => (
								<>
									<TableRow
										key={log.id}
										className="cursor-pointer hover:bg-surface-active-hover"
										onClick={() => toggleRowExpand(log.id)}
									>
										<TableCell>
											{expandedRows[log.id] ? (
												<ChevronDown className="size-icon-xs" />
											) : (
												<ChevronRight className="size-icon-xs" />
											)}
										</TableCell>
										<TableCell className="whitespace-nowrap">
											{log.timestamp}
										</TableCell>
										<TableCell>{log.user}</TableCell>
										<TableCell className="max-w-sm truncate">
											{log.prompt}
										</TableCell>
										<TableCell>{log.tokens}</TableCell>
										<TableCell>{log.toolCalls}</TableCell>
										<TableCell>${log.cost.toFixed(4)}</TableCell>
										<TableCell>
											<span
												className={`px-2 py-1 rounded-full text-xs font-medium ${
													log.status === "completed"
														? "bg-highlight-green/20 text-success-500"
														: log.status === "error"
															? "bg-highlight-error/20 text-error-500"
															: "bg-highlight-warning/20 text-warning-500"
												}`}
											>
												{log.status}
											</span>
										</TableCell>
									</TableRow>
									{expandedRows[log.id] && (
										<TableRow key={`${log.id}-expanded`}>
											<TableCell></TableCell>
											<TableCell colSpan={7}>
												<div className="py-4 px-2">
													<h4 className="text-sm font-medium mb-2">
														Prompt Details
													</h4>
													<div className="bg-surface-tertiary p-3 rounded mb-3 font-mono text-sm whitespace-pre-wrap">
														{log.prompt}
													</div>

													<h4 className="text-sm font-medium mb-2">
														Response Summary
													</h4>
													<div className="bg-surface-tertiary p-3 rounded mb-3 font-mono text-sm">
														<div className="grid grid-cols-2 gap-4">
															<div>
																<p className="text-xs text-content-secondary mb-1">
																	Completion Tokens
																</p>
																<p>{Math.round(log.tokens * 0.7)}</p>
															</div>
															<div>
																<p className="text-xs text-content-secondary mb-1">
																	Prompt Tokens
																</p>
																<p>{Math.round(log.tokens * 0.3)}</p>
															</div>
															<div>
																<p className="text-xs text-content-secondary mb-1">
																	Model
																</p>
																<p>claude-3-5-sonnet</p>
															</div>
															<div>
																<p className="text-xs text-content-secondary mb-1">
																	Duration
																</p>
																<p>{Math.round(log.tokens / 30)}ms</p>
															</div>
														</div>
													</div>

													<h4 className="text-sm font-medium mb-2">
														Tool Calls
													</h4>
													{log.toolCalls > 0 ? (
														<div className="bg-surface-tertiary p-3 rounded font-mono text-sm">
															<ul className="list-disc pl-5 space-y-2">
																{Array.from(
																	{ length: Math.min(log.toolCalls, 5) },
																	(_, i) => (
																		<li key={i} className="text-xs">
																			<span className="text-content-secondary">
																				Tool:{" "}
																			</span>
																			<span className="font-medium">
																				{
																					[
																						"Read",
																						"Write",
																						"Bash",
																						"Grep",
																						"Task",
																						"MultiEdit",
																					][i % 6]
																				}
																			</span>
																		</li>
																	),
																)}
																{log.toolCalls > 5 && (
																	<li className="text-xs text-content-secondary">
																		+ {log.toolCalls - 5} more tool calls
																	</li>
																)}
															</ul>
														</div>
													) : (
														<div className="text-sm text-content-secondary">
															No tool calls for this request
														</div>
													)}
												</div>
											</TableCell>
										</TableRow>
									)}
								</>
							))
						)}
					</TableBody>
				</Table>
			</div>
		</div>
	);
};
