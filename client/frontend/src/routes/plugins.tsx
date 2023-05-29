import { Card, CardHeader, CardTitle } from "@components/ui/card";
import React from "react";

interface PluginsProps {}

export const Plugins: React.FC<PluginsProps> = ({}) => {
  return <div className="container py-12">

    <h1 className="text-5xl font-bold text-primary">Plugins</h1>
    <div className="h-10"/>
    <h2 className="text-2xl font-bold">Plugins available</h2>
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-7 mt-10"> 
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">
          Rexo 
        </CardTitle>
      </CardHeader>
    </Card>
<Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">
          Rexo 
        </CardTitle>
      </CardHeader>
    </Card>

<Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">
          Rexo 
        </CardTitle>
      </CardHeader>
    </Card>
    </div>
  </div>;
};

export default Plugins;
