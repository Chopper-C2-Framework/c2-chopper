import React from "react";
import { useParams } from "react-router-dom";
import { usePluginResults } from "@hooks/queries/plugins/plugin-results";
import { usePluginsDetails } from "@hooks/queries/plugins/plugin-detail";
import PluginResultsDisplay from "@components/plugin-results-display";

interface PluginResultsProps {}

export const PluginResults: React.FC<PluginResultsProps> = ({}) => {
  const { pluginName } = useParams();
  if (pluginName == undefined) {
    window.location.href = "/app/plugins";
    return <></>;
  }
  const pluginResults = usePluginResults(pluginName);
  const plugin = usePluginsDetails(pluginName);

  if (pluginResults.isError || plugin.isError) {
    window.location.href = "/app/plugins";
    return <></>;
  }

  return (
    <div className="px-10">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <div className="flex items-center justify-between space-y-2">
          <h2 className="text-3xl font-bold tracking-tight">Plugin Executions</h2>
        </div>
      </div>
      <div className="px-16">
        {pluginResults.data != null && !pluginResults.isLoading && plugin.data && (
          <PluginResultsDisplay
            refresh={()=>{pluginResults.refetch()}}
            plugin={plugin.data}
            results={pluginResults.data.results}
          />
        )}
      </div>
    </div>
  );
};

export default PluginResults;
