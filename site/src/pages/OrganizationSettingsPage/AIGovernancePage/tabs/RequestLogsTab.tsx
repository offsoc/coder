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
	const [expandedToolCalls, setExpandedToolCalls] = useState<
		Record<string, boolean>
	>({});

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
												<div className="p-4">
													<dl className="text-xs m-0 grid grid-cols-[auto_1fr] gap-x-4 gap-y-3 items-start [&_dt]:font-medium [&_dt]:text-content-secondary [&_dd]:text-content-primary">
														<dt>Request ID:</dt>
														<dd className="font-mono">{log.id}</dd>

														<dt>Timestamp:</dt>
														<dd className="font-mono">{log.timestamp}</dd>

														<dt>User:</dt>
														<dd>{log.user}</dd>

														<dt>Model:</dt>
														<dd className="font-mono">claude-3-5-sonnet</dd>

														<dt>Status:</dt>
														<dd>
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
														</dd>

														<dt>Duration:</dt>
														<dd className="font-mono">
															{Math.round(log.tokens / 30)}ms
														</dd>

														<dt>Total Tokens:</dt>
														<dd className="font-mono">
															{log.tokens.toLocaleString()}
														</dd>

														<dt className="flex items-center gap-1">
															<span>Token Usage:</span>
														</dt>
														<dd>
															<div className="flex items-center gap-3">
																<div className="flex items-center gap-1">
																	<div className="w-3 h-3 rounded-full bg-surface-tertiary"></div>
																	<span className="font-mono">
																		{Math.round(log.tokens * 0.3)} prompt
																	</span>
																</div>
																<div className="flex items-center gap-1">
																	<div className="w-3 h-3 rounded-full bg-surface-primary-hover"></div>
																	<span className="font-mono">
																		{Math.round(log.tokens * 0.7)} completion
																	</span>
																</div>
															</div>
														</dd>

														<dt>Tool Calls:</dt>
														<dd className="font-mono">{log.toolCalls}</dd>

														<dt>Cost:</dt>
														<dd className="font-mono">
															${log.cost.toFixed(4)}
														</dd>
													</dl>

													<div className="mt-4">
														<h4 className="text-sm font-medium mb-2 flex items-center gap-1">
															<span>Prompt</span>
														</h4>
														<div className="bg-surface-tertiary p-3 rounded font-mono text-sm whitespace-pre-wrap border border-solid border-border">
															{log.prompt}
														</div>
													</div>

													{log.toolCalls > 0 && (
														<div className="mt-4">
															<h4 className="text-sm font-medium mb-2 flex items-center gap-1">
																<span>Tool Usage</span>
															</h4>
															<div className="bg-surface-tertiary p-3 rounded border border-solid border-border">
																<div className="space-y-2">
																	{Array.from(
																		{ length: Math.min(log.toolCalls, 5) },
																		(_, i) => {
																			const tool = [
																				"Read",
																				"Write",
																				"Bash",
																				"Grep",
																				"Task",
																				"MultiEdit",
																			][i % 6];
																			return (
																				<div
																					key={i}
																					className="flex items-start gap-2"
																				>
																					<div className="bg-surface-primary-hover text-content-primary px-2 py-1 rounded text-xs font-medium mt-0.5">
																						{tool}
																					</div>
																					<div className="text-xs">
																						<span className="font-mono text-content-secondary">
																							{tool === "Read"
																								? "Read file at path: /home/user/project/src/main.ts"
																								: tool === "Write"
																									? "Created file at path: /home/user/project/src/utils.ts"
																									: tool === "Bash"
																										? "Executed command: ls -la"
																										: tool === "Grep"
																											? "Searched for pattern: function main"
																											: tool === "Task"
																												? "Executed task: analyze codebase"
																												: "Modified multiple files in a single operation"}
																						</span>
																					</div>
																				</div>
																			);
																		},
																	)}
																	{log.toolCalls > 5 &&
																		!expandedToolCalls[log.id] && (
																			<div
																				className="text-xs text-content-secondary pl-7 flex items-center gap-1 cursor-pointer"
																				onClick={(e) => {
																					e.stopPropagation();
																					setExpandedToolCalls((prev) => ({
																						...prev,
																						[log.id]: true,
																					}));
																				}}
																			>
																				<span>+ {log.toolCalls - 5} more</span>
																				<ChevronDown className="size-icon-xxs" />
																			</div>
																		)}
																	{expandedToolCalls[log.id] && (
																		<>
																			{Array.from(
																				{
																					length: Math.min(
																						log.toolCalls - 5,
																						10,
																					),
																				},
																				(_, i) => {
																					const idx = i + 5;
																					const tool = [
																						"Read",
																						"Write",
																						"Bash",
																						"Grep",
																						"Task",
																						"MultiEdit",
																					][idx % 6];
																					return (
																						<div
																							key={idx}
																							className="flex items-start gap-2"
																						>
																							<div className="bg-surface-primary-hover text-content-primary px-2 py-1 rounded text-xs font-medium mt-0.5">
																								{tool}
																							</div>
																							<div className="text-xs">
																								<span className="font-mono text-content-secondary">
																									{tool === "Read"
																										? `Read file at path: /home/user/project/src/components/utils${idx}.ts`
																										: tool === "Write"
																											? `Created file at path: /home/user/project/src/helpers/format${idx}.ts`
																											: tool === "Bash"
																												? `Executed command: git status --short`
																												: tool === "Grep"
																													? `Searched for pattern: function render${idx}`
																													: tool === "Task"
																														? `Executed task: format code in file ${idx}`
																														: `Modified ${idx} files in a single operation`}
																								</span>
																							</div>
																						</div>
																					);
																				},
																			)}
																			{log.toolCalls > 15 && (
																				<div className="text-xs text-content-secondary pl-7">
																					{log.toolCalls > 15 &&
																						`+ ${log.toolCalls - 15} more tool calls not shown`}
																				</div>
																			)}
																		</>
																	)}
																</div>
															</div>
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
