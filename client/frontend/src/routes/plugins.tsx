import {LoadedPluginCard, PluginCard} from "@components/plugin-card";
import { Button } from "@components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@components/ui/card";
import { Input } from "@components/ui/input";
import { useLoadPluginMutation } from "@hooks/mutations/plugin/load-plugin";
import { useAllLoadedPlugins } from "@hooks/queries/plugins/all-loaded-plugins";
import { useAllPluginsQuery } from "@hooks/queries/plugins/all-plugins";
import { useAllActiveTasks } from "@hooks/queries/tasks/all-active-tasks";
import { Download, SearchIcon, ToyBrick } from "lucide-react";
import React from "react";

interface PluginsProps {}

export const Plugins: React.FC<PluginsProps> = ({}) => {
  const { data: allPluginsData } = useAllPluginsQuery();
  const { data: loadedPluginsData } = useAllLoadedPlugins();

  const { mutate: loadPlugin, data, error } = useLoadPluginMutation();

  const availablePlugins= allPluginsData?.filter(plugin=>loadedPluginsData?!loadedPluginsData.includes(plugin):true)
  return (
    <div className="container py-12">
      <h1 className="text-5xl font-bold text-primary">Plugins</h1>
      <div className="h-10" />
      <div className="flex items-center px-10 py-8">
        <SearchIcon />
        <Input placeholder="Search for plugins" className="ml-4" />
      </div>
      <div className="h-10" />
      <h2 className="text-2xl font-bold">Plugins available</h2>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-12 w-full  mt-10">
        {availablePlugins && availablePlugins.length > 0 ? (
          availablePlugins?.map((plugin) => (
            <Card className="col-span-3">
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium font-black text-lg text-primary">
                  {plugin}
                </CardTitle>
                <ToyBrick className="w-6 h-6 text-primary" />
              </CardHeader>
              <CardFooter className="mt-10">
                <Button
                  className="w-full space-x-4"
                  onClick={() => loadPlugin({ file_name: plugin })}
                >
                  <Download />
                  <p>Load</p>
                </Button>
              </CardFooter>
            </Card>
          ))
        ) : (
          <div className="w-full h-10 items-center justify-centter col-span-12">
            <p className="text-center text-secondary font-bold">
                {allPluginsData && availablePlugins && availablePlugins?.length == 0 ? "No plugins to display, They are all loaded" : "No plugins to dispaly"}
            </p>
          </div>
        )}
      </div>

      <div className="mt-10" />
      <h2 className="text-2xl font-bold">Plugins Loaded</h2>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-12 w-full  mt-10">
        {loadedPluginsData && loadedPluginsData.length > 0 ? (
          loadedPluginsData?.map((plugin) => (
            <LoadedPluginCard plugin={plugin} /> 
          ))
        ) : (
          <div className="w-full h-10 items-center justify-centter col-span-12">
            <p className="text-center text-secondary font-bold">
              No plugins were loaded for now
            </p>
          </div>
        )}
      </div>
    </div>
  );
};

export default Plugins;
