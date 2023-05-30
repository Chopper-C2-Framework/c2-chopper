import { Plugin, PluginResult } from "@src/types";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@components/ui/card";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { useDeletePluginResult } from "@hooks/mutations/plugin/delete-plugin-result";

interface IPluginResultsDisplay {
  results: PluginResult[];
  plugin: Plugin;
  refresh: () => void;
}

export default function PluginResultsDisplay({plugin, results, refresh}: IPluginResultsDisplay){
  const deletePluginResult = useDeletePluginResult()
  return (
    <div className="container gap-5 px-8 py-5 flex flex-col align-center">
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center">
            <div>
              {plugin.info.Name}
              <CardDescription className="pl-2 font-normal">
                Version: {plugin.Metadata.version}
              </CardDescription>
            </div>
          </CardTitle>
        </CardHeader>
        <CardContent className="flex flex-row pr-20">
          <div className="w-1/2 flex flex-col gap-2">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="agentId" className="text-right">
                Author
              </Label>
              {
                <Input id="agentId" value={plugin.Metadata.author} readOnly className="col-span-3"/>
              }
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="args" className="text-right">
                Description
              </Label>
              <Input
                id="args"
                value={plugin.Metadata.description}
                readOnly
                className="col-span-3"
              />
            </div>
          </div>
          <div className="w-1/2 flex flex-col gap-2">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="creatorId" className="text-right">
                Return Type
              </Label>
              <Input
                id="creatorId"
                value={plugin.info.ReturnType}
                readOnly
                className="col-span-3"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="createdAt" className="text-right">
                Release Date
              </Label>
              <Input
                id="createdAt"
                value={plugin.Metadata.release_date}
                readOnly
                className="col-span-3"
              />
            </div>
          </div>
        </CardContent>
      </Card>
      
      <div className="container px-10">
        {
          results.map((result, idx) => (
            <Card key={idx} >
              <CardHeader className="flex flex-row justify-between">
                <CardTitle className="flex items-center">
                  <div>
                    {result.createdAt}
                    <CardDescription className="pl-2 font-normal">
                      Output
                    </CardDescription>
                  </div>
                </CardTitle>
                <Button onClick={ async ()=> {
                  await deletePluginResult.mutateAsync({
                    result_id: result.id
                  })
                  refresh()
                }}>Delete</Button>
              </CardHeader>
              <CardContent className="flex flex-row pr-20 bg-black p-10">
                <pre>{result.output}</pre>
              </CardContent>
            </Card>
          ))
        }
      </div>
    </div>
  );
}
