import { type FC } from "react";
import {
	Cell,
	Legend,
	Pie,
	PieChart,
	ResponsiveContainer,
	Tooltip,
} from "recharts";
import { aiToolsUsageData } from "../data/mockData";

interface AIToolsUsageChartProps {
	className?: string;
}

export const AIToolsUsageChart: FC<AIToolsUsageChartProps> = ({
	className,
}) => {
	return (
		<div
			className={`bg-surface-secondary rounded-lg p-4 border border-solid ${className}`}
		>
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
								label={({ name, percent }) => `${(percent * 100).toFixed(0)}%`}
								labelLine={true}
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
	);
};
