import type { Organization } from "api/typesGenerated";
import { Badge } from "components/Badge/Badge";
import {
	SettingsHeader,
	SettingsHeaderDescription,
	SettingsHeaderTitle,
} from "components/SettingsHeader/SettingsHeader";
import { TabLink, Tabs, TabsList } from "components/Tabs/Tabs";
import { type FC, useState } from "react";
import { AnalyticsTab } from "./tabs/AnalyticsTab";
import { ConfigurationTab } from "./tabs/ConfigurationTab";
import { RequestLogsTab } from "./tabs/RequestLogsTab";

interface AIGovernancePageViewProps {
	organization: Organization;
}

export const AIGovernancePageView: FC<AIGovernancePageViewProps> = ({
	organization,
}) => {
	const [activeTab, setActiveTab] = useState<"logs" | "analytics" | "config">(
		"logs",
	);

	return (
		<div>
			<SettingsHeader>
				<SettingsHeaderTitle>AI Governance</SettingsHeaderTitle>
				<SettingsHeaderDescription>
					Manage AI policies and usage for your organization.
				</SettingsHeaderDescription>
			</SettingsHeader>

			<Tabs active={activeTab} className="mb-6">
				<TabsList>
					<TabLink
						value="logs"
						to="#"
						onClick={(e) => {
							e.preventDefault();
							setActiveTab("logs");
						}}
					>
						Request Logs
					</TabLink>
					<TabLink
						value="analytics"
						to="#"
						onClick={(e) => {
							e.preventDefault();
							setActiveTab("analytics");
						}}
					>
						Analytics
					</TabLink>
					<TabLink
						value="config"
						to="#"
						onClick={(e) => {
							e.preventDefault();
							setActiveTab("config");
						}}
					>
						Configuration
					</TabLink>
				</TabsList>
			</Tabs>

			<div className="relative">
				<Badge className="absolute right-0 top-0 z-10 bg-highlight-warning border border-solid border-warning-500 text-content-primary">
					Prototype
				</Badge>

				{activeTab === "logs" && <RequestLogsTab organization={organization} />}
				{activeTab === "analytics" && (
					<AnalyticsTab organization={organization} />
				)}
				{activeTab === "config" && (
					<ConfigurationTab organization={organization} />
				)}
			</div>
		</div>
	);
};
