import { ErrorAlert } from "components/Alert/ErrorAlert";
import { useOrganizationSettings } from "modules/management/OrganizationSettingsLayout";
import { RequirePermission } from "modules/permissions/RequirePermission";
import type { FC } from "react";
import { Helmet } from "react-helmet-async";
import { pageTitle } from "utils/page";
import { AIGovernancePageView } from "./AIGovernancePageView";

const AIGovernancePage: FC = () => {
	const { organization, organizationPermissions } = useOrganizationSettings();

	if (!organization) {
		return <ErrorAlert error="Organization not found" />;
	}

	const helmet = (
		<Helmet>
			<title>
				{pageTitle(
					"AI Governance",
					organization.display_name || organization.name,
				)}
			</title>
		</Helmet>
	);

	// For now, reuse the same permissions as other settings pages
	if (!organizationPermissions?.editSettings) {
		return (
			<>
				{helmet}
				<RequirePermission isFeatureVisible={false} />
			</>
		);
	}

	return (
		<>
			{helmet}
			<AIGovernancePageView organization={organization} />
		</>
	);
};

export default AIGovernancePage;
