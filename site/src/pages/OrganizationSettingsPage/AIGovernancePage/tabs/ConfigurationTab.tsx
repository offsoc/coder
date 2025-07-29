import type { Organization } from "api/typesGenerated";
import { Button } from "components/Button/Button";
import {
	HelpTooltip,
	HelpTooltipContent,
	HelpTooltipTrigger,
} from "components/HelpTooltip/HelpTooltip";
import { Link } from "components/Link/Link";
import {
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue,
} from "components/Select/Select";
import { InfoIcon } from "lucide-react";
import { type FC, useState } from "react";
import { docs } from "utils/docs";

interface ConfigurationTabProps {
	organization: Organization;
}

export const ConfigurationTab: FC<ConfigurationTabProps> = ({
	organization,
}) => {
	const [selectedProvider, setSelectedProvider] = useState<string>("anthropic");

	const providers = [
		{ value: "anthropic", label: "Anthropic" },
		{ value: "openai", label: "OpenAI" },
		{ value: "azure", label: "Azure OpenAI" },
		{ value: "google", label: "Google Vertex AI" },
	];

	return (
		<div className="space-y-8">
			<div className="bg-surface-secondary rounded-lg p-6 border border-solid">
				<div className="flex justify-between mb-6">
					<div>
						<h3 className="text-lg font-medium m-0">AI Bridge Configuration</h3>
						<p className="text-content-secondary text-sm mt-1 mb-0">
							Current configuration for AI services in this organization.{" "}
							<Link href={docs("/")}>Learn more</Link>
						</p>
					</div>
					<Button size="sm" disabled>
						Edit Configuration
					</Button>
				</div>

				<div className="grid grid-cols-[200px_1fr] gap-x-6 gap-y-4 text-sm">
					<div className="text-content-secondary font-medium">API Provider</div>
					<div className="font-mono">
						<div className="w-48">
							<Select disabled>
								<SelectTrigger>
									<SelectValue
										placeholder={
											providers.find((p) => p.value === selectedProvider)?.label
										}
									/>
								</SelectTrigger>
								<SelectContent>
									{providers.map((provider) => (
										<SelectItem key={provider.value} value={provider.value}>
											{provider.label}
										</SelectItem>
									))}
								</SelectContent>
							</Select>
						</div>
						<HelpTooltip>
							<HelpTooltipTrigger>
								<InfoIcon className="ml-2 size-icon-xs inline-block text-content-secondary" />
							</HelpTooltipTrigger>
							<HelpTooltipContent>
								The AI service provider used for this organization.
							</HelpTooltipContent>
						</HelpTooltip>
					</div>

					<div className="text-content-secondary font-medium">API Key</div>
					<div className="font-mono">
						••••••••••••••••••
						<HelpTooltip>
							<HelpTooltipTrigger>
								<InfoIcon className="ml-2 size-icon-xs inline-block text-content-secondary" />
							</HelpTooltipTrigger>
							<HelpTooltipContent>
								The API key is masked for security reasons.
							</HelpTooltipContent>
						</HelpTooltip>
					</div>

					<div className="text-content-secondary font-medium">
						API Endpoint URL
					</div>
					<div className="font-mono">
						{selectedProvider === "anthropic" && "https://api.anthropic.com/v1"}
						{selectedProvider === "openai" && "https://api.openai.com/v1"}
						{selectedProvider === "azure" && "https://example.openai.azure.com"}
						{selectedProvider === "google" &&
							"https://us-central1-aiplatform.googleapis.com/v1"}
						<HelpTooltip>
							<HelpTooltipTrigger>
								<InfoIcon className="ml-2 size-icon-xs inline-block text-content-secondary" />
							</HelpTooltipTrigger>
							<HelpTooltipContent>
								The endpoint URL for the AI service API.
							</HelpTooltipContent>
						</HelpTooltip>
					</div>

					<div className="text-content-secondary font-medium">
						Default Model
					</div>
					<div className="font-mono">
						{selectedProvider === "anthropic" && "claude-3-5-sonnet"}
						{selectedProvider === "openai" && "gpt-4o"}
						{selectedProvider === "azure" && "gpt-4-turbo"}
						{selectedProvider === "google" && "gemini-1.5-pro"}
						<HelpTooltip>
							<HelpTooltipTrigger>
								<InfoIcon className="ml-2 size-icon-xs inline-block text-content-secondary" />
							</HelpTooltipTrigger>
							<HelpTooltipContent>
								The default AI model used for requests.
							</HelpTooltipContent>
						</HelpTooltip>
					</div>

					<div className="text-content-secondary font-medium">Rate Limit</div>
					<div className="font-mono">
						100 RPM
						<HelpTooltip>
							<HelpTooltipTrigger>
								<InfoIcon className="ml-2 size-icon-xs inline-block text-content-secondary" />
							</HelpTooltipTrigger>
							<HelpTooltipContent>
								Maximum requests per minute allowed by the API provider.
							</HelpTooltipContent>
						</HelpTooltip>
					</div>

					<div className="text-content-secondary font-medium">Status</div>
					<div>
						<span className="bg-highlight-green/20 text-success-500 px-2 py-1 rounded-full text-xs font-medium">
							Active
						</span>
					</div>
				</div>
			</div>

			<div className="bg-surface-secondary rounded-lg p-6 border border-solid">
				<div className="flex justify-between mb-6">
					<div>
						<h3 className="text-lg font-medium m-0">Policy Settings</h3>
						<p className="text-content-secondary text-sm mt-1 mb-0">
							Organization-wide policies for AI usage.
						</p>
					</div>
					<Button size="sm" disabled>
						Edit Policies
					</Button>
				</div>

				<div className="grid grid-cols-[200px_1fr] gap-x-6 gap-y-4 text-sm">
					<div className="text-content-secondary font-medium">
						Content Filtering
					</div>
					<div>
						<span className="bg-highlight-green/20 text-success-500 px-2 py-1 rounded-full text-xs font-medium">
							Enabled
						</span>
					</div>

					<div className="text-content-secondary font-medium">
						Request Logging
					</div>
					<div>
						<span className="bg-highlight-green/20 text-success-500 px-2 py-1 rounded-full text-xs font-medium">
							Enabled
						</span>
					</div>

					<div className="text-content-secondary font-medium">
						Budget Controls
					</div>
					<div>
						<span className="bg-highlight-warning/20 text-warning-500 px-2 py-1 rounded-full text-xs font-medium">
							Alert Only
						</span>
					</div>

					<div className="text-content-secondary font-medium">
						Data Retention
					</div>
					<div className="font-mono">90 days</div>
				</div>
			</div>
		</div>
	);
};
