import type { WorkspaceApp } from "api/typesGenerated";
import { Button } from "components/Button/Button";
import { Input } from "components/Input/Input";
import { Switch } from "components/Switch/Switch";
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuItem,
	DropdownMenuTrigger,
} from "components/DropdownMenu/DropdownMenu";
import { ExternalImage } from "components/ExternalImage/ExternalImage";
import { InfoTooltip } from "components/InfoTooltip/InfoTooltip";
import {
	ChevronDownIcon,
	LayoutGridIcon,
	ShieldIcon,
	PlusIcon,
	TrashIcon,
	SaveIcon,
} from "lucide-react";
import { useAppLink } from "modules/apps/useAppLink";
import type { Task } from "modules/tasks/tasks";
import type React from "react";
import { type FC, useState, useCallback } from "react";
import { Link as RouterLink } from "react-router-dom";
import { cn } from "utils/cn";
import { TaskAppIFrame } from "./TaskAppIframe";

type TaskAppsProps = {
	task: Task;
};

export const TaskApps: FC<TaskAppsProps> = ({ task }) => {
	const [allowedCommands, setAllowedCommands] = useState<string[]>([
		"npm run dev",
		"npm run test",
	]);
	const [allowedHosts, setAllowedHosts] = useState<string[]>([
		"artifacts.corporate.com",
	]);
	const [auditingEnabled, setAuditingEnabled] = useState<boolean>(true);

	const handleCommandChange = useCallback((index: number, value: string) => {
		setAllowedCommands((prev) => {
			const updated = [...prev];
			updated[index] = value;
			return updated;
		});
	}, []);

	const removeCommand = useCallback((index: number) => {
		setAllowedCommands((prev) => prev.filter((_, i) => i !== index));
	}, []);

	const addCommand = useCallback(() => {
		setAllowedCommands((prev) => [...prev, ""]);
	}, []);

	const handleHostChange = useCallback((index: number, value: string) => {
		setAllowedHosts((prev) => {
			const updated = [...prev];
			updated[index] = value;
			return updated;
		});
	}, []);

	const removeHost = useCallback((index: number) => {
		setAllowedHosts((prev) => prev.filter((_, i) => i !== index));
	}, []);

	const addHost = useCallback(() => {
		setAllowedHosts((prev) => [...prev, ""]);
	}, []);

	const saveSettings = useCallback(() => {
		// Here you would typically save the settings to your backend
		console.log("Saving security settings:", {
			allowedCommands,
			allowedHosts,
			auditingEnabled,
		});
		// For now, we just log to console as this is a prototype
	}, [allowedCommands, allowedHosts, auditingEnabled]);
	const agents = task.workspace.latest_build.resources
		.flatMap((r) => r.agents)
		.filter((a) => !!a);

	// The Chat UI app will be displayed in the sidebar, so we don't want to show
	// it here
	const apps = agents
		.flatMap((a) => a?.apps)
		.filter(
			(a) => !!a && a.id !== task.workspace.latest_build.ai_task_sidebar_app_id,
		);

	const embeddedApps = apps.filter((app) => !app.external);
	const externalApps = apps.filter((app) => app.external);

	const [activeAppId, setActiveAppId] = useState<string>(() => {
		const appId = embeddedApps[0]?.id;
		if (!appId) {
			throw new Error("No apps found in task");
		}
		return appId;
	});

	const activeApp = apps.find((app) => app.id === activeAppId);
	if (!activeApp) {
		throw new Error(`Active app with ID ${activeAppId} not found in task`);
	}

	const agent = agents.find((a) =>
		a.apps.some((app) => app.id === activeAppId),
	);
	if (!agent) {
		throw new Error(`Agent for app ${activeAppId} not found in task workspace`);
	}

	return (
		<main className="flex flex-col">
			<div className="w-full flex items-center border-0 border-b border-border border-solid">
				<div className="p-2 pb-0 flex gap-2 items-center">
					{embeddedApps.map((app) => (
						<TaskAppTab
							key={app.id}
							task={task}
							app={app}
							active={app.id === activeAppId}
							onClick={(e) => {
								e.preventDefault();
								setActiveAppId(app.id);
							}}
						/>
					))}
				</div>

				<div className="ml-auto flex gap-2">
					<DropdownMenu>
						<DropdownMenuTrigger asChild>
							<Button size="sm" variant="subtle">
								<ShieldIcon className="mr-1.5 h-4 w-4" />
								Security
								<ChevronDownIcon className="ml-1" />
							</Button>
						</DropdownMenuTrigger>
						<DropdownMenuContent align="end" className="w-[400px]">
							<div className="p-4 space-y-4">
								<div className="border-b border-border pb-2 mb-2 flex items-center justify-between">
									<h3 className="text-lg font-medium">Security Settings</h3>
								</div>

								<div>
									<h4 className="text-sm font-medium mb-2">Allowed Commands</h4>
									<div className="space-y-2">
										{allowedCommands.map((command, index) => (
											<div key={index} className="flex items-center gap-2">
												<Input
													className="flex-1 font-mono text-sm"
													value={command}
													onChange={(e) =>
														handleCommandChange(index, e.target.value)
													}
												/>
												<Button
													size="icon"
													variant="outline"
													onClick={() => removeCommand(index)}
												>
													<TrashIcon className="h-4 w-4" />
													<span className="sr-only">Remove</span>
												</Button>
											</div>
										))}
										<Button
											size="sm"
											variant="outline"
											className="w-full mt-1"
											onClick={addCommand}
										>
											<PlusIcon className="h-4 w-4 mr-1" />
											Add Command
										</Button>
									</div>
								</div>

								<div className="mt-4">
									<h4 className="text-sm font-medium mb-2">Allowed Hosts</h4>
									<div className="space-y-2">
										{allowedHosts.map((host, index) => (
											<div key={index} className="flex items-center gap-2">
												<Input
													className="flex-1 font-mono text-sm"
													value={host}
													onChange={(e) =>
														handleHostChange(index, e.target.value)
													}
												/>
												<Button
													size="icon"
													variant="outline"
													onClick={() => removeHost(index)}
												>
													<TrashIcon className="h-4 w-4" />
													<span className="sr-only">Remove</span>
												</Button>
											</div>
										))}
										<Button
											size="sm"
											variant="outline"
											className="w-full mt-1"
											onClick={addHost}
										>
											<PlusIcon className="h-4 w-4 mr-1" />
											Add Host
										</Button>
									</div>
								</div>

								<div className="flex items-center justify-between mt-4">
									<span className="text-sm font-medium">Auditing</span>
									<Switch
										checked={auditingEnabled}
										onCheckedChange={setAuditingEnabled}
									/>
								</div>

								<Button
									className="w-full mt-4"
									size="sm"
									onClick={saveSettings}
								>
									<SaveIcon className="h-4 w-4 mr-1.5" />
									Save Settings
								</Button>
							</div>
						</DropdownMenuContent>
					</DropdownMenu>

					{externalApps.length > 0 && (
						<DropdownMenu>
							<DropdownMenuTrigger asChild>
								<Button size="sm" variant="subtle">
									Open locally
									<ChevronDownIcon />
								</Button>
							</DropdownMenuTrigger>
							<DropdownMenuContent>
								{externalApps.map((app) => {
									const link = useAppLink(app, {
										agent,
										workspace: task.workspace,
									});

									return (
										<DropdownMenuItem key={app.id} asChild>
											<RouterLink to={link.href}>
												{app.icon ? (
													<ExternalImage src={app.icon} />
												) : (
													<LayoutGridIcon />
												)}
												{link.label}
											</RouterLink>
										</DropdownMenuItem>
									);
								})}
							</DropdownMenuContent>
						</DropdownMenu>
					)}
				</div>
			</div>

			<div className="flex-1">
				{embeddedApps.map((app) => {
					return (
						<TaskAppIFrame
							key={app.id}
							active={activeAppId === app.id}
							app={app}
							task={task}
						/>
					);
				})}
			</div>
		</main>
	);
};

