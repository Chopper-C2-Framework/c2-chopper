import { SearchIcon } from "lucide-react";
import React from "react";
import { Input } from "./ui/input";
import usePluginsSearchStore from "@src/stores/plugins/search";
import { Button } from "./ui/button";

interface PluginSearchProps {}

export const PluginSearch: React.FC<PluginSearchProps> = ({}) => {
    const { clearSearchField,updateSearchField}=usePluginsSearchStore()
  return (
    <div className="flex items-center px-10 py-8">
      <SearchIcon />
      <Input placeholder="Search for plugins" className="ml-4" onChange={({target})=>updateSearchField(target.value)}/>
      <Button variant="outline" className="ml-4" onClick={()=>clearSearchField()}>Clear</Button>
    </div>
  );
};

export default PluginSearch;