type TaskAppTabProps = {
	task: Task;
	app: WorkspaceApp;
	active: boolean;
	onClick: (e: React.MouseEvent<HTMLAnchorElement>) => void;
};

const TaskAppTab: FC<TaskAppTabProps> = ({ task, app, active, onClick }) => {
	const agent = task.workspace.latest_build.resources
		.flatMap((r) => r.agents)
		.filter((a) => !!a)
		.find((a) => a.apps.some((a) => a.id === app.id));

	if (!agent) {
		throw new Error(`Agent for app ${app.id} not found in task workspace`);
	}

	const link = useAppLink(app, {
		agent,
		workspace: task.workspace,
	});

	return (
		<Button
			size="sm"
			variant="subtle"
			key={app.id}
			asChild
			className={cn([
				"px-3",
				{
					"text-content-primary bg-surface-tertiary rounded-sm rounded-b-none":
						active,
				},
				{ "opacity-75 hover:opacity-100": !active },
			])}
		>
			<RouterLink to={link.href} onClick={onClick}>
				{app.icon ? <ExternalImage src={app.icon} /> : <LayoutGridIcon />}
				{link.label}
				{app.health === "unhealthy" && (
					<InfoTooltip
						title="This app is unhealthy."
						message="The health check failed."
						type="warning"
					/>
				)}
			</RouterLink>
		</Button>
	);
};
